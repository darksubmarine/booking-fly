package fixtures

import (
	"github.com/darksubmarine/booking-fly/domain/entities/trip"
	"time"
)

const (
	TripId = "01J91TYX985Q0DSPWMXTFH3ZM1"

	TripDeparture = "CFO"
	TripArrival   = "JFK"
	TripMiles     = 2700
	TripFrom      = 1726837954000
	TripTo        = 1727529154000
)

// Trip returns a trip.TripEntity with mocked data and the given userId
func Trip(userId string) *trip.TripEntity {

	nowMillis := time.Now().UnixMilli()

	tripModel := trip.New()
	tripModel.SetId(TripId)
	tripModel.SetCreated(nowMillis)
	tripModel.SetUpdated(nowMillis)
	tripModel.SetMiles(TripMiles)
	tripModel.SetFrom(TripFrom)
	tripModel.SetTo(TripTo)
	tripModel.SetDeparture(TripDeparture)
	tripModel.SetArrival(TripArrival)
	tripModel.SetUserId(userId)

	return tripModel
}
