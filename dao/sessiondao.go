package dao

import (
	"bookstore/model"
	"log"
	"net/http"
)

func AddSession(sess *model.Session) error {
	result := db.Create(sess)
	log.Println("Rows affected:", result.RowsAffected)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func DeleteSession(sessionID string) error {
	result := db.Delete(&model.Session{}, "session_id = ?", sessionID)
	log.Println("Rows affected:", result.RowsAffected)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func GetSession(sessionID string) (*model.Session, error) {
	sess := model.Session{}
	err := db.Where("session_id = ?", sessionID).First(&sess).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &sess, nil
}

func IsLogin(r *http.Request) (bool, *model.Session) {
	cookie, _ := r.Cookie("session_id")
	if cookie != nil {
		session_id := cookie.Value
		session, err := GetSession(session_id)
		if err != nil {
			log.Println(err)
		} else if session.Username != "" {
			return true, session
		}
	}
	return false, nil
}