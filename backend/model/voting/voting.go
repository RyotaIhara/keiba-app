// Package voting (Model)
package voting

import (
	raceModel "keiba-app-backend/model/race"
	userModel "keiba-app-backend/model/user"
	"keiba-app-backend/model/voting/types"
)

type Voting struct {
	ID             int64                // 投票ID
	User           *userModel.User      // ユーザー
	Race           *raceModel.Race      // レース
	PurchaseMethod types.PurchaseMethod // 購入方法
	BettingMethod  types.BetMethod      // 賭け方
}
