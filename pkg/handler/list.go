package handler

import (
	rest_api "api-learn/rest-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
@Summary Create a new todo list
@Security ApiKeyAuth
@Tags Lists
@Description Create a new todo list
@ID create-list
@Accept json
@Produce json
@Param input body rest_api.TodoList true "list info"
@Success 200 {integer} integer 1
@Failure 400, 404 {object} errorResponse
@Failure 500 {object} errorResponse
@Failure default {object} errorResponse
@Router /api/lists [post]
*/
func (h Handler) createList() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, err := getUserId(context)
		if err != nil {
			return
		}

		var input rest_api.TodoList
		if err := context.BindJSON(&input); err != nil {
			newErrorResponse(context, http.StatusBadRequest, err.Error())
			return
		}

		id, err := h.service.TodoList.Create(userId, input)
		if err != nil {
			newErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}
}

type getAllListsResponse struct {
	Data []rest_api.TodoList `json:"data"`
}

/*
@Summary Get All Lists
@Security ApiKeyAuth
@Tags Lists
@Description Get all lists
@ID get-all-lists
@Accept json
@Produce json
@Success 200 {object} getAllListsResponse
@Failure 400, 404 {object} errorResponse
@Failure 500 {object} errorResponse
@Failure default {object} errorResponse
@Router /api/lists [get]
*/
func (h Handler) getAllLists() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, err := getUserId(context)
		if err != nil {
			return
		}

		lists, err := h.service.TodoList.GetAll(userId)
		if err != nil {
			newErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, getAllListsResponse{
			Data: lists,
		})
	}
}

/*
@Summary Get List By Id
@Security ApiKeyAuth
@Tags Lists
@Description Get list by id
@ID get-list-by-id
@Accept json
@Produce json
@Success 200 {object} rest_api.TodoList
@Failure 400, 404 {object} errorResponse
@Failure 500 {object} errorResponse
@Failure default {object} errorResponse
@Router /api/lists [get]
*/
func (h Handler) getListById() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, err := getUserId(context)
		if err != nil {
			return
		}

		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			newErrorResponse(context, http.StatusBadRequest, "invalid id param")
			return
		}

		list, err := h.service.TodoList.GetById(userId, id)
		if err != nil {
			newErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, list)
	}
}

func (h Handler) updateList() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, err := getUserId(context)
		if err != nil {
			return
		}

		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			newErrorResponse(context, http.StatusBadRequest, "invalid id param")
			return
		}

		var input rest_api.UpdateListInput
		if err := context.BindJSON(&input); err != nil {
			newErrorResponse(context, http.StatusBadRequest, err.Error())
			return
		}

		if err := h.service.TodoList.Update(userId, id, input); err != nil {
			newErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, statusResponse{Status: "ok"})
	}
}

func (h Handler) deleteList() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, err := getUserId(context)
		if err != nil {
			return
		}

		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			newErrorResponse(context, http.StatusBadRequest, "invalid id param")
			return
		}

		err = h.service.TodoList.Delete(userId, id)
		if err != nil {
			newErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, statusResponse{
			Status: "ok",
		})
	}
}
