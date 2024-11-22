package handler

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

func (h *Handler) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			accessToken, err := GetBearerToken(r)
			if err != nil {
				log.Error().Msgf("failed GetBearerToken: %v", err)
				h.HandleError(w, http.StatusUnauthorized, err.Error())
				return
			}

			token, err := jwt.Parse(
				accessToken, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}

					return h.auc.GetJWTToken(), nil
				},
			)

			log.Debug().Msgf("token is valid: %v", token.Valid)

			if err != nil || !token.Valid {
				h.HandleError(w, http.StatusUnauthorized, err.Error())
				return
			}

			id := token.Claims.(jwt.MapClaims)["id"].(string)
			err = h.auc.AuthorizeUser(r.Context(), model.AuthorizeUserV1Request{UserID: id, RoutePath: r.URL.Path, Method: r.Method})
			if err != nil {
				log.Error().Msgf("failed h.auc.Authorize: %v", err)
				h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
			}

			r.Header.Set(HeaderUserID, id)

			next.ServeHTTP(w, r)
		},
	)
}

func GetBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get(HeaderAuthorization)
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is missing")
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return "", fmt.Errorf("authorization header does not start with 'Bearer '")
	}

	token := strings.TrimPrefix(authHeader, bearerPrefix)
	return token, nil
}
