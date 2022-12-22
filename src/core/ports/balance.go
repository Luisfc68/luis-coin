package ports

import (
	"math/big"
)

type BalanceRepository interface {
	GetBalance(address string) (*big.Int, error)
}

type BalanceService interface {
	GetBalance(address string) (*big.Int, error)
}
