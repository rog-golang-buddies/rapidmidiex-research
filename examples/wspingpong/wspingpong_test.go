package wspingpong

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNoFavicon(t *testing.T) {
	srv := httptest.NewServer(WSPingPongServer{})

	// TODO: currently the server just returns for this route. What happens then exactly?

	req, err := http.NewRequest("GET", srv.URL+"/favicon.ico", strings.NewReader(""))
	if err != nil {
		t.Errorf("Request error: %v\n", err)
	}
	res, err := srv.Client().Do(req)
	if err != nil {
		t.Errorf("Client error: %v\n", err)
	}
	t.Errorf("Response: %v\n", res)
}

func TestAllowOnlyGet(t *testing.T) {

	methods := []string{"POST", "PUT", "PATCH", "HACK"}

	srv := httptest.NewServer(WSPingPongServer{LogLevel: LogLevelBasicWithHeaders})

	for _, m := range methods {
		req, err := http.NewRequest(m, srv.URL, strings.NewReader(""))
		if err != nil {
			t.Errorf("%v\n", err)
		}
		res, err := srv.Client().Do(req)
		if err != nil {
			t.Errorf("%v\n", err)
		}
		if res.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Status should be MethodNotAllowed but is %d.\n", res.StatusCode)
		}
	}

}
