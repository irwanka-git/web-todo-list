package middleware

import (
	"context"
	"fmt"
	"irwanka/webtodolist/helper"
	"net/http"

	"github.com/go-chi/jwtauth"
)

type customMiddleware struct{}

type UserMiddleware interface {
	SetValueContext(next http.Handler) http.Handler
}

type ValueContext struct {
	m map[string]interface{}
}

func NewUserMiddleware() UserMiddleware {
	return &customMiddleware{}
}

func (v ValueContext) GetValueContext(key string) interface{} {
	return v.m[key]
}

func (*customMiddleware) SetValueContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, claims, errToken := jwtauth.FromContext(r.Context())

		if errToken != nil {
			fmt.Println("Gagal Ambil sub dari token jwt ")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		sub := fmt.Sprintf("%v", claims["sub"])
		id := fmt.Sprintf("%v", claims["id"])

		v := ValueContext{map[string]interface{}{
			helper.USER_UUID: sub,
			helper.USER_ID:   id,
		}}
		ctx := context.WithValue(r.Context(), helper.VALUE_CONTEXT, v)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
