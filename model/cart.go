package model

type WrappedCart struct {
	ID string
	CartItems []*WrappedCartItem
	TotalCount int64
	TotalAmount float64
	UserID int
	Username string
}

type Cart struct {
	ID string
	TotalCount int64
	TotalAmount float64
	UserID int
}

func (wrappedCart *WrappedCart) FillCountAndAmount() {
	wrappedCart.TotalCount = 0
	for _, item := range wrappedCart.CartItems {
		wrappedCart.TotalCount += item.Count
	}
	wrappedCart.TotalAmount = 0
	for _, item := range wrappedCart.CartItems {
		item.FillAmount()
		wrappedCart.TotalAmount += item.Amount
	}
}