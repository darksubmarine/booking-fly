// Package booking_fly Fly reservation use case

package booking_fly

import (
	"github.com/darksubmarine/torpedo-lib-go/context"
	"github.com/darksubmarine/torpedo-lib-go/log"

	"github.com/darksubmarine/booking-fly/domain/entities/user"

	"github.com/darksubmarine/booking-fly/domain/entities/trip"
)

// UseCase struct that implements the use case business logic.
type UseCase struct {
	*UseCaseBase

	/*  Put here your custom use case attributes */
}

// NewUseCase creates a new instance.
func NewUseCase(logger log.ILogger, userSrv user.IService, tripSrv trip.IService) *UseCase {

	return &UseCase{
		UseCaseBase: NewUseCaseBase(logger,
			userSrv, tripSrv,
		)}
}

// DoBooking method that performs a trip reservation updating user plan and miles.
func (uc *UseCase) DoBooking(tripModel *trip.TripEntity) (*trip.TripEntity, error) {

	//1. creates the trip data
	tripCreated, err := uc.tripSrv.Create(context.NoOpDataMap, tripModel)
	if err != nil {
		uc.logger.Error("something was wrong at booking creation", "error", err)
		return nil, err
	}

	//2. fetch the user linked to the trip
	userModel, err := uc.userSrv.Read(context.NoOpDataMap, tripCreated.UserId())
	if err != nil {
		uc.logger.Error("something was wrong getting the booking user", "error", err)
		return nil, ErrUserNotFound
	}

	//3. prepare trip award information
	tripMiles := tripCreated.Miles()
	accumulatedMiles := userModel.Miles() + tripMiles
	awardMiles := accumulatedMiles
	userPlan := userModel.Plan()

	switch userPlan {
	case "GOLD":
		if accumulatedMiles > 15000 && tripMiles > 2000 {
			awardMiles += 1000
		} else {
			awardMiles += 200
		}
	case "SILVER":
		if accumulatedMiles > 8000 {
			awardMiles += 500
			userPlan = "GOLD"
		}
	case "BRONZE":
		if accumulatedMiles > 4000 {
			awardMiles += 300
			userPlan = "SILVER"
		}
	}

	userModel.SetMiles(awardMiles)
	userModel.SetPlan(userPlan)

	//4. save user with updated award information based on the previous rules.
	if _, err := uc.userSrv.Update(context.NoOpDataMap, userModel); err != nil {
		uc.logger.Error("something was wrong updating the user", "error", err)
		return nil, err
	}

	return tripCreated, nil
}
