package services

import (
	"github.com/luisfc68/luis-coin/src/core/ports"
	"math/big"
)

type BalanceService struct {
	balanceRepository ports.BalanceRepository
}

func NewBalanceService(balanceRepository ports.BalanceRepository) *BalanceService {
	return &BalanceService{
		balanceRepository: balanceRepository,
	}
}

func (service *BalanceService) GetBalance(userAddress string) (*big.Int, error) {
	balance, err := service.balanceRepository.GetBalance(userAddress)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
