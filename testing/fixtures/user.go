package fixtures

import (
	"github.com/darksubmarine/booking-fly/domain/entities/user"
	"time"
)

const (
	UserId = "01J91TY3GBN9Z1GR47PSXYKX7H"

	UserPlanBronze = "BRONZE"
	UserPlanSilver = "SILVER"
	UserPlanGold   = "GOLD"

	UserMiles        = 1500
	UserUpdatedMiles = 4500
)

// User returns a user.UserEntity mocked data
func User() *user.UserEntity {
	nowMillis := time.Now().UnixMilli()
	userModel := user.New()
	userModel.SetId(UserId)
	userModel.SetCreated(nowMillis)
	userModel.SetUpdated(nowMillis)
	userModel.SetPlan(UserPlanBronze)
	userModel.SetMiles(UserMiles)
	userModel.SetEmail("some@email.com")
	userModel.SetName("Jon Doe")
	userModel.SetPassword("super-secure")

	return userModel
}
