package configuration

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/luisfc68/luis-coin/src/contracts"
	"github.com/luisfc68/luis-coin/src/utils"
)

type DeployedContract struct {
	Address     common.Address
	Transaction *types.Transaction
	Contract    *contracts.Contracts
}

func DeployContract(client *ethclient.Client, accountAddress string) *DeployedContract {
	auth, err := utils.AuthenticateAccount(client, accountAddress)
	if err != nil {
		panic(err)
	}
	address, transaction, contract, err := contracts.DeployContracts(auth, client)
	if err != nil {
		panic(err)
	}
	return &DeployedContract{
		Address:     address,
		Transaction: transaction,
		Contract:    contract,
	}
}
