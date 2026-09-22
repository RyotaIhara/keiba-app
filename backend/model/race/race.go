// Package race (Model)
package race

import (
	"time"

	"keiba-app-backend/model/race/types"
	raceCourseModel "keiba-app-backend/model/race_course"
)

type Race struct {
	ID             int64                       // レースID
	RaceDate       time.Time                   // 開催日
	Racecourse     *raceCourseModel.Racecourse // 開催競馬場
	RaceNumber     int                         // レース番号
	RaceName       string                      // レース名
	StartTime      time.Time                   // 発走時刻
	Surface        types.Surface               // 馬場
	Distance       int                         // 距離（メートル）
	Direction      types.Direction             // コースの回り方向
	Weather        types.Weather               // 天候
	TrackCondition types.TrackCondition        // 馬場状態
	RaceConditions string                      // 出走条件
}
