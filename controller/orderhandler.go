package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"log"
	"net/http"
	"text/template"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
	_, sess := dao.IsLogin(r)
	userID := sess.UserID
	cart, err := dao.GetCartByUserID(userID)
	if err != nil {
		log.Println(err)
	}
	order := &model.Order{
		ID: utils.CreateUUID(),
		TotalCount: cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State: 0,
		UserID: userID,
		Username: sess.Username,
	}
	err = dao.AddOrder(order)
	if err != nil {
		log.Println(err)
	}
	for _, cartItem := range cart.CartItems {
		orderItem := &model.OrderItem{
			Count: cartItem.Count,
			Amount: cartItem.Amount,
			Title: cartItem.Book.Title,
			Author: cartItem.Book.Author,
			Price: cartItem.Book.Price,
			ImgPath: cartItem.Book.ImgPath,
			OrderID: order.ID,
		}
		err = dao.AddOrderItem(orderItem)
		if err != nil {
			log.Println(err)
		}
		book := cartItem.Book
		book.Sales += int(cartItem.Count)
		book.Stock -= int(cartItem.Count)
		err = dao.UpdateBook(book)
		if err != nil {
			log.Println(err)
		}
	}
	cartID := cart.ID
	err = dao.DeleteCartItemsByCartID(cartID)
	if err != nil {
		log.Println(err)
	}
	err = dao.DeleteCartByCartID(cartID)
	if err != nil {
		log.Println(err)
	}
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w, order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := dao.GetOrders()
	if err != nil {
		log.Println(err)
	}
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}

func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	_, sess := dao.IsLogin(r)
	userID := sess.UserID
	orders, err := dao.GetMyOrders(userID)
	if err != nil {
		log.Println(err)
	}
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}

func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderID")
	orderItems, err := dao.GetOrderItemsByOrderID(orderID)
	if err != nil {
		log.Println(err)
	}
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w, orderItems)
}

func SendOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderID")
	err := dao.UpdateOrderState(orderID, 1)
	if err != nil {
		log.Println(err)
	}
	isLogin, _ := dao.IsLogin(r)
	if isLogin {
		GetMyOrders(w, r)
	} else {
		GetOrders(w, r)
	}
}

func ConfirmReceipt(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderID")
	err := dao.UpdateOrderState(orderID, 2)
	if err != nil {
		log.Println(err)
	}
	isLogin, _ := dao.IsLogin(r)
	if isLogin {
		GetMyOrders(w, r)
	} else {
		GetOrders(w, r)
	}
}