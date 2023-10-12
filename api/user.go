package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/kevinbarrero/auth-service/db/sqlc"
	"github.com/kevinbarrero/auth-service/util"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Email    string `json:"email"     binding:"required,email"`
	FullName string `json:"full_name" binding:"required"`
<<<<<<< HEAD
	Password string `json: "password" binding:"required,min=6"`
=======
	Password string `json:"password"  binding:"required,min=6"`
>>>>>>> e708bd5adf104a93d5919efdc61d7835a5f50c8e
}

type userResponse struct {
	Email             string    `json:"email"`
	FullName          string    `json:"full_name"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
}

func newUserResponse(user db.Users) userResponse {
	return userResponse{
		Email:             user.Email,
		FullName:          user.FullName,
		PasswordChangedAt: user.PasswordChangedAt,
	}
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
		Email:          req.Email,
		FullName:       req.FullName,
		HashedPassword: hashedPassword,
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
	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
<<<<<<< HEAD
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
=======
	Email    string `json: "email" binding:"required,email"`
	Password string `json: "password" binding: "required,min=6"`
>>>>>>> e708bd5adf104a93d5919efdc61d7835a5f50c8e
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	accessToken, err := server.tokenMaker.CreateToken(user.Email, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}
