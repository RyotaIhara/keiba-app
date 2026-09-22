// Package types
package types

type BetMethod string

const (
	BetMethodWin      BetMethod = "win"      // 単勝
	BetMethodPlace    BetMethod = "place"    // 複勝
	BetMethodQuinella BetMethod = "quinella" // 馬連
	BetMethodExacta   BetMethod = "exacta"   // 馬単
	BetMethodTrio     BetMethod = "trio"     // 3連複
	BetMethodTrifecta BetMethod = "trifecta" // 3連単
)
