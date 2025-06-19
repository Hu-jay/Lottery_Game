package repository

import (
	"gorm.io/gorm"

	"github.com/Hu-jay/Lottery_Game/app/models"
)

type MySQLRepo struct{ DB *gorm.DB }

func NewMySQLRepo(db *gorm.DB) *MySQLRepo {
	db.AutoMigrate(&models.BetRecord{})
	return &MySQLRepo{DB: db}
}

func (r *MySQLRepo) SaveBets(bets []models.UserBet) {
	for _, b := range bets {
		r.DB.Create(&models.BetRecord{UserID: b.Id, Round: b.Round, Amount: b.Amount})
	}
}

func (r *MySQLRepo) GetHistory(userID string) []models.BetRecord {
	var recs []models.BetRecord
	r.DB.Where("user_id = ?", userID).Find(&recs)
	return recs
}
