package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/fxmbx/go-simple-bank/db/mock"
	db "github.com/fxmbx/go-simple-bank/db/sqlc"
	"github.com/fxmbx/go-simple-bank/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccounAPI(t *testing.T) {
	account := randomAccount()

	testCases := []struct {
		name          string
		accountId     int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			accountId: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				//build studs for this mock store
				store.EXPECT().
					GetAccountByID(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name:      "NotFound",
			accountId: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				//build studs for this mock store
				store.EXPECT().
					GetAccountByID(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrNoRows)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
				// requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name:      "InternalServerError",
			accountId: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				//build studs for this mock store
				store.EXPECT().
					GetAccountByID(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
				// requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name:      "Invalid Id",
			accountId: 0,
			buildStubs: func(store *mockdb.MockStore) {
				//build studs for this mock store
				store.EXPECT().
					GetAccountByID(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				// requireBodyMatchAccount(t, recorder.Body, account)
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

			// //build studs for this mock store
			// store.EXPECT().
			// 	GetAccountByID(gomock.Any(), gomock.Eq(account.ID)).
			// 	Times(1).
			// 	Return(account, nil)

			//start server
			server := NewServer(store)
			receorder := httptest.NewRecorder()

			url := fmt.Sprintf("/api/accounts/%d", tc.accountId)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			server.router.ServeHTTP(receorder, request)
			tc.checkResponse(t, receorder)
			// require.Equal(t, http.StatusOK, receorder.Code)
			// requireBodyMatchAccount(t, receorder.Body, account)
		})

	}

}

func randomAccount() db.Account {
	return db.Account{
		ID:       utils.RandomInt(1, 1000),
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {

	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotaccount db.Account
	err = json.Unmarshal(data, &gotaccount)
	require.NoError(t, err)

	require.Equal(t, gotaccount, account)

}
