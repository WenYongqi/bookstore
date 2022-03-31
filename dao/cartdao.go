package dao

import (
	"bookstore/model"
	"gorm.io/gorm"
	"log"
)

func WrapCart(cart *model.Cart) *model.WrappedCart {
	return &model.WrappedCart{
		ID: cart.ID,
		TotalCount: cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		UserID: cart.UserID,
	}
}

func UnwrapCart(wrappedCart *model.WrappedCart) *model.Cart {
	return &model.Cart{
		ID: wrappedCart.ID,
		TotalCount: wrappedCart.TotalCount,
		TotalAmount: wrappedCart.TotalAmount,
		UserID: wrappedCart.UserID,
	}
}

func AddCart(cart *model.WrappedCart) error {
	cart.FillCountAndAmount()
	err := db.Create(UnwrapCart(cart)).Error
	if err != nil {
		log.Println(err)
		return err
	}
	for _,cartItem := range cart.CartItems {
		err := AddCartItem(cartItem)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func GetCartByUserID(userID int) (*model.WrappedCart, error) {
	cart := model.Cart{}
	err := db.Where("user_id = ?", userID).First(&cart).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	wrappedCart := *WrapCart(&cart)
	cartItems, err := GetCartItemsByCartID(cart.ID)
	if err != nil {
		log.Println(err)
	}
	wrappedCartItems := []*model.WrappedCartItem{}
	for _, cartItem := range cartItems {
		wrappedCartItems = append(wrappedCartItems, WrapCartItem(cartItem))
	}
	wrappedCart.CartItems = wrappedCartItems
	return &wrappedCart, nil
}

func UpdateCart(cart *model.WrappedCart) error {
	cart.FillCountAndAmount()
	err := db.Save(UnwrapCart(cart)).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteCartByCartID(cartID string) error {
	err := db.Delete(&model.Cart{}, "id = ?", cartID).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}