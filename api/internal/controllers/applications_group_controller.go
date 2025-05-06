package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/database"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/models"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// gin-swagger middleware
// swagger embed files

type ApplicationGroupController struct {
	response.Response
	logger *zap.Logger
}

var controllerInstance *ApplicationGroupController

func NewApplicationGroupController(logger *zap.Logger) *ApplicationGroupController {
	if controllerInstance == nil {
		controllerInstance = &ApplicationGroupController{logger: logger}
	}

	return controllerInstance
}

//	@BasePath	/api/v1

// PingExample godoc
//
//	@Summary	ping example
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/example/helloworld [get]
func (controller *ApplicationGroupController) Index(c *gin.Context) {
	var applicationsGroups []models.ApplicationGroup

	database.DbConnection().Db().Find(&applicationsGroups)

	controller.SendSuccessResponse(gin.H{
		"application_groups": applicationsGroups,
	})
	return
}

func (controller *ApplicationGroupController) Store(c *gin.Context) {
	var group models.ApplicationGroup

	if err := c.ShouldBindJSON(&group); err != nil {
		controller.SendBadRequestWithErr(err)
		return
	}

	group.UserID = uint(1)

	result := database.DbConnection().Db().Create(&group)
	if result.Error != nil {
		controller.logger.Error("Error on storing application group", zap.Error(result.Error))
		controller.SendInternalServerWithErr()
		return
	}

	controller.SendCreatedResponse(gin.H{
		"id": group.ID,
	})

	return
}

func (controller *ApplicationGroupController) Show(c *gin.Context) {
	type idParams struct { // Todo this is like form request in laravel
		ID uint `uri:"id" binding:"required,min=1"`
	}

	var params idParams

	if err := c.ShouldBindUri(&params); err != nil {
		controller.SendBadRequestWithErr(errors.New("Invalid ID " + err.Error()))
		return
	}

	var group models.ApplicationGroup

	result := database.DbConnection().Db().First(&group, params.ID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			controller.SendNotFoundWithErr(
				errors.New("application group not found"),
			)

			return
		}
		controller.logger.Error("Error on showing application group", zap.Error(result.Error))
		controller.SendInternalServerWithErr()
		return
	}

	controller.SendSuccessResponse(gin.H{"group": &group})
	return
}

func (controller *ApplicationGroupController) Update(c *gin.Context) {
	type idParams struct { // Todo this is like form request in laravel
		ID   uint   `json:"id" binding:"required,min=1"`
		Name string `json:"name" binding:"required,min=1,max=255"`
	}

	var params idParams

	if err := c.ShouldBindUri(&params); err != nil {
		controller.SendBadRequestWithErr(
			errors.New("invalid ID " + err.Error()),
		)
		return
	}

	var group models.ApplicationGroup

	result := database.DbConnection().Db().First(&group, params.ID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			controller.SendNotFoundWithErr(
				errors.New("application group not found"),
			)
			return
		}
		controller.logger.Error("Error on updating application group", zap.Error(result.Error))
		controller.SendInternalServerWithErr()

		return
	}

	group.Name = params.Name
	database.DbConnection().Db().Save(&group)

	controller.SendSuccessResponse(gin.H{"group": &group})
	return
}

func (controller *ApplicationGroupController) Delete(c *gin.Context) {

	type idParams struct { // Todo this is like form request in laravel
		ID uint `uri:"id" binding:"required,min=1"`
	}

	var params idParams

	if err := c.ShouldBindUri(&params); err != nil {
		controller.SendBadRequestWithErr(errors.New("Invalid ID " + err.Error()))
		return
	}

	var group models.ApplicationGroup

	result := database.DbConnection().Db().First(&group, params.ID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			controller.SendBadRequestWithErr(errors.New("application group not found"))
			return
		}

		controller.logger.Error("error on deleting group", zap.Uint("group ID", group.ID), zap.Error(result.Error))
		controller.SendInternalServerWithErr()
		return
	}

	database.DbConnection().Db().Delete(&group)
	controller.SendNotContentResponse()
}
