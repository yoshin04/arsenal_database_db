package domain

import (
	card "app/domain/card"
	tactical "app/domain/tactical_card"
)

type GameDeck struct {
	Id            string                // デッキID
	UserId        string                // ユーザID
	MsCard1       card.MsCard           // MSカード1
	MsCard2       card.MsCard           // MSカード2
	MsCard3       card.MsCard           // MSカード3
	MsCard4       card.MsCard           // MSカード4
	PlCard1       card.PlCard           // PLカード1
	PlCard2       card.PlCard           // PLカード2
	PlCard3       card.PlCard           // PLカード3
	PlCard4       card.PlCard           // PLカード4
	TacticalCard1 tactical.TacticalCard // TCカード1
	TacticalCard2 tactical.TacticalCard // TCカード2
}
