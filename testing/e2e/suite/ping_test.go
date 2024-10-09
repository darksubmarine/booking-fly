package suite

import (
	"net/http"
	"net/http/httptest"
)

// TestPing exec a call to test the endpoint '/ping'
func (s *DomainSuite) TestPing() {
	request := httptest.NewRequest(http.MethodGet, "/ping", nil)

	recorder := httptest.NewRecorder()

	s.router.ServeHTTP(recorder, request)

	s.EqualValues(http.StatusOK, recorder.Result().StatusCode)
	s.JSONEq(`{"message":"pong"}`, s.body(recorder))
}
