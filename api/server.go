package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/kevinbarrero/auth-service/db/sqlc"
	"github.com/kevinbarrero/auth-service/token"
	"github.com/kevinbarrero/auth-service/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

