package configuration

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/luisfc68/luis-coin/src/core/ports"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type ServerConfig struct {
	Port string
}

type Server interface {
	ServerConfig() *ServerConfig
	BalanceService() ports.BalanceService
	TransferService() ports.TransferService
}

type Broker struct {
	config          *ServerConfig
	router          *mux.Router
	balanceService  ports.BalanceService
	transferService ports.TransferService
}

func NewServer(config *ServerConfig, balanceService ports.BalanceService, transferService ports.TransferService) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("missing port")
	}
	return &Broker{
		config:          config,
		router:          mux.NewRouter(),
		balanceService:  balanceService,
		transferService: transferService,
	}, nil
}

func (broker *Broker) ServerConfig() *ServerConfig {
	return broker.config
}

func (broker *Broker) BalanceService() ports.BalanceService {
	return broker.balanceService
}

func (broker *Broker) TransferService() ports.TransferService {
	return broker.transferService
}

func (broker *Broker) Start(binder func(server Server, router *mux.Router)) {
	binder(broker, broker.router)
	handler := cors.AllowAll().Handler(broker.router)
	log.Println("Starting server on port", broker.config.Port)
	if err := http.ListenAndServe(broker.config.Port, handler); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
