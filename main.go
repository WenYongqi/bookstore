package main

import (
	"bookstore/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/main", controller.IndexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.FileServer(http.Dir("views")))

	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/checkUsername", controller.CheckUsername)

	//http.HandleFunc("/getBooks", controller.GetBooks)
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	http.HandleFunc("/addBook", controller.AddBook)
	http.HandleFunc("/updateBook", controller.UpdateBook)
	http.HandleFunc("/toUpdateOrAddBookPage", controller.ToUpdateOrAddBookPage)

	http.HandleFunc("/addBookToCart", controller.AddBookToCart)
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	http.HandleFunc("/emptyCart", controller.EmptyCart)
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)

	http.HandleFunc("/checkout", controller.Checkout)
	http.HandleFunc("/getOrders", controller.GetOrders)
	http.HandleFunc("/getMyOrders", controller.GetMyOrders)
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)
	http.HandleFunc("/sendOrder", controller.SendOrder)
	http.HandleFunc("/confirmReceipt", controller.ConfirmReceipt)
	http.ListenAndServe(":8080", nil)
}
