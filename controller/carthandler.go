package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func AddBookToCart(w http.ResponseWriter, r *http.Request) {
	isLogin, sess := dao.IsLogin(r)
	userID := 0
	if isLogin {
		userID = sess.UserID
	} else {
		w.Write([]byte("请先登录！"))
		return
	}
	bookID_s := r.FormValue("bookID")
	bookID,err := strconv.Atoi(bookID_s)
	if err != nil {
		log.Println(err)
	}
	book, err := dao.GetBookByID(bookID)
	if err != nil {
		log.Println(err)
	}
	cart, err := dao.GetCartByUserID(userID)
	if err != nil {
		log.Println(err)
	}
	if cart != nil {
		cartItem, err := dao.GetCartItemByBookIDAndCartID(bookID_s, cart.ID)
		if err != nil {
			log.Println(err)
		}
		if cartItem != nil {
			for _, item := range cart.CartItems {
				if item.ID == cartItem.ID {
					item.Count += 1
					err := dao.UpdateCartItem(item)
					if err != nil {
						log.Println(err)
					}
					break
				}
			}
		} else {
			cartItem := &model.WrappedCartItem{
				Book: book,
				Count: 1,
				CartID: cart.ID,
			}
			cart.CartItems = append(cart.CartItems, cartItem)
			dao.AddCartItem(cartItem)
		}
		err = dao.UpdateCart(cart)
		if err != nil {
			log.Println(err)
		}
	} else {
		cartID := utils.CreateUUID()
		cart := model.WrappedCart{
			ID: cartID,
			UserID: userID,
		}
		cartItems := []*model.WrappedCartItem{
			&model.WrappedCartItem{
				Book: book,
				Count: 1,
				CartID: cartID,
			},
		}
		cart.CartItems = cartItems
		err := dao.AddCart(&cart)
		if err != nil {
			log.Println(err)
		}
	}
	w.Write([]byte("您刚刚将《" + book.Title + "》加入到了购物车中"))
}

func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	isLogin, sess := dao.IsLogin(r)
	userID := 0
	if isLogin {
		userID = sess.UserID
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
		return
	}
	cart, err := dao.GetCartByUserID(userID)
	if err != nil {
		log.Println(err)
	}
	if cart == nil {
		cart = &model.WrappedCart{}
	}
	cart.Username = sess.Username
	t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
	t.Execute(w, cart)
}

func EmptyCart(w http.ResponseWriter, r *http.Request) {
	cartID := r.FormValue("cartID")
	err := dao.DeleteCartItemsByCartID(cartID)
	if err != nil {
		log.Println(err)
	}
	err = dao.DeleteCartByCartID(cartID)
	if err != nil {
		log.Println(err)
	}
	GetCartInfo(w, r)
}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID_s := r.FormValue("cartItemID")
	cartItemID, err := strconv.Atoi(cartItemID_s)
	if err != nil {
		log.Println(err)
	}
	_, sess := dao.IsLogin(r)
	userID := sess.UserID
	cart, err := dao.GetCartByUserID(userID)
	if err != nil {
		log.Println(err)
	}
	newCartItems := []*model.WrappedCartItem{}
	if cart != nil {
		for _, cartItem := range cart.CartItems {
			if cartItem.ID == cartItemID {
				if cartItem.Count == 1 {
					err := dao.DeleteCartItemByCartItemID(cartItemID)
					if err != nil {
						log.Println(err)
					}
				} else {
					cartItem.Count -= 1
					err := dao.UpdateCartItem(cartItem)
					if err != nil {
						log.Println(err)
					}
					newCartItems = append(newCartItems, cartItem)
				}
			} else {
				newCartItems = append(newCartItems, cartItem)
			}
		}
		cart.CartItems = newCartItems
		dao.UpdateCart(cart)
	}
	GetCartInfo(w, r)
}