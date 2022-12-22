package handlers

import (
	"encoding/json"
	"errors"
	"github.com/luisfc68/luis-coin/src/configuration"
	"github.com/luisfc68/luis-coin/src/core/domain"
	"github.com/luisfc68/luis-coin/src/utils"
	"math/big"
	"net/http"
)

type TransferRQ struct {
	Key                string   `json:"key"  validate:"required"`
	DestinationAddress string   `json:"destinationAddress"  validate:"required"`
	Amount             *big.Int `json:"amount" validate:"required,gte=0"`
}

func convertTransfer(dto TransferRQ, originAccount string) *domain.Transfer {
	return &domain.Transfer{
		Key:                dto.Key,
		DestinationAccount: dto.DestinationAddress,
		Amount:             dto.Amount,
		OriginAccount:      originAccount,
	}
}

func PostTransferHandler(server configuration.Server) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		transferRQ := TransferRQ{}
		err := json.NewDecoder(request.Body).Decode(&transferRQ)
		if err != nil {
			jsonError(writer, err.Error(), http.StatusBadRequest)
			return
		}
		err = validate.Struct(transferRQ)
		if err != nil {
			jsonError(writer, err.Error(), http.StatusBadRequest)
			return
		}

		accountInfo, err := utils.GetAccountInfo(transferRQ.Key)
		if err != nil {
			jsonError(writer, err.Error(), http.StatusUnauthorized)
			return
		}

		err = server.TransferService().MakeTransfer(convertTransfer(transferRQ, accountInfo.Address.String()))

		if errors.Is(err, domain.InsufficientFundsError) {
			jsonError(writer, err.Error(), http.StatusConflict)
			return
		} else if errors.Is(err, domain.InvalidAddressError) {
			jsonError(writer, err.Error(), http.StatusBadRequest)
			return
		} else if err != nil {
			jsonError(writer, err.Error(), http.StatusInternalServerError)
			return
		} else {
			writer.WriteHeader(http.StatusNoContent)
			return
		}
	}
}
