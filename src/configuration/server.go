package configuration

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/luisfc68/luis-coin/src/core/services"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type ServerConfig struct {
	Port string
}

type Server interface {
	ServerConfig() *ServerConfig
	BalanceService() *services.BalanceService
	TransferService() *services.TransferService
}

type Broker struct {
	config          *ServerConfig
	router          *mux.Router
	balanceService  *services.BalanceService
	transferService *services.TransferService
}

func NewServer(config *ServerConfig, balanceService *services.BalanceService, transferService *services.TransferService) (*Broker, error) {
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

func (broker *Broker) BalanceService() *services.BalanceService {
	return broker.balanceService
}

func (broker *Broker) TransferService() *services.TransferService {
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
