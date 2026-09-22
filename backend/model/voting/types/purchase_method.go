// Package types
package types

type PurchaseMethod string

const (
	PurchaseMethodNormal    PurchaseMethod = "normal"    // 通常
	PurchaseMethodBox       PurchaseMethod = "box"       // BOX
	PurchaseMethodFormation PurchaseMethod = "formation" // フォーメーション
	PurchaseMethodNagashi   PurchaseMethod = "nagashi"   // ながし
)
