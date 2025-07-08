# Lottery_Game

輕量級的即時輪盤式投注遊戲 API，使用 Go、Gin 與 Redis 架構。

---

## 🌟 Features

- 玩家註冊與餘額管理（初始餘額 2000）。
- 支援下注功能，並將金額存入當前回合獎金池。
- 每 60 秒由後台開獎，透過加權抽獎選出贏家並派發獎金。
- 分離架構：config / models / controllers / router / redis 服務。
- RESTful API 介面，易於整合與測試。

---

## 📘 專案說明

此專案提供 HTTP API，讓玩家可：

1. `GET /players/:user` – 註冊新玩家或查詢餘額。
2. `GET /players/:user/:amount` – 向當前回合下注。
3. `GET /prize` – 查詢本回合總獎金。
4. `GET /players` – 查詢本回合所有下注紀錄。

後台會每 60 秒觸發一次開獎，包括：
- 累加各玩家下注計算總獎金池；
- 使用加權隨機方式選出贏家；
- 將獎金自動存入贏家帳戶；
- 清空本回合下注記錄，準備下一輪。

---

## 🚀 快速啟動指南

### 環境需求

- Go ≥ 1.20
- Redis Server (預設連線至 `localhost:6379`)
- Gin Framework、go-redis 套件依賴 (已於 `go.mod` 中定義)

### 開發環境啟動步驟

```bash
git clone https://github.com/your_username/Lottery_Game.git
cd Lottery_Game

go mod tidy             # 安裝依賴
go run cmd/server/main.go

伺服器啟動後，瀏覽器或使用 curl/Postman 操作 API，例如：

# 查詢或註冊玩家
curl http://localhost:8080/players/Alice

# 對回合下注
curl http://localhost:8080/players/Alice/500

# 查詢本回合獎金
curl http://localhost:8080/prize

# 取得所有下注資訊
curl http://localhost:8080/players


⸻

📂 專案結構

Lottery_Game/
├── cmd/server/main.go         # 專案進入點，負責初始化 Redis 與 HTTP Server
├── app/
│   ├── config/config.go       # 全域參數設定（如回合時間、初始餘額）
│   ├── models/                # 資料模型（User、UserBet、Ret）
│   ├── redisservice/          # 背景開獎服務，Tick 每 60 秒執行一次
│   ├── controllers/           # API 控制器，處理 HTTP 請求
│   └── router/router.go       # 路由設定與 Redis 注入中介
└── go.mod                     # 專案依賴與版本管理