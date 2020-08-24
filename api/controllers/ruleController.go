package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	model "github.com/aperea/go-mlproxy/api/models"

)


func RulesHandler(c *gin.Context) {

	NrawBody, _ := c.GetRawData()
	// Marshall the requrest body
	var thisItem model.IP
	json.Unmarshal([]byte(string(NrawBody)), &thisItem)

	result := validateRequest(thisItem)
	if result != true {
		c.JSON(http.StatusOK, gin.H{})
		return
	}else{
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}
}

func validateRequest(thisItem model.IP) bool {

	rules := model.GetRulesByIP(thisItem)

	if rules != nil {
		for _, rule := range rules {
			if rule.RuleType == 1  || rule.RuleType == 2 {
				return true
			}		
		}
	}
	return false

}
