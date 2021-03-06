//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

package anonymous

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/semi-technologies/weaviate/usecases/config"
	"github.com/stretchr/testify/assert"
)

func Test_AnonymousMiddleware_Enabled(t *testing.T) {
	// when anonymous access is enabled, we don't need to do anything and can
	// savely call the next next handler

	r := httptest.NewRequest("GET", "/foo", nil)
	w := httptest.NewRecorder()

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(900)
	})

	cfg := config.Config{
		Authentication: config.Authentication{
			AnonymousAccess: config.AnonymousAccess{
				Enabled: true,
			},
		},
	}

	New(cfg).Middleware(next).ServeHTTP(w, r)
	response := w.Result()

	assert.Equal(t, response.StatusCode, 900)
}

func Test_AnonymousMiddleware_Disabled(t *testing.T) {
	t.Run("when OIDC is enabled, but no token provided", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/foo", nil)
		w := httptest.NewRecorder()

		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(900)
		})

		cfg := config.Config{
			Authentication: config.Authentication{
				AnonymousAccess: config.AnonymousAccess{
					Enabled: false,
				},
				OIDC: config.OIDC{
					Enabled: true,
				},
			},
		}

		New(cfg).Middleware(next).ServeHTTP(w, r)
		response := w.Result()

		assert.Equal(t, response.StatusCode, 401)
	})

	t.Run("when OIDC is enabled, and a Bearer Header provided", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/foo", nil)
		r.Header.Add("Authorization", "Bearer foo")
		w := httptest.NewRecorder()

		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(900)
		})

		cfg := config.Config{
			Authentication: config.Authentication{
				AnonymousAccess: config.AnonymousAccess{
					Enabled: false,
				},
				OIDC: config.OIDC{
					Enabled: true,
				},
			},
		}

		New(cfg).Middleware(next).ServeHTTP(w, r)
		response := w.Result()

		assert.Equal(t, response.StatusCode, 900)
	})
}
