package domain

type TacticalCard struct {
	Id             string // カードID
	Name           string // カード名
	Cost           uint8  // コスト
	Detail         string // カード詳細
	UnlockCriteria string // 解放条件
}
