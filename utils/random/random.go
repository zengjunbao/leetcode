package random

import (
	rand2 "crypto/rand"
	"errors"
	"math/big"
	"math/rand"
	"time"
)

type Gift struct {
	Id     int
	Weight float64
	Type   string
}

func Random(list []Gift) (Gift, error) {
	if len(list) == 0 {
		return Gift{}, nil
	}
	total := 0
	for _, v := range list {
		total += int(v.Weight * 100)
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(total) // [0,99 )
	current := 0
	for _, v := range list {
		current = current + int(v.Weight*100)
		if index < current {
			return v, nil
		}

	}
	return Gift{}, nil
}

type DrawLotteryItemsDTO struct {
	ID          int     `gorm:"column:id"`
	RewardId    int     `gorm:"column:reward_id"`   // 奖品id
	RewardName  string  `gorm:"column:reward_name"` // 奖品名称
	Probability float32 `gorm:"column:probability"` // 中奖概率
}

func Lottery1() (*DrawLotteryItemsDTO, error) {
	list := []*DrawLotteryItemsDTO{{
		ID:          1,
		RewardId:    1,
		RewardName:  "谢谢惠顾",
		Probability: 0.7,
	}, {
		ID:          2,
		RewardId:    2,
		RewardName:  "一等奖",
		Probability: 0.1,
	}, {
		ID:          3,
		RewardId:    3,
		RewardName:  "三等奖",
		Probability: 0.2,
	}}

	total := 0
	for _, v := range list {
		total += int(v.Probability * 10000)
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(total) // [0,99 )
	current := 0
	for _, v := range list {
		current = current + int(v.Probability*10000)
		if index < current {
			return v, nil
		}
	}
	return nil, errors.New("抽奖失败")
}

func Lottery2() (*DrawLotteryItemsDTO, error) {
	list := []*DrawLotteryItemsDTO{{
		ID:          1,
		RewardId:    1,
		RewardName:  "谢谢惠顾",
		Probability: 0.7,
	}, {
		ID:          2,
		RewardId:    2,
		RewardName:  "一等奖",
		Probability: 0.1,
	}, {
		ID:          3,
		RewardId:    3,
		RewardName:  "三等奖",
		Probability: 0.2,
	}}

	total := 0
	for _, v := range list {
		total += int(v.Probability * 10000)
	}

	b := new(big.Int).SetInt64(int64(total))
	value, err := rand2.Int(rand2.Reader, b)
	if err != nil {
		return nil, err
	}

	current := 0
	for _, v := range list {
		current = current + int(v.Probability*10000)
		if value.Int64() < int64(current) {
			return v, nil
		}
	}
	return nil, errors.New("抽奖失败")
}
