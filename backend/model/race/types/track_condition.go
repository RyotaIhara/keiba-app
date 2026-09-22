// Package types
package types

type TrackCondition string

const (
	TrackConditionFirm  TrackCondition = "firm"  // 良
	TrackConditionGood  TrackCondition = "good"  // 稍良
	TrackConditionYield TrackCondition = "yield" // 重
	TrackConditionSoft  TrackCondition = "soft"  // 不良
)
