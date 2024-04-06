package domain

// TargetType - SP攻撃対象タイプ
var TargetType = map[string]string{
	"SingleTarget": "SINGLE", // 単体
	"PierceTarget": "PIERCE", // 貫通
	"AreaOfEffect": "AOE",    // 範囲
}

func CsvTextConvertToSpTargetType(csvText string) string {
	switch csvText {
	case "単体(敵)":
		return TargetType["SingleTarget"]
	case "範囲(敵)":
		return TargetType["AreaOfEffect"]
	case "貫通(敵)":
		return TargetType["PierceTarget"]
	default:
		return ""
	}
}
