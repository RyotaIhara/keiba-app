// Package racing(Service)
package racing

import (
	"time"

	racingModel "tmp-app-backend/model/racing"
	raceType "tmp-app-backend/model/racing/types/race"
)

func GetRaces() []racingModel.Race {
	races := []racingModel.Race{
		{
			ID:             1,
			RaceDate:       time.Date(2026, time.May, 24, 0, 0, 0, 0, time.Local),
			Racecourse:     &racingModel.Racecourse{ID: 1, Code: "TKY", Name: "東京競馬場"},
			RaceNumber:     11,
			RaceName:       "優駿牝馬",
			StartTime:      time.Date(2026, time.May, 24, 15, 40, 0, 0, time.Local),
			Surface:        raceType.SurfaceTurf,
			Distance:       2400,
			Direction:      raceType.DirectionLeft,
			Weather:        raceType.WeatherSunny,
			TrackCondition: raceType.TrackConditionGood,
			RaceConditions: "3歳牝馬",
		},
		{
			ID:             2,
			RaceDate:       time.Date(2026, time.May, 24, 0, 0, 0, 0, time.Local),
			Racecourse:     &racingModel.Racecourse{ID: 2, Code: "NAK", Name: "中山競馬場"},
			RaceNumber:     10,
			RaceName:       "サンプルレース",
			StartTime:      time.Date(2026, time.May, 24, 15, 10, 0, 0, time.Local),
			Surface:        raceType.SurfaceDirt,
			Distance:       1800,
			Direction:      raceType.DirectionRight,
			Weather:        raceType.WeatherCloudy,
			TrackCondition: raceType.TrackConditionFirm,
			RaceConditions: "4歳以上",
		},
	}

	return races
}
