package handler

import (
	rest_api "api-learn/rest-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h Handler) createItem() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, err := getUserId(context)
		if err != nil {
			return
		}

		listId, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			newErrorResponse(context, http.StatusBadRequest, "invalid list id param")
			return
		}

		var input rest_api.TodoItem
		if err := context.BindJSON(&input); err != nil {
			newErrorResponse(context, http.StatusBadRequest, err.Error())
			return
		}

		id, err := h.service.TodoItem.Create(userId, listId, input)
		if err != nil {
			newErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}
}

func (h Handler) getAllItems() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, err := getUserId(context)
		if err != nil {
			return
		}

		listId, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			newErrorResponse(context, http.StatusBadRequest, "invalid list id param")
			return
		}

		items, err := h.service.TodoItem.GetAll(userId, listId)
		if err != nil {
			newErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, items)
	}
}

func (h Handler) getItemById() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, err := getUserId(context)
		if err != nil {
			return
		}

		itemId, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			newErrorResponse(context, http.StatusBadRequest, "invalid list id param")
			return
		}

		item, err := h.service.TodoItem.GetById(userId, itemId)
		if err != nil {
			newErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, item)
	}
}

func (h Handler) updateItem() gin.HandlerFunc {
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

		var input rest_api.UpdateItemInput
		if err := context.BindJSON(&input); err != nil {
			newErrorResponse(context, http.StatusBadRequest, err.Error())
			return
		}

		if err := h.service.TodoItem.Update(userId, id, input); err != nil {
			newErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, statusResponse{Status: "ok"})
	}
}

func (h Handler) deleteItem() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId, err := getUserId(context)
		if err != nil {
			return
		}

		itemId, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			newErrorResponse(context, http.StatusBadRequest, "invalid list id param")
			return
		}

		err = h.service.TodoItem.Delete(userId, itemId)
		if err != nil {
			newErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, statusResponse{
			Status: "ok",
		})
	}
}
