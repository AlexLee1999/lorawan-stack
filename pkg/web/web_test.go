// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package web

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/random"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestIsZeros(t *testing.T) {
	a := assertions.New(t)
	{
		res := isZeros([]byte{0, 0, 0, 0, 0})
		a.So(res, should.BeTrue)
	}
	{
		res := isZeros([]byte{0, 0, 0, 1, 0})
		a.So(res, should.BeFalse)
	}
}

func TestStatic(t *testing.T) {
	a := assertions.New(t)
	s, err := New(test.Context())
	if !a.So(err, should.BeNil) {
		t.Fatal("Could not create a web instance")
	}

	dir, err := os.Getwd()

	if !a.So(err, should.BeNil) {
		t.Fatal("Could not create resolve testing directory")
	}

	s.Static("/assets", http.Dir(dir))

	// HTTP server returns 200 on valid file request
	{
		req := httptest.NewRequest(http.MethodGet, "/assets/web_test.go", nil)
		rec := httptest.NewRecorder()

		s.ServeHTTP(rec, req)

		resp := rec.Result()
		body, _ := io.ReadAll(resp.Body)

		a.So(resp.StatusCode, should.Equal, http.StatusOK)
		a.So(strings.HasPrefix(string(body), "//"), should.BeTrue)
	}

	// HTTP server returns 404 on invalid file request
	{
		req := httptest.NewRequest(http.MethodGet, "/assets/null.txt", nil)
		rec := httptest.NewRecorder()

		s.Static("/assets", http.Dir(dir+"/teststatic"))

		s.ServeHTTP(rec, req)

		resp := rec.Result()
		a.So(resp.StatusCode, should.Equal, http.StatusNotFound)
	}
}

func TestCookies(t *testing.T) {
	a := assertions.New(t)
	// Errors on illegal hash key byte size
	{
		_, err := New(test.Context(), WithCookieKeys(random.Bytes(2), nil))
		a.So(err, should.NotBeNil)
	}
	// Errors on non 32bit block key
	{
		_, err := New(test.Context(), WithCookieKeys(nil, random.Bytes(31)))
		a.So(err, should.NotBeNil)
	}
}
