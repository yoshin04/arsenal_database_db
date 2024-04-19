package domain

import (
	"app/db/models"
)

type MsCard struct {
	No                    string      // カードNo
	ImageUrl              string      // イメージURL
	IncludedCode          string      // 収録コード
	Rarity                string      // レアリティ
	Name                  string      // カード名
	Type                  string      // タイプ
	Cost                  uint8       // コスト
	Mobility              uint16      // 機動力
	LongRangeAttack       uint16      // 遠距離攻撃力
	CloseRangeAttack      uint16      // 近距離攻撃力
	Hp                    uint16      // HP
	TotalScore            uint16      // 総合スコア
	MainRange             uint16      // メイン射程
	SubRange              uint16      // サブ射程
	GroundSuitability     string      // 地上適正
	SpaceSuitability      string      // 宇宙適正
	DesertSuitability     string      // 砂漠適正
	UnderwaterSuitability string      // 水中適正
	SpAttack              SpAttack    // 特殊攻撃
	MsAbility             MsAbility   // MSアビリティ
	LinkAbility1          LinkAbility // リンクアビリティ
	LinkAbility2          LinkAbility // リンクアビリティ
	SeriesTitle           string      // 出典作品
}

func ToDomainMsCard(m *models.MSCard) *MsCard {
	if m == nil {
		return nil
	}
	msAbility := MsAbility{}
	if m.AbilityName != nil && m.AbilityCost != nil && m.AbilityRange != nil && m.AbilityDetail != nil {
		msAbility = MsAbility{
			Name:   *m.AbilityName,
			Cost:   *m.AbilityCost,
			Range:  *m.AbilityRange,
			Detail: *m.AbilityDetail,
		}
	}

	var linkAbility1, linkAbility2 LinkAbility
	if m.FirstLinkAbility != nil {
		linkAbility1 = LinkAbility{
			ID:                m.FirstLinkAbility.ID,
			Name:              m.FirstLinkAbility.Name,
			RequiredCardCount: m.FirstLinkAbility.RequiredCardCount,
			Effect:            m.FirstLinkAbility.Effect,
		}
	}
	if m.SecondLinkAbility != nil {
		linkAbility2 = LinkAbility{
			ID:                m.SecondLinkAbility.ID,
			Name:              m.SecondLinkAbility.Name,
			RequiredCardCount: m.SecondLinkAbility.RequiredCardCount,
			Effect:            m.SecondLinkAbility.Effect,
		}
	}

	return &MsCard{
		No:                    m.No,
		ImageUrl:              m.ImageURL,
		IncludedCode:          m.IncludedCode,
		Rarity:                m.Rarity,
		Name:                  m.Name,
		Type:                  m.Type,
		Cost:                  m.Cost,
		Mobility:              m.Mobility,
		LongRangeAttack:       m.LongRangeAttack,
		CloseRangeAttack:      m.CloseRangeAttack,
		Hp:                    m.HP,
		TotalScore:            m.TotalScore,
		MainRange:             m.MainRange,
		SubRange:              m.SubRange,
		GroundSuitability:     m.GroundSuitability,
		SpaceSuitability:      m.SpaceSuitability,
		DesertSuitability:     m.DesertSuitability,
		UnderwaterSuitability: m.UnderwaterSuitability,
		SpAttack: SpAttack{
			Name:       m.SpAttackName,
			Cost:       m.SpCost,
			Power:      m.SpPower,
			Range:      m.SpRange,
			Type:       m.SpType,
			TargetType: m.SpTargetType,
			Detail:     m.SpDetail,
		},
		MsAbility:    msAbility,
		LinkAbility1: linkAbility1,
		LinkAbility2: linkAbility2,
		SeriesTitle:  m.SeriesTitle,
	}
}
