package handlers

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/luisfc68/luis-coin/src/utils/mocks"
	"github.com/luisfc68/luis-coin/src/utils/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

type postTransferTestCase struct {
	test.HandlerTestCase
	transfer string
}

func TestPostTransferHandler(t *testing.T) {
	server := mocks.NewMockServer()
	postTransferHandler := PostTransferHandler(server)

	router := mux.NewRouter()
	router.HandleFunc("/api/transfers", postTransferHandler).Methods(http.MethodPost)

	testCases := []postTransferTestCase{
		{
			HandlerTestCase: test.HandlerTestCase{
				Name:         "Returns 400 on incomplete body",
				ExpectedCode: http.StatusBadRequest,
				ExpectedBody: `{"error":"Key: 'TransferRQ.Key' Error:Field validation for 'Key' failed on the 'required' tag\nKey: 'TransferRQ.DestinationAddress' Error:Field validation for 'DestinationAddress' failed on the 'required' tag\nKey: 'TransferRQ.Amount' Error:Field validation for 'Amount' failed on the 'required' tag"}`,
			},
			transfer: `{
				
			}`,
		},
		{
			HandlerTestCase: test.HandlerTestCase{
				Name:         "Returns 400 on invalid body",
				ExpectedCode: http.StatusBadRequest,
				ExpectedBody: `{"error":"invalid character 'a' looking for beginning of object key string"}`,
			},
			transfer: `{ asda }`,
		},
		{
			HandlerTestCase: test.HandlerTestCase{
				Name:         "Returns 401 on invalid key",
				ExpectedCode: http.StatusUnauthorized,
				ExpectedBody: `{"error":"invalid hex character 's' in private key"}`,
			},
			transfer: `{
				"amount": 10, 
				"key": "asdasda",
				"destinationAddress": "0xf6ac737fD028c99b616cf806D887FD85634BD6dd"
			}`,
		},
		{
			HandlerTestCase: test.HandlerTestCase{
				Name:         "Returns 409 on insufficient funds",
				ExpectedCode: http.StatusConflict,
				ExpectedBody: `{"error":"insufficient funds"}`,
			},
			transfer: `{
				"amount": 10, 
				"key": "e6df1e73b2716b40141e0665433ee5fa4834e89dab533e3d505ae6a2aff61d99",
				"destinationAddress": "INSUFFICIENT_AMOUNT_TRANSFER"
			}`,
		},
		{
			HandlerTestCase: test.HandlerTestCase{
				Name:         "Returns 400 on invalid destination address",
				ExpectedCode: http.StatusBadRequest,
				ExpectedBody: `{"error":"invalid address"}`,
			},
			transfer: `{
				"amount": 10, 
				"key": "e6df1e73b2716b40141e0665433ee5fa4834e89dab533e3d505ae6a2aff61d99",
				"destinationAddress": "NOT_VALID"
			}`,
		},
		{
			HandlerTestCase: test.HandlerTestCase{
				Name:         "Returns 204 on successful transfer",
				ExpectedCode: http.StatusNoContent,
			},
			transfer: `{
				"amount": 10, 
				"key": "e6df1e73b2716b40141e0665433ee5fa4834e89dab533e3d505ae6a2aff61d99",
				"destinationAddress": "VALID"
			}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			doPostTransferTestCase(t, router, &testCase)
		})
	}
}

func doPostTransferTestCase(t *testing.T, router *mux.Router, testCase *postTransferTestCase) {
	req, err := http.NewRequest(http.MethodPost, "/api/transfers", bytes.NewReader([]byte(testCase.transfer)))
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
