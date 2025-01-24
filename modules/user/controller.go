package user

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateUserController(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	createdUser, err := CreateUser(user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H {
		"status":  "success",
		"message": "create user success",
		"data":    createdUser,
	})
}

func GetAllUserController(ctx *gin.Context) {
	user, err := GetAllUser()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"status":  "success",
		"message": "get all user success",
		"data":    user,
	})
}

func GetUserByIdController(ctx *gin.Context) {
	getId := ctx.Param("id")

	id, err := strconv.Atoi(getId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	user, err := GetUserById(id)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
				"status":  "fail",
				"message": err.Error(),
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
				"status":  "fail",
				"message": err.Error(),
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, gin.H {
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	updatedUser, err := UpdateUserById(id, user)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
				"status":  "fail",
				"message": err.Error(),
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
				"status":  "fail",
				"message": err.Error(),
			})
		}

		return
	}

	ctx.JSON(http.StatusCreated, gin.H {
		"status":  "success",
		"message": fmt.Sprintf("update user by id \"%v\" success", id),
		"data":    updatedUser,
	})
}

func DeleteUserByIdController(ctx *gin.Context) {
	getId := ctx.Param("id")

	id, err := strconv.Atoi(getId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"status":  "fail",
			"message": err.Error(),
		})

		return
	}

	deletedUser, err := DeleteUserById(id)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
				"status":  "fail",
				"message": err.Error(),
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
				"status":  "fail",
				"message": err.Error(),
			})
		}

		return
	}

	ctx.JSON(http.StatusCreated, gin.H {
		"status":  "success",
		"message": fmt.Sprintf("delete user by id \"%v\" success", id),
		"data":    deletedUser,
	})
}