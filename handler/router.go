package handler

import (
	"github.com/labstack/echo"
)

// handler/router.goでは、URLのルーティングを定義
func InitRouting(e *echo.Echo, dictHandler DictHandler) {

	e.GET("/", dictHandler.View())

	e.GET("/searchWord", dictHandler.SearchWord())

	e.GET("/searchTag", dictHandler.SearchTag())

	e.POST("/dictCreate", dictHandler.Add())

	e.POST("/dictEdit", dictHandler.Edit())
}
