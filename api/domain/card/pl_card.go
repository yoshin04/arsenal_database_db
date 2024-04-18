package domain

import "app/db/models"

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

func ToDomainPlCard(m *models.PLCard) *PlCard {
	if m == nil {
		return nil
	}
	plSkill := PlSkill{
		Name:      m.PlSkillName,
		Condition: m.PlSkillCondition,
		Detail:    m.PlSkillDetail,
	}

	var linkAbility1, linkAbility2 LinkAbility
	if m.FirstLinkAbility != nil {
		linkAbility1 = LinkAbility{
			Name:              m.FirstLinkAbility.Name,
			RequiredCardCount: m.FirstLinkAbility.RequiredCardCount,
			Effect:            m.FirstLinkAbility.Effect,
		}
	}
	if m.SecondLinkAbility != nil {
		linkAbility2 = LinkAbility{
			Name:              m.SecondLinkAbility.Name,
			RequiredCardCount: m.SecondLinkAbility.RequiredCardCount,
			Effect:            m.SecondLinkAbility.Effect,
		}
	}

	return &PlCard{
		No:               m.No,
		IncludeCode:      m.IncludeCode,
		ImageUrl:         m.ImageURL,
		Rarity:           m.Rarity,
		Name:             m.Name,
		Type:             m.Type,
		Cost:             m.Cost,
		Mobility:         m.Mobility,
		LongRangeAttack:  m.LongRangeAttack,
		CloseRangeAttack: m.CloseRangeAttack,
		Hp:               m.HP,
		TotalScore:       m.TotalScore,
		PlSkill:          plSkill,
		LinkAbility1:     linkAbility1,
		LinkAbility2:     linkAbility2,
		SeriesTitle:      m.SeriesTitle,
	}
}
