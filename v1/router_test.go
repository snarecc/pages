package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubVersionController struct {
}

func NewStubVersionController() VersionController {
	return &StubVersionController{}
}

func (c *StubVersionController) HandlerFunc() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("the body"))
		},
	)
}

func TestRouter(t *testing.T) {
	router := NewRouter(
		NewStubVersionController(),
	)

	t.Run("Route not found", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		gotStatusCode := response.Result().StatusCode
		wantStatusCode := 404

		if gotStatusCode != wantStatusCode {
			t.Errorf("got status code %d want %d", gotStatusCode, wantStatusCode)
		}

		gotBody := string(response.Body.Bytes())
		wantBody := "404 page not found\n"

		if gotBody != wantBody {
			t.Errorf("got body %s want %s", gotBody, wantBody)
		}
	})

	t.Run("Route /version to version controller", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/version", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		gotStatusCode := response.Result().StatusCode
		wantStatusCode := 200

		if gotStatusCode != wantStatusCode {
			t.Errorf("got status code %d want %d", gotStatusCode, wantStatusCode)
		}

		gotBody := string(response.Body.Bytes())
		wantBody := "the body"

		if gotBody != wantBody {
			t.Errorf("got body %s want %s", gotBody, wantBody)
		}
	})
}
