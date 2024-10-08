package middleware

import (
	"CloudStorage/pkg/response"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/context"
)

func (mw *Middleware) JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp response.Response

		token := r.Header.Get("Authorization")

		tokenList := strings.Split(token, " ")

		if len(tokenList) != 2 {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		token = tokenList[1]

		tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(mw.config.JwtSecretKey), nil
		})

		if err != nil {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		if !tokenObj.Valid {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		tokenMap, ok := tokenObj.Claims.(jwt.MapClaims)

		if !ok {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		expireDate, ok := tokenMap["exp"].(float64)

		if !ok {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		if time.Now().Unix() > int64(expireDate) {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		userID, ok := tokenMap["user_id"].(float64)

		if !ok {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		roleID, ok := tokenMap["role_id"].(float64)

		if !ok {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		if strings.Contains(r.URL.String(), "/admin/") && roleID != 1 {
			resp = response.Forbidden
			resp.WriteJSON(w)
			return
		}

		_, err = mw.service.GetUserByID(int(userID))

		if err != nil {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		context.Set(r, "role_id", int64(roleID))
		context.Set(r, "user_id", int64(userID))

		next.ServeHTTP(w, r)
	})
}
