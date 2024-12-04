package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/patyukin/mbs-api-gateway/internal/model"
	"github.com/rs/zerolog/log"
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
						return nil, fmt.Errorf("%w: %v", ErrUnexpectedSigningMethod, token.Header["alg"])
					}

					return h.auc.GetJWTToken(), nil
				},
			)

			log.Debug().Msgf("token is valid: %v", token.Valid)

			if err != nil || !token.Valid {
				h.HandleError(w, http.StatusUnauthorized, err.Error())
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				log.Error().Msg("failed to get claims from token")
				h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			idValue, idExists := claims["id"]
			roleValue, roleExists := claims["role"]
			if !idExists || !roleExists {
				log.Error().Msg("failed to get id or role from claims")
				h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			id, ok := idValue.(string)
			if !ok {
				log.Error().Msg("failed to get id from claims")
				h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			role, ok := roleValue.(string)
			if !ok {
				log.Error().Msg("failed to get role from claims")
				h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			err = h.auc.AuthorizeUserV1UseCase(
				r.Context(), model.AuthorizeUserV1Request{
					UserID:    id,
					RoutePath: r.URL.Path,
					Method:    r.Method,
				},
			)
			if err != nil {
				log.Error().Msgf("failed h.auc.AuthorizeUserV1UseCase: %v", err)
				h.HandleError(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			r.Header.Set(HeaderUserID, id)
			r.Header.Set(HeaderUserRole, role)

			next.ServeHTTP(w, r)
		},
	)
}

func GetBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get(HeaderAuthorization)
	if authHeader == "" {
		return "", ErrMissingAuthorizationHeader
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return "", ErrInvalidAuthorizationHeaderFormat
	}

	token := strings.TrimPrefix(authHeader, bearerPrefix)
	return token, nil
}
