package model

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sales struct {
	Id       int64              `bson:"_id,omitempty"`
	Date     primitive.DateTime `bson:"date,omitempty"`
	Item     string             `bson:"item,omitempty"`
	Price    float64            `bson:"price,omitempty"`
	Quantity int64              `bson:"quantity,omitempty"`
}

func (s *Sales) String() string {
	return fmt.Sprintf("Id : %v, Date : %v, Item : %v, Price :%v, Quantity :%v", s.Id, s.Date, s.Item, s.Price, s.Quantity)
}

// ? rank solo
type RankSoloDocument struct {
	Tier    string `bson:"tier,omitempty"`
	TierNum int64  `bson:"tierNum,omitempty"`
	Score   int64  `bson:"score,omitempty"`
}

type Users struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Summoner string             `bson:"summoner,omitempty"`
	RankSolo RankSoloDocument   `bson:"rank_solo"`
}

func (s *Users) String() string {
	return fmt.Sprintf("Id : %v, summoner : %v, Tier : %v, TierNum :%v, Score :%v",
		s.Id, s.Summoner, s.RankSolo.Tier, s.RankSolo.TierNum, s.RankSolo.Score)
}
