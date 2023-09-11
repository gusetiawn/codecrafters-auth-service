package app

import (
	_ "github.com/gusetiawn/codecrafters-auth-service/internal/config"
	"github.com/gusetiawn/codecrafters-auth-service/internal/database"
	"github.com/gusetiawn/codecrafters-auth-service/internal/handlers"
	"github.com/gusetiawn/codecrafters-auth-service/internal/repository"
	http "github.com/gusetiawn/codecrafters-auth-service/internal/server"
	"github.com/gusetiawn/codecrafters-auth-service/internal/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "boy",
	}
	serveCmd = &cobra.Command{
		Use: "serve",
		Run: serve,
	}
	handlers *http.ServerHandlers
	server   http.Server
)

func init() {
	cobra.OnInitialize(initProject)
	rootCmd.AddCommand(serveCmd)
}

func initProject() {
	server = http.Serve
	dbConnection := database.New()
	repository := repository.New(dbConnection)
	service := service.New(repository)
	handler := handler.New(service)
	handlers = &http.ServerHandlers{
		Handler: handler,
	}
}

func serve(cmd *cobra.Command, args []string) {
	logger := logrus.WithField("func", "serve")
	logger.Info("serve")
	err := server(handlers)
	if err != nil {
		logger.WithError(err).Error("error while running")
		panic(err)
	}
	logger.Info("done")
}

func Execute() {
	rootCmd.Execute()
}
