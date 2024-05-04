package controller

import (
	libs "app/libs/firebase"
	usecase "app/usecase/commands"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Name string `json:"name"`
}

type IAuthController interface {
	Register(c echo.Context) error
}

type authController struct {
	registerUsecase usecase.IRegisterUsecase
	firebaseService libs.IFirebaseService
}

func NewAuthController(registerUc usecase.IRegisterUsecase, fs libs.IFirebaseService) IAuthController {
	return &authController{
		registerUsecase: registerUc,
		firebaseService: fs,
	}
}

func (ac *authController) Register(c echo.Context) error {
	log.Println("Running UserController.register")

	idToken := c.Request().Header.Get("Authorization")
	if idToken == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証トークンが提供されていません")
	}

	// "Bearer "プレフィックスがあれば削除
	if strings.HasPrefix(idToken, "Bearer ") {
		idToken = strings.TrimPrefix(idToken, "Bearer ")
	}

	token, err := ac.firebaseService.VerifyIDToken(c.Request().Context(), idToken)
	if err != nil {
		log.Printf("IDトークンの検証エラー: %v", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid ID token")
	}

	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Request body can't be read")
	}

	if req.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Name is required")
	}

	user, err := ac.registerUsecase.Run(usecase.RegisterUsecaseInput{
		Name:        req.Name,
		FirebaseUid: token.UID,
	})
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to register user")
	}

	return c.JSON(http.StatusOK, user)
}
