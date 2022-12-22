package utils

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type AccountInfo struct {
	PublicKey  *crypto.PublicKey
	PrivateKey *ecdsa.PrivateKey
	Address    *common.Address
}

func GetAccountInfo(privateKeyString string) (*AccountInfo, error) {
	privateKey, err := eth.HexToECDSA(privateKeyString)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid key")
	}

	address := eth.PubkeyToAddress(*publicKeyECDSA)

	return &AccountInfo{
		Address:    &address,
		PublicKey:  &publicKey,
		PrivateKey: privateKey,
	}, nil
}

func AuthenticateAccount(client *ethclient.Client, privateKeyString string) (*bind.TransactOpts, error) {
	accountInfo, err := GetAccountInfo(privateKeyString)

	if err != nil {
		return nil, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), *accountInfo.Address)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(accountInfo.PrivateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = 0
	auth.GasPrice = nil

	return auth, nil
}
