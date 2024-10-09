package suite

import (
	"bytes"
	"encoding/json"
	"fmt"
	tripHTTP "github.com/darksubmarine/booking-fly/domain/entities/trip/inputs/http/gin"
	"github.com/darksubmarine/booking-fly/domain/entities/user"
	userHTTP "github.com/darksubmarine/booking-fly/domain/entities/user/inputs/http/gin"
	"github.com/darksubmarine/booking-fly/testing/fixtures"
	"github.com/darksubmarine/torpedo-lib-go/entity"
	"net/http"
	"net/http/httptest"
)

// TestBookingFly_UserNotFound
// This test tries to book a fly with a unknown user.
func (s *DomainSuite) TestBookingFly_UserNotFound() {

	var data = map[string]interface{}{
		"departure": "CFO", "arrival": "JFK",
		"miles": 2700, "from": 1726837954000, "to": 1727529154000, "userId": fixtures.UserId}

	jsonData, jsonErr := json.Marshal(data)
	s.Nil(jsonErr)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/booking", bytes.NewBuffer(jsonData))
	recorder := httptest.NewRecorder()

	s.router.ServeHTTP(recorder, request)

	s.EqualValues(http.StatusBadRequest, recorder.Result().StatusCode)
	s.JSONEq(`{"code":"4007","error":"user not found with given Id"}`, s.body(recorder))
}

// helperFetchUser helper function to fetch a user given its ID
func (s *DomainSuite) helperFetchUser(userId string) *user.UserEntity {
	s.NotEmpty(userId)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/users/%s", userId), nil)
	recorder := httptest.NewRecorder()

	s.router.ServeHTTP(recorder, request)
	s.EqualValues(http.StatusOK, recorder.Result().StatusCode)

	_body := s.body(recorder)

	var userDTO userHTTP.FullDTO
	s.Nil(json.Unmarshal([]byte(_body), &userDTO))

	toRet := user.New()
	s.Nil(entity.From(&userDTO, toRet))

	return toRet
}

// helperCreateUser helper function to create a user and return it
func (s *DomainSuite) helperCreateUser() *user.UserEntity {

	var userDTO userHTTP.WriteableDTO
	userModel := fixtures.User()
	s.Nil(entity.To(userModel, &userDTO))

	jsonData, jsonErr := json.Marshal(userDTO)
	s.Nil(jsonErr)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(jsonData))
	recorder := httptest.NewRecorder()

	s.router.ServeHTTP(recorder, request)

	s.EqualValues(http.StatusOK, recorder.Result().StatusCode)

	_body := s.body(recorder)

	var createdUserDTO userHTTP.FullDTO
	s.Nil(json.Unmarshal([]byte(_body), &createdUserDTO))

	toRet := user.New()
	s.Nil(entity.From(&createdUserDTO, toRet))

	return toRet
}

// TestBookingFly_Ok test that books a fly with a valid user and get user plan upgraded.
//
//	Given a user frequent flyer with plan BRONZE and accumulated miles 1500
//	When the user books a trip from the airport CFO  to the airport JFK with a trip distance of 2700 miles
//	Then the trip is booked successfully and the user plan is upgraded to SILVER and the accumulated miles is 4500
func (s *DomainSuite) TestBookingFly_Ok() {
	userModel := s.helperCreateUser()
	tripModel := fixtures.Trip(userModel.Id())

	var tripDTO tripHTTP.WriteableDTO
	s.Nil(entity.To(tripModel, &tripDTO))

	jsonData, jsonErr := json.Marshal(tripDTO)
	s.Nil(jsonErr)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/booking", bytes.NewBuffer(jsonData))
	recorder := httptest.NewRecorder()

	s.router.ServeHTTP(recorder, request)

	s.EqualValues(http.StatusOK, recorder.Result().StatusCode)

	_body := s.body(recorder)

	var tripDTOCreated tripHTTP.FullDTO
	s.Nil(json.Unmarshal([]byte(_body), &tripDTOCreated))

	s.NotEmpty(tripDTOCreated.Id())

	// test user status previous to booking fly
	s.EqualValues(fixtures.UserMiles, userModel.Miles())
	s.EqualValues(fixtures.UserPlanBronze, userModel.Plan())

	// test user status after booking fly
	updatedUserModel := s.helperFetchUser(userModel.Id())
	s.EqualValues(fixtures.UserUpdatedMiles, updatedUserModel.Miles())
	s.EqualValues(fixtures.UserPlanSilver, updatedUserModel.Plan())
}
