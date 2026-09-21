// Package racing(Service)
package racing

import (
	racingModel "tmp-app-backend/model/racing"
)

func GetRaceCourses() []racingModel.Racecourse {
	raceCourses := []racingModel.Racecourse{
		{
			ID:   1,
			Code: "TKY",
			Name: "東京競馬場",
		},
		{
			ID:   2,
			Code: "NAK",
			Name: "中山競馬場",
		},
	}

	return raceCourses
}
