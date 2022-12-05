package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
	"net/http"
)

type createEntryRequest struct {
	AccountID int64 `json:"account_id" binding:"required"`
	Amount    int64 `json:"amount" binding:"required,min=1"`
}

func (server *Server) createEntry(ctx *gin.Context) {
	var req createEntryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateEntryParams{
		AccountID: req.AccountID,
		Amount:    req.Amount,
	}

	entry, err := server.store.CreateEntry(ctx, arg)
	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, entry)
}

type getEntryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getEntry(ctx *gin.Context) {
	var req getEntryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	entry, err := server.store.GetEntry(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entry)
}

type listEntryRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listEntries(ctx *gin.Context) {
	var req listEntryRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListEntriesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	entries, err := server.store.ListEntries(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entries)
}
