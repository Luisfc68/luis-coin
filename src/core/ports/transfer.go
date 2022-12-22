package ports

import "github.com/luisfc68/luis-coin/src/core/domain"

type TransferRepository interface {
	InsertTransfer(transfer *domain.Transfer) error
}
