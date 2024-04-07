package domain

// SPAttackはMSカードのSP攻撃を表す構造体。
type SpAttack struct {
	Name       string // SP攻撃名
	Cost       uint8  // SPコスト
	Power      uint16 // SP攻撃威力
	Range      uint8  // SP攻撃射程
	Type       string // SP攻撃タイプ
	TargetType string // SP攻撃対象タイプ
	Detail     string // SP攻撃詳細
}
