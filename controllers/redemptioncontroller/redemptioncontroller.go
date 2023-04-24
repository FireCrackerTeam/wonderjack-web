package redemptioncontroller

import (
	"encoding/json"
	"net/http"

	"github.com/FireCrackerTeam/wonderjack-web/helper"
	"github.com/FireCrackerTeam/wonderjack-web/models"
)

func ClaimRedemtionCode(w http.ResponseWriter, r *http.Request){
	//get redemption code from body
	var codeInput models.RedemptionCodes
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&codeInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// get user id from token
	userId, err := helper.JWTClaim(r)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	codeInput.UserId = userId

	// check if the redemption code has been used with current user
	var redemptionCode models.RedemptionCodes
	models.DB.Where("user_id = ?", codeInput.UserId).Where("code = ?", codeInput.Code).Find(&redemptionCode)
	if redemptionCode != (models.RedemptionCodes{}) {
		response := map[string]string{"message": "Current user has claimed the redemption code!"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	

	// check if the redemption code has been used twice
	var count int64
	models.DB.Table("redemption_codes").Group("code").Count(&count)
	if count < 2 {
		// insert to database
		if err := models.DB.Create(&codeInput).Error; err != nil {
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	} else { // if redemption code has been used twice
		response := map[string]string{"message": "Redemption code has been used twice!"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := map[string]string{"message": "Redemption code code successfully claimed!"}
	helper.ResponseJSON(w, http.StatusOK, response)
}