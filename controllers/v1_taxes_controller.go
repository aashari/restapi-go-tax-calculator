package controllers

import (
	"../constants"
	"../helpers"
	"../objects"
	"../services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type V1TaxesController struct {
	service     services.V1TaxesService
	errorHelper helpers.ErrorHelper
}

func V1TaxesControllerHandler(router *gin.Engine) {

	handler := &V1TaxesController{
		service:     services.V1TaxesServiceHandler(),
		errorHelper: helpers.ErrorHelperHandler(),
	}

	group := router.Group("v1/taxes")
	{
		group.POST("", handler.CalculateTax)
	}

}

func (handler *V1TaxesController) CalculateTax(context *gin.Context) {

	requestObject := objects.V1TaxesObjectRequest{}
	context.ShouldBind(&requestObject)

	result, err := handler.service.CalculateTax(requestObject)

	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)

}
