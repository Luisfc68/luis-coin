package mocks

import (
	"github.com/luisfc68/luis-coin/src/core/domain"
	"math/big"
)

type BalanceRepositoryMock struct{}
type BalanceServiceMock struct{}

func (repo *BalanceRepositoryMock) GetBalance(address string) (*big.Int, error) {
	return getBalance(address)
}

func (repo *BalanceServiceMock) GetBalance(address string) (*big.Int, error) {
	return getBalance(address)
}

func getBalance(address string) (*big.Int, error) {
	if address == "NOT_VALID" {
		return nil, domain.InvalidAddressError
	}
	return big.NewInt(50), nil
}
