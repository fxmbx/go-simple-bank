package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/fxmbx/go-simple-bank/db/mock"
	db "github.com/fxmbx/go-simple-bank/db/sqlc"
	"github.com/fxmbx/go-simple-bank/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestTransferAPI(t *testing.T) {
	account1 := randomAccount()
	account2 := randomAccount()
	account1.Currency = utils.CAD
	account2.Currency = utils.CAD
	amount := int64(10)

	currency1 := account1.Currency
	// currency2 := account2.Currency

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"from_account_id": account1.ID,
				"to_account_id":   account2.ID,
				"amount":          amount,
				"currency":        currency1,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccountByID(gomock.Any(), gomock.Eq(account1.ID)).Times(1).Return(account1, nil)
				store.EXPECT().GetAccountByID(gomock.Any(), gomock.Eq(account2.ID)).Times(1).Return(account2, nil)

				arg := db.TransferTxParams{
					FromAccountID: account1.ID,
					ToAccountID:   account2.ID,
					Amount:        amount,
					// Currency: account1.Currency,
				}
				store.EXPECT().TransferTx(gomock.Any(), gomock.Eq(arg)).Times(1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "FromAccountNotFound",
			body: gin.H{
				"from_account_id": account1.ID,
				"to_account_id":   account2.ID,
				"amount":          amount,
				"currency":        utils.CAD,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccountByID(gomock.Any(), gomock.Any()).Times(1).Return(db.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "ToAccountNotFound",
			body: gin.H{
				"from_account_id": account1.ID,
				"to_account_id":   account2.ID,
				"amount":          amount,
				"currency":        utils.CAD,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAccountByID(gomock.Any(), gomock.Eq(account1.ID)).Times(1).Return(account1, nil)
				store.EXPECT().GetAccountByID(gomock.Any(), gomock.Any()).Times(1).Return(db.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
			// store.EXPECT().GetAccountByID(gomock.Any(), gomock.Any()).Times(1).Return(db.Account{}, sql.ErrNoRows)
			// store.EXPECT().GetAccountByID(gomock.Any(), gomock.Eq(account2.ID)).Times(0)

			// arg := db.TransferTxParams{
			// 	FromAccountID: account1.ID,
			// 	ToAccountID:   account2.ID,
			// 	Amount:        amount,
			// }
			// store.EXPECT().TransferTx(gomock.Any(), gomock.Eq(arg)).Times(1)

			server := NewServer(store)
			recorder := httptest.NewRecorder()
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)
			require.NotEmpty(t, data)

			url := "/api/transfer"
			// body := gin.H{
			// 	"from_account_id": account1.ID,
			// 	"to_account_id":   account2.ID,
			// 	"account":         amount,
			// 	"currency":        currency1,
			// }
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}
