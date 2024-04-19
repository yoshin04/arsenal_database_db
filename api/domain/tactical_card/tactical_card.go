package domainn

import "app/db/models"

type TacticalCard struct {
	Id             string // カードID
	ImageUrl       string // カード画像URL
	Name           string // カード名
	Cost           uint8  // コスト
	Detail         string // カード詳細
	UnlockCriteria string // 解放条件
}

func ToDomainTacticalCard(m *models.TacticalCard) *TacticalCard {
	if m == nil {
		return nil
	}

	return &TacticalCard{
		Id:             m.ID,
		ImageUrl:       m.ImageUrl,
		Name:           m.Name,
		Cost:           m.Cost,
		Detail:         m.Detail,
		UnlockCriteria: m.UnlockCriteria,
	}
}
