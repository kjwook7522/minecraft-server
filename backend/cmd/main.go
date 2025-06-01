package main

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"minecraft-server/config"
	"minecraft-server/handler"
	"minecraft-server/service"
)

func main() {
	// 설정 초기화
	if err := config.Init(); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 서비스 초기화
	ctx := context.Background()
	instanceService, err := service.NewInstanceService(ctx)
	if err != nil {
		log.Fatalf("failed to create instance service: %v", err)
	}

	// 핸들러 초기화
	instanceHandler := handler.NewInstanceHandler(instanceService)

	// Echo 인스턴스 생성
	e := echo.New()

	// 미들웨어 설정
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// 라우터 설정
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	// 서버 관리 API
	api := e.Group("/api")
	{
		server := api.Group("/server")
		server.POST("/start", instanceHandler.StartServer)
		server.POST("/stop", instanceHandler.StopServer)
		server.GET("/status", instanceHandler.GetServerStatus)
	}

	// 서버 시작
	e.Logger.Fatal(e.Start(":8080"))
}
