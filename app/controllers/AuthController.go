package controllers

import (
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
		return utils.Error(
			c,
			401,
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

	user, accessToken, refreshToken, err :=
		ctrl.AuthService.Register(
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
		"User registered successfully",
		fiber.Map{
			"user": user,
			"tokens": fiber.Map{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			},
		},
	)
}
