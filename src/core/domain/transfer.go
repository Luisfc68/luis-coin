package domain

import (
	"errors"
	"math/big"
)

var InsufficientFoundsError = errors.New("insufficient founds")

type Transfer struct {
	Amount             *big.Int
	Key                string
	DestinationAccount string
	OriginAccount      string
}
