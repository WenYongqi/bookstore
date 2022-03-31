package dao

import (
	"bookstore/model"
	"gorm.io/gorm"
	"log"
)

func WrapCartItem(cartItem *model.CartItem) *model.WrappedCartItem {
	book, err := GetBookByID(cartItem.BookID)
	if err != nil {
		log.Println(err)
	}
	return &model.WrappedCartItem{
		ID: cartItem.ID,
		Count: cartItem.Count,
		Amount: cartItem.Amount,
		CartID: cartItem.CartID,
		Book: book,
	}
}

func UnwrapCartItem(wrappedCartItem *model.WrappedCartItem) *model.CartItem {
	return &model.CartItem{
		ID: wrappedCartItem.ID,
		BookID: wrappedCartItem.Book.ID,
		Count: wrappedCartItem.Count,
		Amount: wrappedCartItem.Amount,
		CartID: wrappedCartItem.CartID,
	}
}

func AddCartItem(cartItem *model.WrappedCartItem) error {
	cartItem.FillAmount()
	err := db.Create(UnwrapCartItem(cartItem)).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetCartItemByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {
	cartItem := model.CartItem{}
	err := db.Where("book_id = ? and cart_id = ?", bookID, cartID).First(&cartItem).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &cartItem, nil
}

func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	cartItem := []*model.CartItem{}
	err := db.Where("cart_id = ?", cartID).Find(&cartItem).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return cartItem, nil
}

func UpdateCartItem(cartItem *model.WrappedCartItem) error {
	cartItem.FillAmount()
	err := db.Save(UnwrapCartItem(cartItem)).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteCartItemsByCartID(cartID string) error {
	err := db.Delete(&model.CartItem{}, "cart_id = ?", cartID).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteCartItemByCartItemID(cartItemID int) error {
	err := db.Delete(&model.CartItem{}, "id = ?", cartItemID).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}