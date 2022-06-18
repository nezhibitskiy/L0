package main

import (
	"L0/connections"
	"L0/internal/handlers"
	"L0/internal/repository"
	usecase2 "L0/internal/usecase"
	"context"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func main() {

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, err := config.Build()
	if err != nil {
		log.Fatal("zap logger build error")
	}
	logger := prLogger.Sugar()
	//defer func(prLogger *zap.Logger) {
	//	err = prLogger.Sync()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(prLogger)

	if err := godotenv.Load(".env"); err != nil {
		logger.Fatal("error while loading environment ", err)
	}

	stanComposite, err := connections.ConnectToStan()
	if err != nil {
		logger.Fatal("stan composite failed ", err)
	}

	dbComp, err := connections.NewPostgresDBComposite()
	if err != nil {
		logger.Fatal("postgres composite failed ", err)
	}
	r := repository.NewRepository(dbComp.DB)

	redisComp, err := connections.NewRedisComposite()
	if err != nil {
		logger.Fatal("redis composite failed", err)
	}
	redis := repository.NewRedisRepository(redisComp.Redis)

	u := usecase2.NewUsecase(r, redis)

	echoServer := echo.New()

	_, err = handlers.RegisterHandlers(stanComposite.SC, u, echoServer, logger)
	if err != nil {
		logger.Fatal("register handlers failed")
	}

	err = u.InitCache(context.Background())
	if err != nil {
		logger.Error("error during cache initing. Old data will be get from postgres")
	}
	echoServer.Logger.Fatal(echoServer.Start(":8080"))
}
