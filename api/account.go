package api

import (
	"database/sql"
	"net/http"

	db "github.com/fxmbx/go-simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// if pqErr, ok := err.(*pq.Error); ok {
		// 	log.Println(pqErr.Code)
		// }
		// ctx.JSON(http.StatusBadRequest, errorResponse("error binging req to json 🎣", err))
		errroHandler(ctx, err)
		return
	}
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		errroHandler(ctx, err)

		// ctx.JSON(http.StatusInternalServerError, errorResponse("error creating account 🎣", err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type GetAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("error binging req to uri 🎣", err))
		return
	}
	account, err := server.store.GetAccountByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse("requested account doesn't exist 🎣", err))
			return

		}
		ctx.JSON(http.StatusInternalServerError, errorResponse("error getting account 🎣", err))
		return
	}
	// account = db.Account{}

	ctx.JSON(http.StatusOK, account)
}

type GetAccountsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getAccounts(ctx *gin.Context) {
	var req GetAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("error binging req to query 🎣", err))
		return
	}
	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	account, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse("requested row doesn't exists🎣", err))
			return

		}
		ctx.JSON(http.StatusInternalServerError, errorResponse("error getting accounts 🎣", err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type GetAccountUpdateRequest struct {
	ID      int64 `json:"id" binding:"required,min=1"`
	Balance int64 `json:"balance" binding:"required,min=0"`
}

func (server *Server) updateAccount(ctx *gin.Context) {
	var req GetAccountUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("error binging req to json 🎣", err))
		return
	}

	arg := db.UpdateAccountParams{
		ID:      req.ID,
		Balance: req.Balance,
	}
	account, err := server.store.UpdateAccount(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse("Requested account for update not found 🎣", err))
			return

		}
		ctx.JSON(http.StatusInternalServerError, errorResponse("error updating account 🎣", err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// func (server *Server) requestBinder(c *gin.Context, req any, reqtype string) error {
// 	var err error
// 	switch reqtype {
// 	case "uri":
// 		err = c.ShouldBindUri(&req)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	case "body":
// 		err = c.ShouldBindJSON(&req)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	case "query":
// 		err = c.ShouldBindQuery(&req)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// 	return nil
// }
