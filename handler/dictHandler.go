package handler

import (
	"github.com/FirstSS-Sub/Dictionary-search/domain/model"
	"github.com/FirstSS-Sub/Dictionary-search/usecase"
	"net/http"
	"github.com/labstack/echo"
)

type DictHandler struct {
	dictUsecase usecase.DictUsecase
}

// handler/dictHandler.goは、routerから呼び出され、
// リクエストで渡されたデータをusecase層への受け渡す
// 戻り値として受け取ったデータを、JSON形式で返す
func NewDictHandler(dictUsecase usecase.DictUsecase) DictHandler {
	dictHandler := DictHandler{dictUsecase: dictUsecase}
	return dictHandler
}

// 表示
func (handler *DictHandler) View() echo.HandlerFunc {
	return func(c echo.Context) error {
		models, err := handler.dictUsecase.View()
		if err != nil {
			return c.JSON(http.StatusBadRequest, models)
		}
		return c.JSON(http.StatusOK, models)
	}

}
func (handler *DictHandler) SearchWord() echo.HandlerFunc {
	return func(c echo.Context) error {
		word := c.QueryParam("word")
		models, err := handler.dictUsecase.SearchWord(word)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models)
		}
		return c.JSON(http.StatusOK, models)
	}
}
func (handler *DictHandler) SearchTag() echo.HandlerFunc {
	return func(c echo.Context) error {
		tag := c.QueryParam("tag")
		models, err := handler.dictUsecase.SearchTag(tag)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models)
		}
		return c.JSON(http.StatusOK, models)
	}
}
func (handler *DictHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var dict model.Dict
		c.Bind(&dict)
		err := handler.dictUsecase.Add(&dict)
		return c.JSON(http.StatusOK, err)
	}
}
func (handler *DictHandler) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		var dict model.Dict
		c.Bind(&dict)
		err := handler.dictUsecase.Edit(&dict)
		return c.JSON(http.StatusOK, err)
	}
}
