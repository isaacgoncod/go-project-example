package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacgoncod/go-project-example/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("Error validating request: %v", err.Error())

		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errf("Error creating opening: %v", err.Error())

		sendError(ctx, http.StatusInternalServerError, "Error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)

}
