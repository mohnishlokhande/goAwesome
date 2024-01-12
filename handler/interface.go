package handler

import "awesomeProject/model"

type Account interface {
	Create(a []model.Account) error
	Get() ([]model.Account, error)
	Transfer(senderAcct, receiverAcct string, amount float64) error
}
