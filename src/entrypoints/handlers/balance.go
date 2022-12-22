package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/luisfc68/luis-coin/src/configuration"
	"github.com/luisfc68/luis-coin/src/core/domain"
	"math/big"
	"net/http"
)

type BalanceRS struct {
	Balance *big.Int `json:"balance"`
}

func GetBalanceHandler(server configuration.Server) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		accountAddress, ok := mux.Vars(request)["account"]
		if !ok {
			jsonError(writer, "invalid path variable name", http.StatusInternalServerError)
			return
		}

		balance, err := server.BalanceService().GetBalance(accountAddress)
		if err != nil && errors.Is(err, domain.InvalidAddressError) {
			jsonError(writer, err.Error(), http.StatusBadRequest)
		} else if err != nil {
			jsonError(writer, err.Error(), http.StatusInternalServerError)
		} else {
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(&BalanceRS{
				Balance: balance,
			})
		}
	}
}
