package blockchain

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/luisfc68/luis-coin/src/contracts"
)

type SmartContractAdapter struct {
	contract *contracts.Contracts
	client   *ethclient.Client
}
