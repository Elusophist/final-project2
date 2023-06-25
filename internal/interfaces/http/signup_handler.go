package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/the-go-dragons/final-project2/internal/domain"
)

type SignupRequest struct {
	Username string
	Password string
}

func (uh *UserHandler) Signup(c echo.Context) error {
	var request SignupRequest

	// Check the body data
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Message: "Invalid body request"})

	}
	if request.Username == "" || request.Password == "" {
		return c.JSON(http.StatusBadRequest, Response{Message: "Missing required fields"})
	}

	_, err = uh.userUsecase.GetUserByUsername(request.Username)
	if err == nil {
		return c.JSON(http.StatusConflict, Response{Message: "User already exists with the given username"})
	}

	user := domain.User{
		Username: request.Username,
		Password: request.Password,
	}

	_, err = uh.userUsecase.CreateUser(&user)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Cant create user"})
	}

	return c.JSON(http.StatusOK, Response{Message: "Created"})
}