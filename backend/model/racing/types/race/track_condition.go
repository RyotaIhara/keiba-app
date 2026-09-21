// Package race
package race

type TrackCondition string

const (
	TrackConditionFirm  TrackCondition = "firm"
	TrackConditionGood  TrackCondition = "good"
	TrackConditionYield TrackCondition = "yield"
	TrackConditionSoft  TrackCondition = "soft"
)
