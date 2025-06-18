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
	RoundSecond    = 60
	DefaultBalance = 2000
	UserMember     = "game"
	BetThisRound   = "bet_this_round"
)

type GameService struct {
	repo  *repository.RedisRepo
	Round int
}

func NewGameService(r *repository.RedisRepo) *GameService {
	r.Del(UserMember, BetThisRound) // 清空
	return &GameService{repo: r}
}

func (s *GameService) GameServer() {
	ticker := time.NewTicker(RoundSecond * time.Second)
	for {
		s.Round++
		start := time.Now()
		log.Println(start.Format("2006-01-02 15:04:05"), "round", s.Round, "start")
		<-ticker.C

		prize := s.GetPrize()
		bets := s.GetBets()
		if len(bets) == 0 {
			log.Println("Round", s.Round, "沒有任何玩家下注")
			continue
		}

		winNum := rand.Intn(prize + 1)
		var winner string
		for _, b := range bets {
			winNum -= b.Amount
			if winNum <= 0 {
				winner = b.Id
				break
			}
		}
		log.Println("獎金池:", prize, "得主:", winner)
		s.repo.ZIncrBy(UserMember, float64(prize), winner)
		s.repo.Del(BetThisRound)
	}
}

func (s *GameService) GetPrize() int {
	bets, _ := s.repo.ZRangeWithScores(BetThisRound)
	sum := 0
	for _, z := range bets {
		sum += int(z.Score)
	}
	return sum
}

func (s *GameService) GetBets() []models.UserBet {
	bets, _ := s.repo.ZRangeWithScores(BetThisRound)
	var ret []models.UserBet
	for _, z := range bets {
		ret = append(ret, models.UserBet{Id: fmt.Sprint(z.Member), Round: s.Round, Amount: int(z.Score)})
	}
	return ret
}

func (s *GameService) GetBalance(uid string) (models.User, error) {
	score, err := s.repo.ZScore(UserMember, uid)
	if err != nil {
		if err == redis.Nil {
			_ = s.repo.ZAdd(UserMember, DefaultBalance, uid)
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
	s.repo.ZIncrBy(UserMember, -float64(amt), uid)
	s.repo.ZIncrBy(BetThisRound, float64(amt), uid)
	return u, nil
}
