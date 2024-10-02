// Package booking_fly Fly reservation use case

package booking_fly

import (
	"github.com/darksubmarine/booking-fly/domain/entities/trip"
	tripMocks "github.com/darksubmarine/booking-fly/domain/entities/trip/testing/mocks"
	"github.com/darksubmarine/booking-fly/domain/entities/user"
	userMocks "github.com/darksubmarine/booking-fly/domain/entities/user/testing/mocks"
	"github.com/darksubmarine/torpedo-lib-go/context"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"time"

	"testing"
)

// TestUseCase_DoBooking test to validate the use case logic.
//
//	Given a user frequent flyer with plan BRONZE and accumulated miles 1500
//	When the user books a trip from the airport CFO  to the airport JFK with a trip distance of 2700 miles
//	Then the trip is booked successfully and the user plan is upgraded to SILVER and the accumulated miles is 4500
func TestUseCase_DoBooking(t *testing.T) {

	// mocked user Id
	userId := "01J91TY3GBN9Z1GR47PSXYKX7H"

	// mocked trip Id
	tripId := "01J91TYX985Q0DSPWMXTFH3ZM1"

	// mocked trip from date
	tripFrom, err := time.Parse(time.RFC3339, "2024-10-01T12:04:34-03:00")
	assert.Nil(t, err)

	// mocked trip to date
	tripTo, err := time.Parse(time.RFC3339, "2024-10-14T11:10:34-03:00")
	assert.Nil(t, err)

	// mocked trip entity to be created
	tripToCreate := getTripFromCFO2JFK(userId, tripFrom, tripTo)

	// logger instance needed as dependency in the use case.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// user service instance needed as dependency in the use case.
	userSrv := userMocks.Service{&userMocks.ServiceBase{

		// Given a user frequent flyer with plan BRONZE and accumulated miles 1500
		Read_: func(ctx context.IDataMap, id string) (*user.UserEntity, error) {
			nowMillis := time.Now().UnixMilli()
			userModel := user.New()
			userModel.SetId(userId)
			userModel.SetCreated(nowMillis)
			userModel.SetUpdated(nowMillis)
			userModel.SetPlan("BRONZE")
			userModel.SetMiles(1500)
			userModel.SetEmail("some@email.com")
			userModel.SetName("Jon Doe")
			userModel.SetPassword("super-secure")

			return userModel, nil
		},

		Update_: func(ctx context.IDataMap, ety *user.UserEntity) (*user.UserEntity, error) {
			//	Then the accumulated miles are 4500
			assert.EqualValues(t, 4500, ety.Miles())

			//	Then the user plan is upgraded to SILVER
			assert.Equal(t, "SILVER", ety.Plan())
			return ety, nil
		},
	}}

	// trip service instance needed as dependency in the use case.
	tripSrv := tripMocks.Service{&tripMocks.ServiceBase{

		// Then the trip is booked
		Create_: func(ctx context.IDataMap, ety *trip.TripEntity) (*trip.TripEntity, error) {
			nowMillis := time.Now().UnixMilli()
			created := trip.New()
			created.SetId(tripId)
			created.SetCreated(nowMillis)
			created.SetUpdated(nowMillis)
			created.SetMiles(ety.Miles())
			created.SetFrom(ety.From())
			created.SetTo(ety.To())
			created.SetDeparture(ety.Departure())
			created.SetArrival(ety.Arrival())
			created.SetUserId(ety.UserId())

			return created, nil
		},
	}}

	// When the user books a trip from the airport CFO  to the airport JFK and the trip distance is 2700 miles
	useCase := NewUseCase(logger, userSrv, tripSrv)
	created, err := useCase.DoBooking(tripToCreate)

	assert.Nil(t, err)
	assert.Equal(t, created.Id(), tripId)
}

// getTripFromCFO2JFK returns a tripEntity instance with mocked data.
func getTripFromCFO2JFK(userId string, from time.Time, to time.Time) *trip.TripEntity {
	tripModel := trip.New()
	tripModel.SetMiles(2700)
	tripModel.SetFrom(from.UnixMilli())
	tripModel.SetTo(to.UnixMilli())
	tripModel.SetDeparture("CFO")
	tripModel.SetArrival("JFK")
	tripModel.SetUserId(userId)

	return tripModel
}
