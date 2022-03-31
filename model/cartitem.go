package model

type WrappedCartItem struct {
	ID int
	Book *Book
	Count int64 //数量
	Amount float64 //总金额
	CartID string //购物项属于哪个购物车
}

type CartItem struct {
	ID int
	BookID int
	Count int64 //数量
	Amount float64 //总金额
	CartID string //购物项属于哪个购物车
}

func (wrappedCartItem *WrappedCartItem) FillAmount() {
	 wrappedCartItem.Amount = float64(wrappedCartItem.Count) * wrappedCartItem.Book.Price
}