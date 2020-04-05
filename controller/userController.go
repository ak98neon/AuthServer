package controller

import (
	"encoding/json"
	"fmt"
	"github.com/ak98neon/authserver/model"
	"github.com/ak98neon/authserver/util"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var db = util.ConnectDb()

func Login(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := FindOne(user.Username, user.Password)
	json.NewEncoder(w).Encode(resp)
}

func FindOne(username, password string) map[string]interface{} {
	user := &model.User{}

	if err := db.Where("Username = ?", username).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Username not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Hour * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &model.Token{
		UserID: user.ID,
		Name:   user.Username,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString(mySigningKey)
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}
