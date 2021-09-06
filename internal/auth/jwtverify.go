package auth

import (
	"errors"
	"fmt"
	"github.com/anish-yadav/lms-api/internal/constants"
	"github.com/anish-yadav/lms-api/internal/pkg/permission"
	"github.com/anish-yadav/lms-api/internal/pkg/user"
	webresponse "github.com/anish-yadav/lms-api/internal/pkg/webservice/response"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	authHeaderKey = "Authorization"
	bearer        = "Bearer "
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		token := r.Header.Get(authHeaderKey)
		tokenString := strings.TrimPrefix(token, bearer)
		tokenString = strings.Trim(tokenString, " ")
		log.Debugf("auth.Middleware: token : %s", tokenString)
		if len(tokenString) == 0 {
			webresponse.RespondWithError(w, http.StatusBadRequest, constants.InvalidToken)
			return
		}
		t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New(constants.InvalidSigningMethod)
			}

			if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
				return nil, errors.New(constants.TokenExpired)
			}
			// todo need to autogenerate this secret on first run
			return []byte("secret"), nil
		})
		if err != nil {
			log.Debugf("auth.Middleware: failed to parse/verify : %s", err.Error())
			webresponse.RespondWithError(w, http.StatusUnauthorized, constants.InvalidToken)
			return
		}
		claims := t.Claims.(jwt.MapClaims)
		data := claims["data"].(map[string]interface{})
		path := strings.TrimPrefix(r.URL.Path, "/api/v1")
		reqPermission := RoutePermissionMap[path]
		log.Debugf("permission required for %s : %s", r.URL.Path, reqPermission)
		userID := data["user.id"]
		currUser := user.GetUserById(fmt.Sprintf("%s", userID))
		userPermission := permission.GetPermissionByName(currUser.Type)

		if !userPermission.HasPermission(reqPermission) {
			webresponse.RespondWithError(w, http.StatusForbidden, constants.Forbidden)
			return
		}

		if t.Valid {
			next.ServeHTTP(w, r)
			return
		}

		webresponse.RespondWithError(w, http.StatusUnauthorized, constants.InvalidToken)
		return
	})
}
