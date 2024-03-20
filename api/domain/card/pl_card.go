package domain

type PlCard struct {
	No               string      // カードNo
	ImageUrl         string      // イメージURL
	IncludeCode      string      // 収録コード
	Rarity           string      // PLカードレアリティ
	Name             string      // カード名
	Type             string      // タイプ
	Cost             uint8       // コスト
	Mobility         uint16      // 機動力
	LongRangeAttack  uint16      // 遠距離攻撃力
	CloseRangeAttack uint16      // 近距離攻撃力
	Hp               uint16      // HP
	TotalScore       uint16      // 総合スコア
	PlSkill          PlSkill     // PLスキル
	LinkAbility1     LinkAbility // リンクアビリティ1
	LinkAbility2     LinkAbility // リンクアビリティ2
	SeriesTitle      string      // 出典作品
}
