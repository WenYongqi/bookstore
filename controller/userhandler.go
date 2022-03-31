package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
	isLogin, _ := dao.IsLogin(r)
	if isLogin {
		GetPageBooksByPrice(w, r)
		return
	}
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user, _ := dao.CheckUsernameAndPassword(username, password)
	if user.ID == 0 {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确")
	} else {
		session_id := utils.CreateUUID()
		sess := &model.Session{
			SessionID: session_id,
			Username: username,
			UserID: user.ID,
		}
		err := dao.AddSession(sess)
		if err != nil {
			log.Println(err)
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "登录失败")
			return
		}
		cookie := http.Cookie{
			Name: "session_id",
			Value: session_id,
		}
		http.SetCookie(w, &cookie)
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, username)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		log.Println(err)
	} else if cookie != nil {
		session_id := cookie.Value
		err := dao.DeleteSession(session_id)
		if err != nil {
			log.Println(err)
		} else {
			cookie.MaxAge = -1
			http.SetCookie(w, cookie)
		}
	}
	http.Redirect(w, r, "/main", 302)
}

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	if username == "" || password == "" {
		t := template.Must(template.ParseFiles("views/pages/user/register.html"))
		t.Execute(w, "用户名和密码不能为空")
		return
	}
	user, _ := dao.CheckUsername(username)
	if user.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/user/register.html"))
		t.Execute(w, "该用户名已被使用")
	} else {
		err := dao.SaveUser(username, password, email)
		if err != nil {
			t := template.Must(template.ParseFiles("views/pages/user/register.html"))
			t.Execute(w, "注册失败")
		} else {
			user, _ := dao.CheckUsernameAndPassword(username, password)
			if user.ID == 0 {
				t := template.Must(template.ParseFiles("views/pages/user/register.html"))
				t.Execute(w, "注册失败")
				return
			}
			session_id := utils.CreateUUID()
			sess := &model.Session{
				SessionID: session_id,
				Username: username,
				UserID: user.ID,
			}
			err := dao.AddSession(sess)
			if err != nil {
				log.Println(err)
				t := template.Must(template.ParseFiles("views/pages/user/register.html"))
				t.Execute(w, "注册成功，但保存cookie失败，请重新登录")
				return
			}
			cookie := http.Cookie{
				Name: "session_id",
				Value: session_id,
			}
			http.SetCookie(w, &cookie)
			t := template.Must(template.ParseFiles("views/pages/user/register_success.html"))
			t.Execute(w, username)
		}
	}
}

func CheckUsername(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user, _ := dao.CheckUsername(username)
	if username == "" {
		user.ID = 0x3f3f3f3f
	}
	w.Write([]byte(strconv.Itoa(user.ID)))
}