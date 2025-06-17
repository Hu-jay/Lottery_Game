package models

type User struct {
	Id      string `json:"Id"`      // 玩家 ID
	Balance int    `json:"balance"` // 玩家餘額
}

type UserBet struct {
	Id     string `json:"Id"`
	Round  int    `json:"round"`  // 局數
	Amount int    `json:"amount"` // 下注金額
}

type Ret struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
