package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacgoncod/go-project-example/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error listing openings on database")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
