package domain

type LinkAbility struct {
	ID                string // リンク能力ID
	Name              string // リンク能力名
	RequiredCardCount uint8  // 必要カード枚数
	Effect            string // リンク能力詳細
}

var CardEffects = map[string]string{
	"MeleeAttackSmallUp":             "MELEE_ATTACK_SMALL_UP",               //近接攻撃力小アップ
	"MeleeAttackMediumUp":            "MELEE_ATTACK_MEDIUM_UP",              //近距離攻撃力中アップ
	"MeleeAttackAndHpSmallUp":        "MELEE_ATTACK_AND_HP_SMALL_UP",        //近距離攻撃力、HP小アップ
	"RangedAttackSmallUp":            "RANGED_ATTACK_SMALL_UP",              //遠距離攻撃力小アップ
	"RangedAttackMediumUp":           "RANGED_ATTACK_MEDIUM_UP",             //遠距離攻撃力中アップ
	"MeleeAndRangedAttackSmallUp":    "MELEE_AND_RANGED_ATTACK_SMALL_UP",    //近遠攻撃力小アップ
	"RangedAttackAndHpSmallUp":       "RANGED_ATTACK_AND_HP_SMALL_UP",       //遠距離攻撃力、HP小アップ
	"MobilitySmallUp":                "MOBILITY_SMALL_UP",                   // 機動力小アップ
	"MobilityAndRangedAttackSmallUp": "MOBILITY_AND_RANGED_ATTACK_SMALL_UP", // 機動力、遠距離攻撃力小アップ
	"MobilityAndMeleeAttackSmallUp":  "MOBILITY_AND_MELEE_ATTACK_SMALL_UP",  // 機動力、近距離攻撃力小アップ
	"HpSmallUp":                      "HP_SMALL_UP",                         // HP小アップ
	"MobilityAndHpSmallUp":           "MOBILITY_AND_HP_SMALL_UP",            // 機動力、HP小アップ
	"TransformationTactic":           "TRANS_FORMATION_TACTIC",              // 変形戦術
}

func CsvTextConvertToLinkAbilityEffects(csvText string) string {
	switch csvText {
	case "近攻":
		return CardEffects["MeleeAttackSmallUp"]
	case "近攻[中↑]":
		return CardEffects["MeleeAttackMediumUp"]
	case "遠攻":
		return CardEffects["RangedAttackSmallUp"]
	case "遠攻[中↑]":
		return CardEffects["RangedAttackMediumUp"]
	case "遠/近", "遠攻/近攻", "遠攻/近攻,作戦":
		return CardEffects["MeleeAndRangedAttackSmallUp"]
	case "近攻/HP":
		return CardEffects["MeleeAttackAndHpSmallUp"]
	case "機動":
		return CardEffects["MobilitySmallUp"]
	case "機動/遠攻":
		return CardEffects["MobilityAndRangedAttackSmallUp"]
	case "機動/近攻":
		return CardEffects["MobilityAndMeleeAttackSmallUp"]
	case "HP":
		return CardEffects["HpSmallUp"]
	case "機動/HP":
		return CardEffects["MobilityAndHpSmallUp"]
	case "遠攻/HP":
		return CardEffects["RangedAttackAndHpSmallUp"]
	case "[変形戦術]":
		return CardEffects["TransformationTactic"]
	case "[秘められた力]":
		return ""
	default:
		return ""
	}
}
