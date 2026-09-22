// Package racing(Model)
package racing

import (
	"keiba-app-backend/model/racing/types/race"
)

type RaceDetail struct {
	ID               int64    // レース詳細ID
	Race             *Race    // レース
	HorseNumber      int      // 馬番
	FrameNumber      int      // 枠番
	HorseName        string   // 馬名
	Sex              race.Sex // 性別
	Age              int      // 年齢
	Weight           int      // 斤量
	Jockey           string   // 騎手
	Stable           string   // 厩舎
	BodyWeight       int      // 馬体重
	BodyWeightChange int      // 馬体重の増減
	Odds             float64  // オッズ
	Popularity       int      // 人気
}
