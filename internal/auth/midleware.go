package auth

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddlewareAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractTokenFromHeader(r)
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := verifyToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			role, ok := claims["role"].(string)
			if !ok || role != "admin" {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Передаем управление следующему обработчику
		next.ServeHTTP(w, r)
	})
}

func extractTokenFromHeader(r *http.Request) string {
	// Извлечение токена из заголовка Authorization или других мест
	// Пример:
	// Authorization: Bearer <token>
	token := r.Header.Get("Authorization")
	if token != "" {
		// Пример: удаление префикса "Bearer " и возврат токена
		return token[len("Bearer "):]
	}
	return ""
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	// Пример: Верификация токена с использованием публичного ключа или секрета
	// Для ES256 с публичным ключом (как в вашем случае с ecdsaPrivateKey.PublicKey)
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверка метода подписи
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		// Возвращаем публичный ключ для проверки подписи
		return &ecdsaPrivateKey.PublicKey, nil
	})
}
