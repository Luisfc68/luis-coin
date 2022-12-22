package services

import (
	"errors"
	"github.com/luisfc68/luis-coin/src/core/domain"
	"github.com/luisfc68/luis-coin/src/utils/mocks"
	"math/big"
	"testing"
)

type balanceTestCase struct {
	name            string
	input           string
	expectedError   error
	expectedBalance *big.Int
}

func areEqualBalances(first, second *big.Int) bool {
	return first == nil && second == nil || first.Cmp(second) == 0
}

func doBalanceTestCase(service *BalanceServiceImpl, testCase *balanceTestCase, t *testing.T) {
	balance, err := service.GetBalance(testCase.input)
	if !errors.Is(err, testCase.expectedError) {
		t.Errorf("Expected error %v got %v", testCase.expectedError, err)
	} else if !areEqualBalances(balance, testCase.expectedBalance) {
		t.Errorf("Expected balance %v got %v", testCase.expectedBalance, err)
	}
}

func TestBalanceService_GetBalance(t *testing.T) {
	balanceService := NewBalanceService(&mocks.BalanceRepositoryMock{})
	testCases := []balanceTestCase{
		{
			name:            "Should fail on invalid destination address",
			input:           "NOT_VALID",
			expectedError:   domain.InvalidAddressError,
			expectedBalance: nil,
		},
		{
			name:            "Should fail on invalid key",
			input:           "123",
			expectedError:   nil,
			expectedBalance: big.NewInt(50),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doBalanceTestCase(balanceService, &testCase, t)
		})
	}
}
