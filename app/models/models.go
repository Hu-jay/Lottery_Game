package models

import "time"

type User struct {
	Id      string `json:"Id"`
	Balance int    `json:"balance"`
}

type UserBet struct {
	Id     string `json:"Id"`
	Round  int    `json:"round"`
	Amount int    `json:"amount"`
}

type BetRecord struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    string    `gorm:"column:user_id;index"`
	Round     int       `gorm:"column:round"`
	Amount    int       `gorm:"column:amount"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Ret struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
