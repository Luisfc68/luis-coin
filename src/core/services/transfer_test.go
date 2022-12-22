package services

import (
	"errors"
	"github.com/luisfc68/luis-coin/src/core/domain"
	"github.com/luisfc68/luis-coin/src/utils/mocks"
	"math/big"
	"testing"
)

type transferTestCase struct {
	name           string
	input          *domain.Transfer
	expectedResult error
}

func doTransferTestCase(service *TransferServiceImpl, testCase *transferTestCase, t *testing.T) {
	err := service.MakeTransfer(testCase.input)
	if !errors.Is(err, testCase.expectedResult) {
		t.Errorf("Expected %v got %v", testCase.expectedResult, err)
	}
}

func TestTransferService_MakeTransfer(t *testing.T) {
	transferService := NewTransferService(&mocks.TransferRepositoryMock{}, &mocks.BalanceRepositoryMock{})
	testCases := []transferTestCase{
		{
			name: "Should fail on invalid destination address",
			input: &domain.Transfer{
				Amount:             big.NewInt(10),
				DestinationAccount: "NOT_VALID",
				OriginAccount:      "123",
				Key:                "456",
			},
			expectedResult: domain.InvalidAddressError,
		},
		{
			name: "Should fail on invalid key",
			input: &domain.Transfer{
				Amount:             big.NewInt(10),
				DestinationAccount: "123",
				OriginAccount:      "NOT_VALID",
				Key:                "NOT_VALID",
			},
			expectedResult: domain.InvalidAddressError,
		},
		{
			name: "Should fail on insufficient funds",
			input: &domain.Transfer{
				Amount:             big.NewInt(60),
				DestinationAccount: "123",
				OriginAccount:      "456",
				Key:                "789",
			},
			expectedResult: domain.InsufficientFundsError,
		},
		{
			name: "Should make transfer ok",
			input: &domain.Transfer{
				Amount:             big.NewInt(30),
				DestinationAccount: "123",
				OriginAccount:      "456",
				Key:                "789",
			},
			expectedResult: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doTransferTestCase(transferService, &testCase, t)
		})
	}
}
