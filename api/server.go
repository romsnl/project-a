package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gitlab "github.com/romsnl/project-a/pkg"
)

type Server struct {
	router *gin.Engine
}

func (s *Server) Serve(addr string) error {
	return s.router.Run(addr)
}

func NewServer() *Server {
	router := gin.Default()
	s := &Server{
		router,
	}

	router.GET("/healthz", s.getHealth)
	router.GET("/api/version", s.getVersion)

	return s
}

func (s *Server) getVersion(c *gin.Context) {
	version := gitlab.GetGitlabVersion()

	c.JSON(http.StatusOK, gin.H{
		"version": version,
	})
}

func (s *Server) getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": "OK",
	})
}
