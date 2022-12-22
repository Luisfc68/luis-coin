package blockchain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/luisfc68/luis-coin/src/contracts"
	"github.com/luisfc68/luis-coin/src/core/domain"
	"github.com/luisfc68/luis-coin/src/utils"
)

type TransferRepositoryImpl struct {
	SmartContractAdapter
}

func NewTransferRepositoryImpl(contract *contracts.Contracts, client *ethclient.Client) *TransferRepositoryImpl {
	return &TransferRepositoryImpl{
		SmartContractAdapter{
			client:   client,
			contract: contract,
		},
	}
}

func (adapter *TransferRepositoryImpl) InsertTransfer(transfer *domain.Transfer) error {
	if !common.IsHexAddress(transfer.DestinationAccount) || !common.IsHexAddress(transfer.OriginAccount) {
		return domain.InvalidAddressError
	}

	auth, err := utils.AuthenticateAccount(adapter.client, transfer.Key)
	if err != nil {
		return err
	}

	_, err = adapter.contract.Send(auth, common.HexToAddress(transfer.DestinationAccount), transfer.Amount)

	return err
}
