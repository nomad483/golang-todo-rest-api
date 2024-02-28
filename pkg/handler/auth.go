package handler

import (
	rest_api "api-learn/rest-api"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
@Summary SignUp
@Tags Authentication
@Description Create a new user
@ID create-account
@Accept json
@Produce json
@Param input body rest_api.User true "account info"
@Success 200 {integer} integer 1
@Failure 400, 404 {object} errorResponse
@Failure 500 {object} errorResponse
@Failure default {object} errorResponse
@Router /auth/sign-up [post]
*/
func (h Handler) signUp() gin.HandlerFunc {
	return func(context *gin.Context) {
		var input rest_api.User

		if err := context.BindJSON(&input); err != nil {
			newErrorResponse(context, http.StatusBadRequest, err.Error())
			return
		}

		id, err := h.service.Authorization.CreateUser(input)
		if err != nil {
			newErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

/*
@Summary SignIn
@Tags Authentication
@Description Login to the service
@ID login
@Accept json
@Produce json
@Param input body signInInput true "credentials"
@Success 200 {string} string "token"
@Failure 400, 404 {object} errorResponse
@Failure 500 {object} errorResponse
@Failure default {object} errorResponse
@Router /auth/sign-in [post]
*/
func (h Handler) signIn() gin.HandlerFunc {
	return func(context *gin.Context) {
		var input signInInput

		if err := context.BindJSON(&input); err != nil {
			newErrorResponse(context, http.StatusBadRequest, err.Error())
			return
		}

		token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
		if err != nil {
			newErrorResponse(context, http.StatusInternalServerError, err.Error())
			return
		}

		context.JSON(http.StatusOK, map[string]interface{}{
			"token": token,
		})
	}
}
