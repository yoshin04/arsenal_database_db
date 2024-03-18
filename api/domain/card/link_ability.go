package domain

type LinkAbility struct {
	Name              string // リンク能力名
	RequiredCardCount uint8  // 必要カード枚数
	Effect            string // リンク能力詳細
}

var CardEffects = map[string]string{
	"MeleeAttackSmallUp":   "MELEE_ATTACK_SMALL_UP",   //近接攻撃力小アップ
	"MeleeAttackMediumUp":  "MELEE_ATTACK_MEDIUM_UP",  //近距離攻撃力中アップ
	"RangedAttackSmallUp":  "RANGED_ATTACK_SMALL_UP",  //遠距離攻撃力小アップ
	"RangedAttackMediumUp": "RANGED_ATTACK_MEDIUM_UP", //遠距離攻撃力中アップ
	"MobilitySmallUp":      "MOBILITY_SMALL_UP",       // 機動力小アップ
	"HpSmallUp":            "HP_SMALL_UP",             // HP小アップ
	"TransformationTactic": "TRANS_FORMATION_TACTIC",  // 変形戦術
}
