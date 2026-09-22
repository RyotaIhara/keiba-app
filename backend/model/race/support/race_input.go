// Package support
package support

import (
	"time"

	raceTypes "keiba-app-backend/model/race/types/race"
)

type RaceInput struct {
	RaceDate       time.Time
	RaceCourseID   int64
	RaceNumber     int
	RaceName       string
	StartTime      time.Time
	Surface        raceTypes.Surface
	Distance       int
	Direction      raceTypes.Direction
	Weather        raceTypes.Weather
	TrackCondition raceTypes.TrackCondition
	RaceConditions string
}
