package services

import (
	"github.com/luisfc68/luis-coin/src/core/ports"
	"math/big"
)

type BalanceServiceImpl struct {
	balanceRepository ports.BalanceRepository
}

func NewBalanceService(balanceRepository ports.BalanceRepository) *BalanceServiceImpl {
	return &BalanceServiceImpl{
		balanceRepository: balanceRepository,
	}
}

func (service *BalanceServiceImpl) GetBalance(userAddress string) (*big.Int, error) {
	balance, err := service.balanceRepository.GetBalance(userAddress)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
