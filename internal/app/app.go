package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	h "gtihub.com/blckvia/go-queue/internal/handler"

	"gtihub.com/blckvia/go-queue/internal/broker"
)

type App struct {
	Server *http.Server
	Logger *zap.Logger
}

type QueueConfig struct {
	Name    string `mapstructure:"name"`
	Size    int    `mapstructure:"size"`
	MaxSubs int    `mapstructure:"max_subs"`
}

func NewApp(ctx context.Context, logger *zap.Logger) *App {
	if err := InitConfig(); err != nil {
		logger.Fatal("error initializing configs: %w", zap.Error(err))
	}

	if err := godotenv.Load(); err != nil {
		logger.Fatal("error loading .env file", zap.Error(err))
	}

	var queueConfigs []QueueConfig
	if err := viper.UnmarshalKey("queues", &queueConfigs); err != nil {
		logger.Fatal("error unmarshalling queues", zap.Error(err))
	}

	b := broker.NewBroker()

	for _, cfg := range queueConfigs {
		if err := b.CreateQueue(cfg.Name, cfg.Size, cfg.MaxSubs); err != nil {
			logger.Fatal("error creating queue", zap.String("queue", cfg.Name), zap.Error(err))
		}
		logger.Info(fmt.Sprintf("created queue %s", cfg.Name))
	}

	handlers := h.New(b)

	srv := &http.Server{
		Addr:           ":" + viper.GetString("port"),
		Handler:        handlers.InitRoutes(),
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return &App{
		Server: srv,
		Logger: logger,
	}
}

func (a *App) Run() error {
	return a.Server.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context, logger *zap.Logger) error {
	return a.Server.Shutdown(ctx)
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	return viper.ReadInConfig()
}
