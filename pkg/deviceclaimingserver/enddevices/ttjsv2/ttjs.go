// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package ttjsv2 provides the claiming client implementation for The Things Join Server 2.0 API.
package ttjsv2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"go.thethings.network/lorawan-stack/v3/pkg/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/config"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/httpclient"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"google.golang.org/grpc"
)

// BasicAuth contains HTTP basic auth settings.
type BasicAuth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// NetworkServer contains information related to the Network Server.
type NetworkServer struct {
	Hostname string
	HomeNSID *types.EUI64
}

// Config is the configuration to communicate with The Things Join Server End Device Claming API.
type Config struct {
	NetID           types.NetID         `yaml:"-"`
	JoinEUIPrefixes []types.EUI64Prefix `yaml:"-"`
	NetworkServer   NetworkServer       `yaml:"-"`

	BasicAuth `yaml:"basic-auth"`
	URL       string `yaml:"url"`
}

// Component abstracts the underlying *component.Component.
type Component interface {
	httpclient.Provider
	GetBaseConfig(ctx context.Context) config.ServiceBase
	GetPeerConn(ctx context.Context, role ttnpb.ClusterRole, ids cluster.EntityIdentifiers) (*grpc.ClientConn, error)
	AllowInsecureForCredentials() bool
}

// TTJS is a client that claims end devices on a The Things Join Server.
type TTJS struct {
	Component

	httpClient *http.Client
	baseURL    *url.URL
	config     Config
}

// NewClient applies the config and returns a new TTJS client.
func (cfg *Config) NewClient(ctx context.Context, c Component) (*TTJS, error) {
	httpClient, err := c.HTTPClient(ctx)
	if err != nil {
		return nil, err
	}
	baseURL, err := url.Parse(cfg.URL)
	if err != nil {
		return nil, err
	}
	return &TTJS{
		config:     *cfg,
		httpClient: httpClient,
		baseURL:    baseURL,
		Component:  c,
	}, nil
}

// SupportsJoinEUI implements EndDeviceClaimer.
func (client *TTJS) SupportsJoinEUI(eui types.EUI64) bool {
	for _, prefix := range client.config.JoinEUIPrefixes {
		if eui.HasPrefix(prefix) {
			return true
		}
	}
	return false
}

var (
	errDeviceNotProvisioned = errors.DefineNotFound("device_not_provisioned", "device with EUI `{dev_eui}` not provisioned") //nolint:lll
	errDeviceNotClaimed     = errors.DefineNotFound("device_not_claimed", "device with EUI `{dev_eui}` not claimed")
	errDeviceAccessDenied   = errors.DefinePermissionDenied("device_access_denied", "access to device with `{dev_eui}` denied. Either device is already claimed or owner token is invalid") //nolint:lll
	errUnauthorized         = errors.DefineUnauthenticated("unauthorized", "client API Key missing or invalid")
)

// Claim implements EndDeviceClaimer.
func (client *TTJS) Claim(ctx context.Context, joinEUI, devEUI types.EUI64, claimAuthenticationCode string) error {
	reqURL := fmt.Sprintf("%s/api/v2/devices/%s/claim", client.baseURL.String(), devEUI.String())
	logger := log.FromContext(ctx).WithFields(log.Fields(
		"dev_eui", devEUI,
		"join_eui", joinEUI,
		"url", reqURL,
	))

	claimReq := &claimRequest{
		OwnerToken: claimAuthenticationCode,
		Lock:       true,
		HomeNetID:  client.config.NetID.String(),
	}
	if client.config.NetworkServer.HomeNSID != nil {
		claimReq.HomeNSID = new(string)
		*claimReq.HomeNSID = client.config.NetworkServer.HomeNSID.String()
	}
	buf, err := json.Marshal(claimReq)
	if err != nil {
		return err
	}

	logger.Debug("Claim end device")
	request, err := http.NewRequestWithContext(ctx, http.MethodPut, reqURL, bytes.NewReader(buf))
	if err != nil {
		return err
	}
	request.SetBasicAuth(client.config.BasicAuth.Username, client.config.BasicAuth.Password)
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if isSuccess(resp.StatusCode) {
		return nil
	}

	var errResp errorResponse
	err = json.Unmarshal(respBody, &errResp)
	if err != nil {
		logger.WithError(err).Warn("Failed to decode error message")
	} else {
		logger.WithField("error", errResp.Message).Warn("Failed to claim end device")
	}

	switch resp.StatusCode {
	case http.StatusNotFound:
		return errDeviceNotProvisioned.WithAttributes("dev_eui", devEUI)
	case http.StatusForbidden:
		return errDeviceAccessDenied.WithAttributes("dev_eui", devEUI)
	case http.StatusUnauthorized:
		return errUnauthorized.New()
	default:
		return errors.FromHTTPStatusCode(resp.StatusCode)
	}
}

// Unclaim implements EndDeviceClaimer.
func (client *TTJS) Unclaim(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers) error {
	devEUI := types.MustEUI64(ids.DevEui)
	reqURL := fmt.Sprintf("%s/api/v2/devices/%s/claim", client.baseURL.String(), devEUI.String())
	logger := log.FromContext(ctx).WithFields(log.Fields(
		"dev_eui", devEUI,
		"join_eui", ids.JoinEui,
		"url", reqURL,
	))

	logger.Debug("Unclaim end device")
	request, err := http.NewRequestWithContext(ctx, http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(client.config.BasicAuth.Username, client.config.BasicAuth.Password)
	resp, err := client.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if isSuccess(resp.StatusCode) {
		return nil
	}

	var errResp errorResponse
	err = json.Unmarshal(respBody, &errResp)
	if err != nil {
		logger.WithError(err).Warn("Failed to decode error message")
	} else {
		logger.WithField("error", errResp.Message).Warn("Failed to unclaim end device")
	}

	switch resp.StatusCode {
	case http.StatusNotFound:
		return errDeviceNotClaimed.WithAttributes("dev_eui", devEUI)
	case http.StatusForbidden:
		return errDeviceAccessDenied.WithAttributes("dev_eui", devEUI)
	case http.StatusUnauthorized:
		return errUnauthorized.New()
	default:
		return errors.FromHTTPStatusCode(resp.StatusCode)
	}
}

// GetClaimStatus implements EndDeviceClaimer.
func (client *TTJS) GetClaimStatus(
	ctx context.Context, ids *ttnpb.EndDeviceIdentifiers,
) (*ttnpb.GetClaimStatusResponse, error) {
	devEUI := types.MustEUI64(ids.DevEui)
	reqURL := fmt.Sprintf("%s/api/v2/devices/%s/claim", client.baseURL.String(), devEUI.String())
	logger := log.FromContext(ctx).WithFields(log.Fields(
		"dev_eui", devEUI,
		"join_eui", ids.JoinEui,
		"url", reqURL,
	))

	logger.Debug("Get claim status for end device")
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(client.config.BasicAuth.Username, client.config.BasicAuth.Password)
	resp, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if isSuccess(resp.StatusCode) {
		var (
			claimData claimData
			ret       = ttnpb.GetClaimStatusResponse{
				EndDeviceIds: ids,
			}
			homeNSID  types.EUI64
			homeNetID types.NetID
		)
		err = json.Unmarshal(respBody, &claimData)
		if err != nil {
			return nil, err
		}
		err = homeNetID.UnmarshalText([]byte(claimData.HomeNetID))
		if err != nil {
			return nil, err
		}
		ret.HomeNetId = homeNetID.Bytes()
		if claimData.HomeNSID != nil {
			err = homeNSID.UnmarshalText([]byte(*claimData.HomeNSID))
			if err != nil {
				return nil, err
			}
			ret.HomeNsId = homeNSID.Bytes()
		}
		return &ret, nil
	}

	var errResp errorResponse
	err = json.Unmarshal(respBody, &errResp)
	if err != nil {
		logger.WithError(err).Warn("Failed to decode error message")
	} else {
		logger.WithField("error", errResp.Message).Warn("Failed to get claim status")
	}

	switch resp.StatusCode {
	case http.StatusNotFound:
		return nil, errDeviceNotClaimed.WithAttributes("dev_eui", devEUI)
	case http.StatusForbidden:
		return nil, errDeviceAccessDenied.WithAttributes("dev_eui", devEUI)
	case http.StatusUnauthorized:
		return nil, errUnauthorized.New()
	default:
		return nil, errors.FromHTTPStatusCode(resp.StatusCode)
	}
}

// isSuccess returns true if the HTTP status code is 2xx.
func isSuccess(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}
