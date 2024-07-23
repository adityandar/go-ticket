package controller

import (
	"go-ticket/constants"
	"go-ticket/database"
	"go-ticket/helpers"
	"go-ticket/repository"
	"go-ticket/request"
	"go-ticket/structs"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterAsOrganizer(ctx *gin.Context) {
	var registerRequestBody request.RegisterRequestBody
	var organizerParam structs.Organizer

	err := ctx.BindJSON(&registerRequestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"loc":   "user binding",
		})
		return
	}

	passwordHash, err := helpers.HashAndSalt([]byte(registerRequestBody.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	organizerParam = structs.Organizer{
		CompanyName: registerRequestBody.CompanyName,
		User: structs.User{
			FullName: registerRequestBody.FullName,
			Email:    registerRequestBody.Email,
			Password: passwordHash,
			Role:     constants.Organizer,
		},
	}
	log.Println(organizerParam)

	err = repository.InsertOrganizer(database.DbConnection, organizerParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"loc":   "insert organizer",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Insert Organizer",
	})
	return
}

func LoginAsOrganizer(ctx *gin.Context) {
	var loginRequestBody request.LoginRequestBody
	var organizer structs.Organizer
	err := ctx.BindJSON(&loginRequestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"loc":   "binding param",
		})
		return
	}
	organizer, err = repository.GetOrganizerByEmail(database.DbConnection, loginRequestBody.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"loc":   "get organizer by email",
		})
		return
	}
	if organizer.User.Email == "" && organizer.User.Password == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Email not found",
		})
		return
	}
	isPasswordCorrect := helpers.ComparePasswords(organizer.User.Password, []byte(loginRequestBody.Password))

	if isPasswordCorrect {
		ctx.JSON(http.StatusOK, gin.H{
			"result": "You are logged in",
			"body":   organizer,
		})
	} else {

		ctx.JSON(http.StatusUnauthorized, gin.H{
			"result": "Password is wrong",
		})
	}
	return

}

// not working
func UpdateOrganizer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"loc":   "convert id",
		})
		return
	}
	var updateOrganizerRequestBody request.UpdateOrganizerRequestBody
	err = ctx.BindJSON(&updateOrganizerRequestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"loc":   "binding param",
		})
		return
	}
	updatedOrganizer := structs.Organizer{
		UserId:      id,
		CompanyName: updateOrganizerRequestBody.CompanyName,
		User: structs.User{
			Id:       id,
			FullName: updateOrganizerRequestBody.FullName,
			Email:    updateOrganizerRequestBody.Email,
		},
	}

	err = repository.UpdateOrganizer(database.DbConnection, updatedOrganizer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"loc":   "update organizer",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": "Update success",
	})
	return

}

func DeleteOrganizer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"loc":   "convert id",
			})
			return
		}
	}
	err = repository.DeleteOrganizer(database.DbConnection, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Organizer",
	})
	return
}
