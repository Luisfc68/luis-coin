package mocks

import (
	"github.com/luisfc68/luis-coin/src/configuration"
	"github.com/luisfc68/luis-coin/src/core/ports"
)

type MockServer struct {
	balanceServiceMock  *BalanceServiceMock
	transferServiceMock *TransferServiceMock
	config              *configuration.ServerConfig
}

func NewMockServer() *MockServer {
	return &MockServer{
		balanceServiceMock:  &BalanceServiceMock{},
		transferServiceMock: &TransferServiceMock{},
		config: &configuration.ServerConfig{
			Port: ":8080",
		},
	}
}

func (server *MockServer) ServerConfig() *configuration.ServerConfig {
	return server.config
}

func (server *MockServer) BalanceService() ports.BalanceService {
	return server.balanceServiceMock
}

func (server *MockServer) TransferService() ports.TransferService {
	return server.transferServiceMock
}
