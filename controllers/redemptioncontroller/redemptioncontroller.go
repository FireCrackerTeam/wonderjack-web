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
		response := map[string]string{
			"success": "failed",
			"message": "Current user has claimed the redemption code!",
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	
	// check if the redemption code has been redeemed by other user
	models.DB.Where("NOT user_id = ?", codeInput.UserId).Where("code = ?", codeInput.Code).Find(&redemptionCode)
	if redemptionCode != (models.RedemptionCodes{}) {
		response := map[string]string{
			"success": "failed",
			"message": "Other user has claimed the redemption code!",
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// check if user already redeemed the code
	models.DB.Where("user_id = ?", codeInput.UserId).Find(&redemptionCode)
	if redemptionCode != (models.RedemptionCodes{}) {
		response := map[string]string{
			"success": "success",
			"message": "Already claimed before!",
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Check if the redemption code is available
	var availableCode models.AvailableCodes
	models.DB.Where("code = ?", codeInput.Code).Find(&availableCode)
	if availableCode != (models.AvailableCodes{}) {
		// insert to database
		if err := models.DB.Create(&codeInput).Error; err != nil {
			response := map[string]string{
				"success": "failed",
				"message": err.Error(),
			}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	} else { // if redemption code unavailable
		response := map[string]string{
			"success": "failed",
			"message": "Redemption code unavailable",
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := map[string]string{
		"success": "success",
		"message": "Redemption code code successfully claimed!",
	}
	helper.ResponseJSON(w, http.StatusOK, response)
}
