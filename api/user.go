package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/husky_dusky/simplebank/db/sqlc"
	"github.com/husky_dusky/simplebank/util"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username         string    `json:"username"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	CreatedAt        time.Time `json:"created_at"`
	PasswordChangeAt time.Time `json:"password_change_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := createUserResponse{
		Username:         user.Username,
		FullName:         user.FullName,
		Email:            user.Email,
		CreatedAt:        user.CreatedAt,
		PasswordChangeAt: user.PasswordChangeAt,
	}

	ctx.JSON(http.StatusOK, response)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
}
