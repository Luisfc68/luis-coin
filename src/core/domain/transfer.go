package domain

import (
	"errors"
	"math/big"
)

var InsufficientFundsError = errors.New("insufficient funds")

type Transfer struct {
	Amount             *big.Int
	Key                string
	DestinationAccount string
	OriginAccount      string
}
