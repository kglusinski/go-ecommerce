package security

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	resErrBadCredentials = ErrorResponse{Message: errBadCredentials.Error()}
	resErrBadRequest     = ErrorResponse{Message: "incorrect request body"}
)

type SecurityService interface {
	Login(email string, password string) (string, error)
	Register(email string, password string) error
}

type SecurityController struct {
	securityService SecurityService
}

func NewSecurityController(service SecurityService) *SecurityController {
	return &SecurityController{
		securityService: service,
	}
}

type CredentialRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Token string `json:"token"`
}

func (s *SecurityController) Login(e echo.Context) error {
	var creds CredentialRequest
	err := e.Bind(&creds)
	if err != nil {
		return e.JSON(http.StatusBadRequest, resErrBadRequest)
	}

	token, err := s.securityService.Login(creds.Email, creds.Password)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, resErrBadCredentials)
	}

	return e.JSON(http.StatusOK, SuccessResponse{Token: token})
}

func (s *SecurityController) Register(e echo.Context) error {
	var creds CredentialRequest
	err := e.Bind(&creds)
	if err != nil {
		return e.JSON(http.StatusBadRequest, resErrBadRequest)
	}

	err = s.securityService.Register(creds.Email, creds.Password)
	if err != nil {
		return e.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	return e.NoContent(http.StatusCreated)
}
