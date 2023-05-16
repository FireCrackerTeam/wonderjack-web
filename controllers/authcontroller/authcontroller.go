package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FireCrackerTeam/wonderjack-web/config"
	"github.com/FireCrackerTeam/wonderjack-web/helper"
	"github.com/FireCrackerTeam/wonderjack-web/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request){
	// get json data
	var userInput models.Users
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// get user data from email
	var user models.Users
	if err := models.DB.Where("user_email = ?", userInput.UserEmail).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Wrong email or password"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Wrong email or password"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// create jwt token
	expTime := time.Now().Add(time.Hour * 24 * 60)
	claims := &config.JWTClaim {
		UserId: user.UserId,
		Email: user.UserEmail,
		RegisteredClaims: jwt.RegisteredClaims {
			Issuer: "wonderjack-web",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// declare algorithm for signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	var redemptionCode models.RedemptionCodes
	models.DB.Where("user_id = ?", user.UserId).Find(&redemptionCode)
	if redemptionCode != (models.RedemptionCodes{}) {
		// response := map[string]string{
		// 	"message": "login successful",
		// 	"user_id": user.userId,
		// 	"token": token,
		// 	"redeem": "success",
		// }
		response := map[string]interface{}{
			"message": "Login Successful",
			"user_id": user.UserId,
			"token": token,
			"redeem": true,
		}
		helper.ResponseJSON(w, http.StatusOK, response)
		return
	}

	// response := map[string]string{
	// 	"message": "login successful",
	// 	"token": token,
	// 	"redeem": "failed",
	// }
	response := map[string]interface{}{
		"message": "Login Successful",
		"user_id": user.UserId,
		"token": token,
		"redeem": false,
	}
	
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	// get json data
	var userInput models.Users
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// hash password using bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert to database
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	expTime := time.Now().Add(time.Hour * 24 * 60)
	claims := &config.JWTClaim {
		UserId: userInput.UserId,
		Email: userInput.UserEmail,
		RegisteredClaims: jwt.RegisteredClaims {
			Issuer: "wonderjack-web",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// declare algorithm for signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	// response := map[string]string{
	// 	"message": "success",
	// 	"user_id": user.userId,
	// 	"token": token,
	// }
	response := map[string]interface{}{
		"message": "Register Successful",
		"user_id": userInput.UserId,
		"token": token,
	}

	helper.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request){
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: "",
		HttpOnly: true,
		MaxAge: -1,
	})

	response := map[string]string{"message": "logout successful"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
