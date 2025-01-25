package user

import (
	"fmt"
	"formative-14/commons"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func indexController(ctx *gin.Context) {
	type AppRoutes struct {
		Route1 string
		Route2 string
		Route3 string
		Route4 string
		Route5 string
	}

	var endpoint = commons.ENDPOINT
	var repo = commons.REPOSITORY

	var indexDescription = struct {
		Details    string
		Repository string
		Endpoint   string
		Routes     AppRoutes
	}{
		Details:    "Fungsionalitas sama, tetapi hanya mengganti nama model person menjadi user",
		Repository: repo,
		Endpoint:   endpoint,
		Routes: AppRoutes{
			Route1: fmt.Sprintf("Get All User: GET %s/users", endpoint),
			Route2: fmt.Sprintf("Get User By Id: GET %s/users/:id", endpoint),
			Route3: fmt.Sprintf("Create User: POST %s/users", endpoint),
			Route4: fmt.Sprintf("Update User By Id: PUT %s/users/:id", endpoint),
			Route5: fmt.Sprintf("Delete User By Id: DELETE %s/users/:id", endpoint),
		},
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "view index success",
		"data":    indexDescription,
	})
}

func CreateUserController(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	createdUser, err := CreateUser(user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "create user success",
		"data":    createdUser,
	})
}

func GetAllUserController(ctx *gin.Context) {
	user, err := GetAllUser()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "get all user success",
		"data":    user,
	})
}

func GetUserByIdController(ctx *gin.Context) {
	getId := ctx.Param("id")

	id, err := strconv.Atoi(getId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	user, err := GetUserById(id)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("get user by id \"%v\" success", id),
		"data":    user,
	})
}

func UpdateUserByIdController(ctx *gin.Context) {
	var user User

	getId := ctx.Param("id")

	id, err := strconv.Atoi(getId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	updatedUser, err := UpdateUserById(id, user)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("update user by id \"%v\" success", id),
		"data":    updatedUser,
	})
}

func DeleteUserByIdController(ctx *gin.Context) {
	getId := ctx.Param("id")

	id, err := strconv.Atoi(getId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	deletedUser, err := DeleteUserById(id)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("delete user by id \"%v\" success", id),
		"data":    deletedUser,
	})
}
