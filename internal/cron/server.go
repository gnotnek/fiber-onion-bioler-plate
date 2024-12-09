package cron

import (
	"context"
	"fiber-onion-boiler-plate/internal/config"
	"fiber-onion-boiler-plate/internal/health"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewServer() *Server {
	cfg := config.Load()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal().Msgf("failed to opening db conn: %s", err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().Msgf("failed to get db object: %s", err.Error())
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	healthRepo := health.NewRepository(db)
	healthSvc := health.NewService(healthRepo)

	return &Server{
		svc: healthSvc,
	}
}

type Server struct {
	svc *health.Service
}

func (s *Server) Run() {
	log.Info().Msg("Cron server is running")

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(1).Minute().Do(s.svc.Check(context.Background()))

	scheduler.StartBlocking()
}
