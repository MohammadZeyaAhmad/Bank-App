package api

import (
	"log"
	"net/http"

	db "github.com/MohammadZeyaAhmad/Bank-App/db/sqlc"
	"github.com/gin-gonic/gin"
)

type VerifyEmailRequest struct {
	ID   int64 `form:"id" binding:"required"`
	Code string `form:"secret_code" binding:"required"`
}

func (server *Server) VerifyEmail(ctx *gin.Context) {
	var req VerifyEmailRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	log.Fatal("passed till params")
	arg := db.UpdateVerifyEmailParams{
		ID:req.ID,
		SecretCode:req.Code}
    log.Fatal("passed till params")
	VerifyEmail, err := server.store.UpdateVerifyEmail(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, VerifyEmail)
}