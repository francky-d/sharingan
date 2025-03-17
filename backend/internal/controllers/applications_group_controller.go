package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/database"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/models"
	"gorm.io/gorm"
)

type ApplicationGroupController struct {
}

var controllerInstance *ApplicationGroupController

func NewApplicationGroupController() *ApplicationGroupController {
	if controllerInstance == nil {
		controllerInstance = &ApplicationGroupController{}
	}

	return controllerInstance
}

func (controller *ApplicationGroupController) Index(c *gin.Context) {
	var applicationsGroups []models.ApplicationGroup

	database.DbConnection().Db().Find(&applicationsGroups)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"application_groups": applicationsGroups,
		},
	})
}

func (controller *ApplicationGroupController) Store(c *gin.Context) {
	var group models.ApplicationGroup

	if err := c.ShouldBindJSON(&group); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	group.UserID = uint(1)

	result := database.DbConnection().Db().Create(&group)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, result.Error.Error())
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": gin.H{
			"id": group.ID,
		}},
	)
}

func (controller *ApplicationGroupController) Show(c *gin.Context) {
	type idParams struct { // Todo this is like form request in laravel
		ID uint `uri:"id" binding:"required,min=1"`
	}

	var params idParams

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid ID " + err.Error(),
		})
		return
	}

	var group models.ApplicationGroup

	result := database.DbConnection().Db().First(&group, params.ID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Application group not found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Something went wrong; Try later",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    &group,
	})
}

func (controller *ApplicationGroupController) Update(c *gin.Context) {
	type idParams struct { // Todo this is like form request in laravel
		ID   uint   `json:"id" binding:"required,min=1"`
		Name string `json:"name" binding:"required,min=1,max=255"`
	}

	var params idParams

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid ID " + err.Error(),
		})
		return
	}

	var group models.ApplicationGroup

	result := database.DbConnection().Db().First(&group, params.ID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Application group not found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Something went wrong; Try later",
		})

		return
	}

	group.Name = params.Name
	database.DbConnection().Db().Save(&group)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    nil,
	})
}

func (controller *ApplicationGroupController) Delete(c *gin.Context) {

	type idParams struct { // Todo this is like form request in laravel
		ID uint `uri:"id" binding:"required,min=1"`
	}

	var params idParams

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid ID " + err.Error(),
		})
		return
	}

	var group models.ApplicationGroup

	result := database.DbConnection().Db().First(&group, params.ID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Application group not found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Something went wrong; Try later",
		})

		return
	}

	database.DbConnection().Db().Delete(&group)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    nil,
	})
}
