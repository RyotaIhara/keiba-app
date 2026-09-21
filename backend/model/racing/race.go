// Package racing(Model)
package racing

import (
	"time"

	"tmp-app-backend/model/racing/types/race"
)

type Race struct {
	ID             int64               // レースID
	RaceDate       time.Time           // 開催日
	Racecourse     *Racecourse         // 開催競馬場
	RaceNumber     int                 // レース番号
	RaceName       string              // レース名
	StartTime      time.Time           // 発走時刻
	Surface        race.Surface        // 馬場
	Distance       int                 // 距離（メートル）
	Direction      race.Direction      // コースの回り方向
	Weather        race.Weather        // 天候
	TrackCondition race.TrackCondition // 馬場状態
	RaceConditions string              // 出走条件
}
