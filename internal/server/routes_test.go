package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHandlers(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	s := &Server{}

	testHandler(req, s.EnglishHandler, t)
}

func testHandler(req *http.Request, h echo.HandlerFunc, t *testing.T) {
	e := echo.New()

	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)

	// Assertions
	if err := h(c); err != nil {
		t.Errorf("handler() error = %v", err)
		return
	}
	if resp.Code != http.StatusOK {
		t.Errorf("handler() wrong status code = %v", resp.Code)
		return
	}

	// Result test
	// expected := map[string]string{"message": "Hello World"}
	//
	// var actual map[string]string
	// // Decode the response body into the actual map
	// if err := json.NewDecoder(resp.Body).Decode(&actual); err != nil {
	// 	t.Errorf("handler() error decoding response body: %v", err)
	// 	return
	// }
	// // Compare the decoded response with the expected value
	// if !reflect.DeepEqual(expected, actual) {
	// 	t.Errorf("handler() wrong response body. expected = %v, actual = %v", expected, actual)
	// 	return
	// }
}
