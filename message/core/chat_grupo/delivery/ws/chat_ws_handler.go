package ws

import (
	// "encoding/json"
	"log"
	// "net/http"
	r "message/domain/repository"

	"github.com/labstack/echo/v4"
	// _jwt "regate-backend/domain/util"

)

type GrupoWsHandler struct {
}

func NewWsHandler(e *echo.Echo,grupoUseCase r.GrupoUseCase){
	go H.Run(grupoUseCase)
	handler :=GrupoWsHandler{
	}
	e.GET("v1/ws/chat-grupo",handler.ChatGrupo)
}

func (ws *GrupoWsHandler) ChatGrupo(c echo.Context) (err error) {
	casoId := c.QueryParam("id")
	log.Println("chat grupo",casoId)
	ServeWs(c.Response(), c.Request(), casoId)
	return nil
}

