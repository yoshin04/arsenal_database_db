package domain

type MSCard struct {
	No                    string      // MSカードNo
	IncludedCode          string      // 収録コード
	Rarity                string      // レアリティ
	Name                  string      // 名
	Type                  string      // タイプ
	Cost                  uint8       // コスト
	Mobility              uint16      // 機動力
	LongRangeAttack       uint16      // 遠距離攻撃力
	CloseRangeAttack      uint16      // 近距離攻撃力
	Hp                    uint16      // HP
	TotalScore            uint16      // 総合スコア
	MainRange             uint8       // メイン射程
	SubRange              uint8       // サブ射程
	GroundSuitability     string      // 地上適正
	SpaceSuitability      string      // 宇宙適正
	DesertSuitability     string      // 砂漠適正
	UnderwaterSuitability string      // 水中適正
	SpAttack              SpAttack    // 特殊攻撃
	MsAbility             MsAbility   // MSアビリティ
	LinkAbility           LinkAbility // リンクアビリティ
	SeriesTitle           string      // 出典作品
}
