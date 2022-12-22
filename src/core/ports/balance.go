package ports

import (
	"math/big"
)

type BalanceRepository interface {
	GetBalance(address string) (*big.Int, error)
}
