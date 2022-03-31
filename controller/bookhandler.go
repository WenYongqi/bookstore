package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pageNo_s := r.FormValue("pageNo")
	pageNo, _ := strconv.ParseInt(pageNo_s, 10, 64)
	if pageNo == 0 {
		pageNo = 1
	}
	page, err := dao.GetPageBooks(pageNo)
	if err != nil {
		log.Println(err)
	}
	isLogin, sess := dao.IsLogin(r)
	if isLogin {
		page.IsLogin = true
		page.Username = sess.Username
	}
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}

//func GetBooks(w http.ResponseWriter, r *http.Request) {
//	books, err := dao.GetBooks()
//	if err != nil {
//		log.Println(err)
//	}
//	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
//	t.Execute(w, books)
//}

func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	pageNo_s := r.FormValue("pageNo")
	pageNo, _ := strconv.ParseInt(pageNo_s, 10, 64)
	if pageNo == 0 {
		pageNo = 1
	}
	page, err := dao.GetPageBooks(pageNo)
	if err != nil {
		log.Println(err)
	}
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}

func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	pageNo_s := r.FormValue("pageNo")
	minPrice_s := r.FormValue("minPrice")
	maxPrice_s := r.FormValue("maxPrice")
	pageNo, _ := strconv.ParseInt(pageNo_s, 10, 64)
	if pageNo == 0 {
		pageNo = 1
	}
	//不带价格全局查询
	if minPrice_s == "" && maxPrice_s == "" {
		IndexHandler(w, r)
		return
	}
	minPrice, err1 := strconv.ParseFloat(minPrice_s, 10)
	maxPrice, err2 := strconv.ParseFloat(maxPrice_s, 10)
	//价格参数不合法
	if err1 != nil || err2 != nil {
		log.Println("price parameter not legal")
		IndexHandler(w, r)
		return
	}
	page, err := dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
	if err != nil {
		log.Println(err)
	}
	isLogin, sess := dao.IsLogin(r)
	if isLogin {
		page.IsLogin = true
		page.Username = sess.Username
	}
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	if r.PostFormValue("title") == "" {
		log.Println("书名不能为空")
		GetPageBooks(w, r)
		return
	}
	price, err := strconv.ParseFloat(r.PostFormValue("price"), 64)
	if err != nil {
		log.Println("解析数字失败")
	}
	sales, err := strconv.Atoi(r.PostFormValue("sales"))
	if err != nil {
		log.Println("解析数字失败")
	}
	stock, err := strconv.Atoi(r.PostFormValue("stock"))
	if err != nil {
		log.Println("解析数字失败")
	}
	book := model.Book{
		Title: r.PostFormValue("title"),
		Author: r.PostFormValue("author"),
		Price: price,
		Sales: sales,
		Stock: stock,
		ImgPath: "static/img/default.jpg",
	}
	err = dao.AddBook(&book)
	if err != nil {
		log.Println(err)
	}
	GetPageBooks(w, r)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID_s := r.FormValue("bookID")
	bookID, err := strconv.Atoi(bookID_s)
	if err != nil {
		log.Println("解析失败")
		GetPageBooks(w, r)
		return
	}
	err = dao.DeleteBook(bookID)
	if err != nil {
		log.Println(err)
	}
	GetPageBooks(w, r)
}

func ToUpdateOrAddBookPage(w http.ResponseWriter, r *http.Request) {
	bookID_s := r.FormValue("bookID")
	//Add
	if bookID_s == "" {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w,"")
		return
	}
	//Update
	bookID, err := strconv.Atoi(bookID_s)
	if err != nil {
		log.Println("解析失败")
		GetPageBooks(w, r)
		return
	}
	book, err := dao.GetBookByID(bookID)
	if err != nil {
		log.Println(err)
		GetPageBooks(w, r)
		return
	}
	t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
	t.Execute(w, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := strconv.Atoi(r.PostFormValue("bookID"))
	if err != nil {
		log.Println("bookID 不合法")
		GetPageBooks(w, r)
		return
	}
	if r.PostFormValue("title") == "" {
		log.Println("书名不能为空")
		GetPageBooks(w, r)
		return
	}
	price, err := strconv.ParseFloat(r.PostFormValue("price"), 64)
	if err != nil {
		log.Println("解析数字失败")
	}
	sales, err := strconv.Atoi(r.PostFormValue("sales"))
	if err != nil {
		log.Println("解析数字失败")
	}
	stock, err := strconv.Atoi(r.PostFormValue("stock"))
	if err != nil {
		log.Println("解析数字失败")
	}
	book := model.Book{
		ID: bookID,
		Title: r.PostFormValue("title"),
		Author: r.PostFormValue("author"),
		Price: price,
		Sales: sales,
		Stock: stock,
	}
	err = dao.UpdateBook(&book)
	if err != nil {
		log.Println(err)
	}
	GetPageBooks(w, r)
}