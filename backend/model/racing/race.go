// Package racing(Model)
package racing

import (
	"time"

	"tmp-app-backend/model/racing/types/race"
)

type Race struct {
	ID             int64
	RaceDate       time.Time
	Racecourse     *Racecourse
	RaceNumber     int
	RaceName       string
	StartTime      time.Time
	Surface        race.Surface
	Distance       int
	Direction      race.Direction
	Weather        race.Weather
	TrackCondition race.TrackCondition
	RaceConditions string
}
