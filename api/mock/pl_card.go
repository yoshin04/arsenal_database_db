package mock

import domain "app/domain/card"

func GenerateDummyPlCards() []*domain.PlCard {
	return []*domain.PlCard{
		{
			Id:               "AB01051",
			No:               "051",
			ImageUrl:         "https://example.com/image.jpg",
			IncludeCode:      "AB01",
			Rarity:           domain.Rarity["M"],
			Name:             "アムロ・レイ",
			Type:             domain.CardType["CONTROL"],
			Cost:             4,
			Mobility:         150,
			LongRangeAttack:  200,
			CloseRangeAttack: 240,
			Hp:               160,
			TotalScore:       750,
			PlSkill: domain.PlSkill{
				Name:      "決定的な一撃",
				Condition: "敵戦艦/拠点をロックオン時",
				Detail:    "敵戦艦/拠点へのダメージを中アップする。",
			},
			LinkAbility1: domain.LinkAbility{
				Name:              "機動戦士ガンダム",
				RequiredCardCount: 3,
				Effect:            domain.CardEffects["MobilitySmallUp"],
			},
			LinkAbility2: domain.LinkAbility{
				Name:              "ニュータイプの潜在能力",
				RequiredCardCount: 3,
				Effect:            domain.CardEffects["MobilitySmallUp"],
			},
			SeriesTitle: "機動戦士ガンダム",
		},
	}
}
