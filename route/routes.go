package route

import (
	"github.com/ak98neon/authserver/controller"
	"github.com/ak98neon/authserver/secret"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
)

func Routes() *mux.Router {

	r := mux.NewRouter()
	r.Use(CommonMiddleware)

	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.Handle("/forms", jwtMiddleware().Handler(controller.GetAllForms)).Methods("GET")
	r.Handle("/forms/approve", jwtMiddleware().Handler(controller.ApproveForm)).Methods("POST")
	r.Handle("/forms/reject", jwtMiddleware().Handler(controller.RejectForm)).Methods("POST")
	return r
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}

func jwtMiddleware() *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return secret.MySigningKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}
