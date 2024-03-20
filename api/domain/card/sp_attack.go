package domain

// SPAttackはMSカードのSP攻撃を表す構造体。
type SpAttack struct {
	Name       string // SP攻撃名
	Cost       uint16 // SP攻撃威力
	Power      uint16 // SP攻撃射程
	Range      uint8 // SP攻撃対象タイプ
	TargetType string // SP攻撃タイプ
	Detail     string // SP攻撃詳細
}
