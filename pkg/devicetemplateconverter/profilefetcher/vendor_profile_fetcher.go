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

package profilefetcher

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// vendorIDProfileFetcher handles the validation and fetching of end-device profile in the Device Repository.
type vendorIDProfileFetcher struct{}

// NewFetcherByVendorIDs returns a end-device's profile fetcher that builds its request with vendorIDs.
func NewFetcherByVendorIDs() EndDeviceProfileFetcher {
	return &vendorIDProfileFetcher{}
}

// ShouldFetchProfile dictactes if the end-device has the necessary fields to fetch its profile.
func (*vendorIDProfileFetcher) ShouldFetchProfile(device *ttnpb.EndDevice) bool {
	return device.GetVersionIds().GetVendorId() != 0 && device.GetVersionIds().GetVendorProfileId() != 0
}

// FetchProfile provides the end-device profile.
func (*vendorIDProfileFetcher) FetchProfile(
	ctx context.Context,
	device *ttnpb.EndDevice,
) (*ttnpb.EndDeviceTemplate, error) {
	profileIdentifiers := &ttnpb.GetTemplateRequest_EndDeviceProfileIdentifiers{
		VendorId:        device.GetVersionIds().GetVendorId(),
		VendorProfileId: device.GetVersionIds().GetVendorProfileId(),
	}

	fetcher, ok := fetcherFromContext(ctx)
	if !ok {
		return nil, nil
	}
	return fetcher.GetTemplate(ctx, &ttnpb.GetTemplateRequest{EndDeviceProfileIds: profileIdentifiers})
}
