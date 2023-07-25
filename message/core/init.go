package core

import (
	"database/sql"
	_conversationDeliveryHttp "message/core/conversation/delivery/http"
	_conversationR "message/core/conversation/repository/pg"
	_conversationU "message/core/conversation/usecase"

	_utilU "message/core/util/usecase"

	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(db *sql.DB){
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,echo.HeaderAccessControlAllowCredentials},
	  }))
	// e.Use(middleware.Logger())
	timeoutContext := time.Duration(5) * time.Second

	utilU := _utilU.NewUseCase()

	//Conversation
	conversationR := _conversationR.NewRepository(db)
	conversationU := _conversationU.NewUseCase(timeoutContext,conversationR,utilU)
	_conversationDeliveryHttp.NewHandler(e,conversationU)

	e.Start("0.0.0.0:9091")
}