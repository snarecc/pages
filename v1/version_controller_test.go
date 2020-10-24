package v1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestVersionController(t *testing.T) {
	t.Run("GET returns version and sha1", func(t *testing.T) {
		versionController := NewVersionController(
			NewBuild(
				"oogabooga",
				"boogaooga",
			),
		)

		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		versionController.HandlerFunc().ServeHTTP(response, request)

		gotStatusCode := response.Result().StatusCode
		wantStatusCode := 200

		if gotStatusCode != wantStatusCode {
			t.Errorf("got status code %d want %d", gotStatusCode, wantStatusCode)
		}

		var gotBody map[string]string
		if err := json.NewDecoder(response.Body).Decode(&gotBody); err != nil {
			t.Fatal(err)
		}

		wantBody := map[string]string{
			"sha1":    "oogabooga",
			"version": "boogaooga",
		}

		if !reflect.DeepEqual(gotBody, wantBody) {
			t.Errorf("got body %q want %q", gotBody, wantBody)
		}
	})
}
