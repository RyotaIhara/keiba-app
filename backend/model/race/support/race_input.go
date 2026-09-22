// Package support
package support

import (
	"time"

	"keiba-app-backend/model/race/types"
)

type RaceInput struct {
	RaceDate       time.Time
	RaceCourseID   int64
	RaceNumber     int
	RaceName       string
	StartTime      time.Time
	Surface        types.Surface
	Distance       int
	Direction      types.Direction
	Weather        types.Weather
	TrackCondition types.TrackCondition
	RaceConditions string
}

type RaceSearchInput struct {
	RaceDate     *time.Time
	RaceCourseID *int64
}
