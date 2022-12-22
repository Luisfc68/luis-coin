package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/luisfc68/luis-coin/src/configuration"
	"github.com/luisfc68/luis-coin/src/contracts"
	"github.com/luisfc68/luis-coin/src/core/services"
	"github.com/luisfc68/luis-coin/src/dataproviders/blockchain"
	"github.com/luisfc68/luis-coin/src/entrypoints/handlers"
	"github.com/luisfc68/luis-coin/src/entrypoints/middlewares"
	"net/http"
	"os"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("RPC_SERVER_URL"))
	if err != nil {
		panic(err)
	}
	conn := deployAndConnectToContract(client)
	server := createServer(conn, client)
	server.Start(routes)
}

func deployAndConnectToContract(client *ethclient.Client) *contracts.Contracts {
	deployedContract := configuration.DeployContract(client, os.Getenv("ADMIN_ACCOUNT_ADDRESS"))
	conn, err := contracts.NewContracts(common.HexToAddress(deployedContract.Address.Hex()), client)
	if err != nil {
		panic(err)
	}
	return conn
}

func createServer(conn *contracts.Contracts, client *ethclient.Client) *configuration.Broker {
	serverConfig := &configuration.ServerConfig{
		Port: os.Getenv("PORT"),
	}

	balanceRepository := blockchain.NewBalanceRepositoryImpl(conn, client)
	balanceService := services.NewBalanceService(balanceRepository)

	transferRepository := blockchain.NewTransferRepositoryImpl(conn, client)
	transferService := services.NewTransferService(transferRepository, balanceRepository)

	server, err := configuration.NewServer(serverConfig, balanceService, transferService)
	if err != nil {
		panic(err)
	}
	return server
}

func routes(server configuration.Server, router *mux.Router) {
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middlewares.LogMiddleware(server))
	api.Use(middlewares.ContentTypeMiddleware(server))

	api.HandleFunc("/balances/{account}", handlers.GetBalanceHandler(server)).Methods(http.MethodGet)
	api.HandleFunc("/transfers", handlers.PostTransferHandler(server)).Methods(http.MethodPost)
}
