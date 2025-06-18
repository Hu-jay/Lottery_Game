package models

type User struct {
	Id      string `json:"Id"`
	Balance int    `json:"balance"`
}

type UserBet struct {
	Id     string `json:"Id"`
	Round  int    `json:"round"`
	Amount int    `json:"amount"`
}
