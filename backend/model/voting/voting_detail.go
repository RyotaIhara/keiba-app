// Package voting (Model)
package voting

import "keiba-app-backend/model/voting/types"

type VotingDetail struct {
	ID            int64               // 投票詳細ID
	Voting        *Voting             // 投票
	BettingTicket types.BettingTicket // 買い目
	Amount        int                 // 投票金額
}
