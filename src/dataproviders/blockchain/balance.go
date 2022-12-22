package blockchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/luisfc68/luis-coin/src/contracts"
	"github.com/luisfc68/luis-coin/src/core/domain"
	"math/big"
)

type BalanceRepositoryImpl struct {
	SmartContractAdapter
}

func NewBalanceRepositoryImpl(contract *contracts.Contracts, client *ethclient.Client) *BalanceRepositoryImpl {
	return &BalanceRepositoryImpl{
		SmartContractAdapter{
			client:   client,
			contract: contract,
		},
	}
}

func (adapter *BalanceRepositoryImpl) GetBalance(address string) (*big.Int, error) {
	if !common.IsHexAddress(address) {
		return nil, domain.InvalidAddressError
	}

	balance, err := adapter.contract.CheckBalance(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}

	return balance, nil
}
