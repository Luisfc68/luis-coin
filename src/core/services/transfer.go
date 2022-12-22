package services

import (
	"github.com/luisfc68/luis-coin/src/core/domain"
	"github.com/luisfc68/luis-coin/src/core/ports"
)

type TransferService struct {
	transferRepository ports.TransferRepository
	balanceRepository  ports.BalanceRepository
}

func NewTransferService(transferRepository ports.TransferRepository, balanceRepository ports.BalanceRepository) *TransferService {
	return &TransferService{
		transferRepository: transferRepository,
		balanceRepository:  balanceRepository,
	}
}

func (service *TransferService) MakeTransfer(transfer *domain.Transfer) error {
	balance, err := service.balanceRepository.GetBalance(transfer.OriginAccount)
	if err != nil {
		return err
	} else if balance.Cmp(transfer.Amount) == -1 {
		return domain.InsufficientFoundsError
	}

	return service.transferRepository.InsertTransfer(transfer)
}
