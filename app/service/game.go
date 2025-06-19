package service

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Hu-jay/Lottery_Game/app/models"
	"github.com/Hu-jay/Lottery_Game/app/repository"
	"github.com/redis/go-redis/v9"
)

const (
	RoundSecond    = 30
	DefaultBalance = 2000
	UserMember     = "game"
	BetThisRound   = "bet_this_round"
)

type GameService struct {
	redis *repository.RedisRepo
	mysql *repository.MySQLRepo
	Round int
}

func NewGameService(r *repository.RedisRepo, m *repository.MySQLRepo) *GameService {
	r.Del(UserMember, BetThisRound) // 清空
	return &GameService{redis: r, mysql: m}
}

func (s *GameService) GameServer() {
	ticker := time.NewTicker(RoundSecond * time.Second)
	for {
		s.Round++
		log.Println("round start", s.Round)
		<-ticker.C
		bets := s.GetBets()
		prize := s.GetPrize()
		if len(bets) == 0 {
			log.Println("沒有玩家下注")
			continue
		}
		win := rand.Intn(prize + 1)
		var w string
		for _, b := range bets {
			win -= b.Amount
			if win <= 0 {
				w = b.Id
				break
			}
		}
		log.Println("獲得金額", prize, "贏家是", w)
		s.redis.ZIncrBy(UserMember, float64(prize), w)
		s.mysql.SaveBets(bets)
		s.redis.Del(BetThisRound)
	}
}

func (s *GameService) GetPrize() int {
	bets, _ := s.redis.ZRangeWithScores(BetThisRound)
	sum := 0
	for _, z := range bets {
		sum += int(z.Score)
	}
	return sum
}

func (s *GameService) GetBets() []models.UserBet {
	bets, _ := s.redis.ZRangeWithScores(BetThisRound)
	var ret []models.UserBet
	for _, z := range bets {
		ret = append(ret, models.UserBet{Id: fmt.Sprint(z.Member), Round: s.Round, Amount: int(z.Score)})
	}
	return ret
}

func (s *GameService) GetBalance(uid string) (models.User, error) {
	score, err := s.redis.ZScore(UserMember, uid)
	if err != nil {
		if err == redis.Nil {
			_ = s.redis.ZAdd(UserMember, DefaultBalance, uid)
			return models.User{Id: uid, Balance: DefaultBalance}, nil
		}
		return models.User{}, err
	}
	return models.User{Id: uid, Balance: int(score)}, nil
}

func (s *GameService) Bet(uid string, amt int) (models.User, error) {
	u, err := s.GetBalance(uid)
	if err != nil {
		return models.User{}, err
	}
	if amt <= 0 {
		return models.User{}, errors.New("下注金額需為正整數")
	}
	if amt > u.Balance {
		return models.User{}, errors.New("餘額不足")
	}
	u.Balance -= amt
	s.redis.ZIncrBy(UserMember, -float64(amt), uid)
	s.redis.ZIncrBy(BetThisRound, float64(amt), uid)
	return u, nil
}

func (s *GameService) GetHistory(user string) ([]models.BetRecord, error) {
	return s.mysql.GetHistory(user), nil
}
