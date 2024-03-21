package service

import (
	"ecommerce/constants"
	"fmt"
	"math/rand"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(email, customerid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      email,
		"customerid": customerid, // Set the customerid claim
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(constants.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractCustomerID(jwtToken string, secretKey string) (string, error) {
	// Parse the JWT token
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err // Handle token parsing errors
	}

	// Check if the token is valid
	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Extract the customer ID from the claims
			customerID, ok := claims["customerid"].(string)
			if ok {
				return customerID, nil
			}
		}
	}

	return "", fmt.Errorf("invalid or expired JWT token")
}


func GenerateUniqueCustomerID() string {
	// Implement your logic to generate a unique customer ID (e.g., UUID, timestamp, etc.)
	// For example, you can use a combination of timestamp and random characters
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), GetRandomString(6))
}

func GenerateUniqueAdminID() string {
	// Implement your logic to generate a unique customer ID (e.g., UUID, timestamp, etc.)
	// For example, you can use a combination of timestamp and random characters
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), GetRandomString(8))
}

// Custom function to generate random characters (for demonstration purposes)
func GetRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateOTP(length int) string {
	const charset = "0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}