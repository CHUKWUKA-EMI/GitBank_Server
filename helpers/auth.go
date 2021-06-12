package helpers

import (
	"log"
	"os"
	"strings"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	jwt.StandardClaims
	ID uint `gorm:"primaryKey"`
}

var loadErr = godotenv.Load()

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateAccessToken(id uint, name string, email string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	if loadErr != nil {
		log.Fatal(loadErr.Error())
	}
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["userId"] = id
	claims["email"] = email
	claims["role"] = role
	claims["Issuer"] = "GitBank"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	signedToken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateVerificationToken(id uint) (string, error) {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "ChukwukaBank",
			IssuedAt:  time.Now().Unix(),
		},
		ID: id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ExtractToken(c *fiber.Ctx) string {
	bearerToken := c.Request().Header.Peek("Authorization")

	token := strings.Split(string(bearerToken), " ")
	if len(token) == 2 {
		return token[1]
	}

	return " "
}

func VerifyToken(c *fiber.Ctx) error {
	tokenString := ExtractToken(c)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok && !token.Valid {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error":   true,
			"message": "Couldn't parse claims",
		})
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Token has expired",
		})
	}
	c.Locals("userId", claims.ID)
	return c.Next()
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
