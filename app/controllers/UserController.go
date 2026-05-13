package controllers

import (
	"ipenpoto/app/repositories"
	"ipenpoto/app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController struct {
	UserRepo *repositories.UserRepository
}

func NewUserController(
	userRepo *repositories.UserRepository,
) *UserController {
	return &UserController{
		UserRepo: userRepo,
	}
}

func (ctrl *UserController) GetProfile(c *fiber.Ctx) error {
	userIDStr := c.Locals("user_id")
	if userIDStr == nil {
		return utils.Error(
			c,
			401,
			"Unauthorized - Missing user_id",
			nil,
		)
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		return utils.Error(
			c,
			401,
			"Unauthorized - Invalid user_id",
			nil,
		)
	}

	user, err := ctrl.UserRepo.FindByID(userID)
	if err != nil {
		return utils.Error(
			c,
			404,
			"User not found",
			nil,
		)
	}

	return utils.Success(
		c,
		"Success",
		fiber.Map{
			"id":                user.ID,
			"name":              user.Name,
			"username":          user.Username,
			"email":             user.Email,
			"role":              string(user.Role),
			"avatar_url":        user.AvatarURL,
			"email_verified_at": user.EmailVerifiedAt,
			"status":            string(user.Status),
			"created_at":        user.CreatedAt,
			"updated_at":        user.UpdatedAt,
		},
	)
}
