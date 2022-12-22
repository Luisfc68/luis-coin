package handlers

import (
	"github.com/gorilla/mux"
	"github.com/luisfc68/luis-coin/src/utils/mocks"
	"github.com/luisfc68/luis-coin/src/utils/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

type getBalanceTestCase struct {
	test.HandlerTestCase
	address string
}

func TestGetBalanceHandler(t *testing.T) {
	server := mocks.NewMockServer()
	balanceHandler := GetBalanceHandler(server)

	router := mux.NewRouter()
	router.HandleFunc("/api/balances/{account}", balanceHandler)

	testCases := []getBalanceTestCase{
		{
			HandlerTestCase: test.HandlerTestCase{
				Name:         "Returns 400 on invalid address",
				ExpectedCode: http.StatusBadRequest,
				ExpectedBody: `{"error":"invalid address"}`,
			},
			address: "NOT_VALID",
		},
		{
			HandlerTestCase: test.HandlerTestCase{
				Name:         "Returns 200 on valid address",
				ExpectedCode: http.StatusOK,
				ExpectedBody: `{"balance":50}`,
			},
			address: "VALID",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			doGetBalanceTestCase(t, router, &testCase)
		})
	}
}

func doGetBalanceTestCase(t *testing.T, router *mux.Router, testCase *getBalanceTestCase) {
	req, err := http.NewRequest(http.MethodGet, "/api/balances/"+testCase.address, nil)
	if err != nil {
		t.Fatal(err)
	}

	writer := httptest.NewRecorder()
	currenTest := &test.HandlerTest{
		T:        t,
		TestCase: testCase.HandlerTestCase,
		Req:      req,
		Writer:   writer,
		Router:   router,
	}
	test.DoHandlerTest(currenTest)
}
