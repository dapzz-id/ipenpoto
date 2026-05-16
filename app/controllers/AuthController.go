package controllers

import (
	"strings"

	"ipenpoto/app/requests"
	"ipenpoto/app/services"
	"ipenpoto/app/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController(
	authService *services.AuthService,
) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (ctrl *AuthController) Login(c *fiber.Ctx) error {
	var req requests.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.Error(
			c,
			400,
			"Invalid request body",
			err.Error(),
		)
	}

	if err := utils.ValidateRequest.Struct(req); err != nil {
		return utils.Error(
			c,
			400,
			"Validation failed",
			utils.FormatValidationError(err),
		)
	}

	user, accessToken, refreshToken, err :=
		ctrl.AuthService.Login(
			req.Username,
			req.Password,
		)

	if err != nil {
		statusCode := 401
		if strings.Contains(err.Error(), "Email not verified") {
			statusCode = 403
		}
		return utils.Error(
			c,
			statusCode,
			err.Error(),
			nil,
		)
	}

	return utils.Success(
		c,
		"Login successful",
		fiber.Map{
			"user": user,
			"tokens": fiber.Map{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			},
		},
	)
}

func (ctrl *AuthController) Register(c *fiber.Ctx) error {
	var req requests.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.Error(
			c,
			400,
			"Invalid request body",
			err.Error(),
		)
	}

	if err := utils.ValidateRequest.Struct(req); err != nil {
		return utils.Error(
			c,
			400,
			"Validation failed",
			utils.FormatValidationError(err),
		)
	}

	user, err := ctrl.AuthService.Register(
		req.Name,
		req.Username,
		req.Email,
		req.Password,
	)

	if err != nil {
		return utils.Error(
			c,
			400,
			err.Error(),
			nil,
		)
	}

	return utils.Success(
		c,
		201,
		"Registration successful. Check your email to verify your account.",
		fiber.Map{
			"id":    user.ID,
			"email": user.Email,
		},
	)
}

func (ctrl *AuthController) VerifyEmail(c *fiber.Ctx) error {
	token := c.Query("token")
	if token == "" {
		return utils.Error(
			c,
			400,
			"Verification token is required",
			nil,
		)
	}

	err := ctrl.AuthService.VerifyEmail(token)
	if err != nil {
		statusCode := 400
		errorMsg := err.Error()

		if strings.Contains(errorMsg, "expired") || strings.Contains(errorMsg, "invalid") {
			statusCode = 400
		} else if strings.Contains(errorMsg, "already verified") {
			statusCode = 400
		}

		return utils.Error(
			c,
			statusCode,
			errorMsg,
			nil,
		)
	}

	return utils.Success(
		c,
		"Email verified successfully. You can now login.",
		fiber.Map{},
	)
}
