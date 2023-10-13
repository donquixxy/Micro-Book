package server

import "github.com/gin-gonic/gin"

type Server struct {
	Router *gin.Engine
}

func NewServer(env string) *Server {
	var r *gin.Engine

	if env == "production" {
		r = gin.Default()
		gin.SetMode(gin.ReleaseMode)
	} else {
		r = gin.Default()
	}

	// Init router

	server := &Server{
		Router: r,
	}

	return server
}
