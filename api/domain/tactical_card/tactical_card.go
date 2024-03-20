package domainn

type TacticalCard struct {
	Id             string // カードID
	ImageUrl       string // カード画像URL
	Name           string // カード名
	Cost           uint8  // コスト
	Detail         string // カード詳細
	UnlockCriteria string // 解放条件
}
