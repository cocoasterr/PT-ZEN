package controllers

import (
	"time"

	"github.com/cocoasterr/kalbe_test/models"
	"github.com/cocoasterr/kalbe_test/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"github.com/wpcodevo/golang-fiber-jwt/initializers"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		Service: service,
	}
}

func (c *UserController) RegisterUser(ctx *fiber.Ctx) error {
	var userInput models.User
	if err := ctx.BodyParser(&userInput); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return err
	}
	userData, _ := c.Service.FindUser(ctx.Context(), userInput.Email)
	if userData != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not Registered!"})
	}

	password := userInput.Password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithError(err).Error("Error hashing password")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	userInput.Password = string(hashedPassword)

	trx, err := c.Service.Repo.BeginTransaction()
	if err != nil {
		logrus.WithError(err).Error("Error beginning transaction")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	userMap := map[string]interface{}{
		"id":           userInput.Id,
		"firstname":    userInput.FirstName,
		"lastname":     userInput.LastName,
		"phone_number": userInput.PhoneNumber,
		"email":        userInput.Email,
		"password":     userInput.Password,
	}

	err = c.Service.CreateService(ctx.Context(), trx, userMap, "")
	if err != nil {
		logrus.WithError(err).Error("Error creating User")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if commitErr := c.Service.Repo.CommitTransaction(trx); commitErr != nil {
		logrus.WithError(commitErr).Errorf("Error committing transaction")
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	config, _ := initializers.LoadConfig(".")
	secretKey := []byte(config.JwtSecret)

	var body models.LoginUser
	if err := ctx.BodyParser(&body); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return err
	}

	userData, _ := c.Service.FindUser(ctx.Context(), body.Email)
	if userData == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found!"})
	}

	hashedPassword := userData["password"].(string)

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(body.Password)); err != nil {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid username or password",
		})
		return nil
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userData["id"]
	claims["name"] = userData["firstname"].(string) + " " + userData["lastname"].(string)
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	s, err := token.SignedString([]byte(secretKey))
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token": s, "message": "login successfully"})
}

func (c *UserController) Profile(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(string)
	userData, _ := c.Service.FindByService(ctx.Context(), "id", userID)
	if userData == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found!"})
	}
	respUser := models.RespUser{
		FirstName:   userData["firstname"].(string),
		LastName:    userData["lastname"].(string),
		PhoneNumber: userData["phone_number"].(string),
		Email:       userData["email"].(string),
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": respUser})
}

func (c *UserController) ResetPassword(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(string)
	var userInput models.User
	if err := ctx.BodyParser(&userInput); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return err
	}
	password := userInput.Password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithError(err).Error("Error hashing password")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	userInput.Password = string(hashedPassword)

	userMap := map[string]interface{}{
		"password": userInput.Password,
	}
	trx, err := c.Service.Repo.BeginTransaction()
	if err != nil {
		logrus.WithError(err).Error("Error beginning transaction")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	if err = c.Service.UpdateService(ctx.Context(), trx, userMap, userID, ""); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if commitErr := c.Service.Repo.CommitTransaction(trx); commitErr != nil {
		logrus.WithError(commitErr).Errorf("Error committing transaction")
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "user successfully updated!"})
}

func (c *UserController) Logout(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
	ctx.Set("Authorization", "")
	ctx.Locals("userID", nil)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logout successful"})
}
