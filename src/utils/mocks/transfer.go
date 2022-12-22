package mocks

import (
	"errors"
	"github.com/luisfc68/luis-coin/src/core/domain"
)

type TransferRepositoryMock struct{}
type TransferServiceMock struct{}

func (repo *TransferRepositoryMock) InsertTransfer(transfer *domain.Transfer) error {
	if transfer.DestinationAccount == "NOT_VALID" || transfer.OriginAccount == "NOT_VALID" {
		return domain.InvalidAddressError
	} else if transfer.Key == "NOT_VALID" {
		return errors.New("invalid key")
	}
	return nil
}
func (repo *TransferServiceMock) MakeTransfer(transfer *domain.Transfer) error {
	if transfer.DestinationAccount == "NOT_VALID" || transfer.OriginAccount == "NOT_VALID" {
		return domain.InvalidAddressError
	} else if transfer.Key == "NOT_VALID_KEY_TRANSFER" {
		return errors.New("invalid key")
	} else if transfer.DestinationAccount == "INSUFFICIENT_AMOUNT_TRANSFER" {
		return domain.InsufficientFundsError
	}
	return nil
}
