package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/wpcodevo/golang-fiber-jwt/initializers"
)

func DeserializeUser(ctx *fiber.Ctx) error {
	var tokenString string
	authorization := ctx.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if ctx.Cookies("token") != "" {
		tokenString = ctx.Cookies("token")
	}

	if tokenString == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	config, _ := initializers.LoadConfig(".")
	secretKey := []byte(config.JwtSecret)

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return secretKey, nil
	})
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid user ID in token"})
	}
	userName, ok := claims["name"].(string)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid user ID in token"})
	}
	// Simpan ID pengguna ke dalam konteks Fiber
	ctx.Locals("userID", userID)
	ctx.Locals("userName", userName)

	return ctx.Next()
}
