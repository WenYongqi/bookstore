package dao

import (
	"fmt"
	"testing"
)

func TestGetCartItemByBookIDAndCartID(t *testing.T) {
	cartItem, err := GetCartItemByBookIDAndCartID("1", "666888")
	fmt.Printf("%+v\n%v\n", cartItem, err)
}

func TestGetCartItemByCartID(t *testing.T) {
	cartItems, err := GetCartItemsByCartID("666888")
	for _, cartItem := range cartItems {
		fmt.Printf("%+v\n", cartItem)
	}
	fmt.Printf("%v\n", err)
}
