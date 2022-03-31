package model

import "time"

type Session struct {
	SessionID string
	Username string
	UserID int
	CreatedAt time.Time
}
