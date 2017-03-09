package main

import (
	"errors"
	"net/http"

	"github.com/go-kit/kit/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const Secret string = "qcon2017"
const SuccessMessage string = "You have won!"
const FailMessage string = "You lose!"
const LambdaURL string = "http://router.fission/slackPost"

type Service interface {
	GetDeal(id int) (Deal, error)
	GetSecret(code string) (string, error)
}

type Deal struct {
	Id   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

func NewDealService(db *mgo.Session, logger log.Logger) Service {
	return &dealService{
		db:     *db,
		logger: logger,
	}
}

type dealService struct {
	db     mgo.Session
	logger log.Logger
}

func (s *dealService) GetDeal(id int) (Deal, error) {
	c := s.db.DB("test").C("deals")
	r := Deal{}
	err := c.Find(bson.M{"id": id}).One(&r)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (s *dealService) GetSecret(code string) (string, error) {
	if code == Secret {
		_, err := http.Get(LambdaURL)
		if err != nil {
			return "", errors.New("Error calling lambda function")
		}
		return SuccessMessage, nil
	}
	return FailMessage, nil
}
