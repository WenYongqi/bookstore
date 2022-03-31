package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestAddCart(t *testing.T) {
	book := &model.Book{
		ID: 1,
		Price: 27.20,
	}
	book2 := &model.Book{
		ID: 2,
		Price: 23.00,
	}
	cartItem := &model.WrappedCartItem{
		Book: book,
		Count: 10,
		CartID: "666888",
	}
	cartItem2 := &model.WrappedCartItem{
		Book: book2,
		Count: 10,
		CartID: "666888",
	}
	cart := &model.WrappedCart{
		ID: "666888",
		CartItems: []*model.WrappedCartItem{cartItem, cartItem2},
		UserID: 1,
	}
	err := AddCart(cart)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetCartByUserID(t *testing.T) {
	cart, err := GetCartByUserID(1)
	fmt.Printf("%+v\n", cart)
	if cart != nil {
		for _, cartItem := range cart.CartItems {
			fmt.Printf("%+v\n", cartItem)
		}
	}
	fmt.Printf("%+v\n", err)
}