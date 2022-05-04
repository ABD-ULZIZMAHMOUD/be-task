package customers

import (
	"be-task/config"
	"be-task/helpers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func filter(g *gin.Context) *gorm.DB {
	db := config.DB
	db = filterByState(helpers.GetParamFromRequest(g, "state"), db)
	_, db = filterByCountry(helpers.GetParamFromRequest(g, "country"), db)
	return db
}

func filterByCountry(country string, db *gorm.DB) (bool, *gorm.DB) {
	// no country
	if country == "" {
		return true, db
	}
	// not valid country
	regex := helpers.GetRegexByCountry(country)
	if regex == "" {
		return false, db
	}
	db = db.Where(`phone REGEXP ?`, regex)
	return true, db
}

func filterByState(state string, db *gorm.DB) *gorm.DB {
	var aggORed string
	var aggORedValues []interface{}
	for _, regex := range helpers.Countries {
		aggORed = appendORedQuery(aggORed, "phone REGEXP ?")
		aggORedValues = append(aggORedValues, regex)
	}
	if state == "valid" {
		db = db.Where(aggORed, aggORedValues...)
	} else if state == "not_valid" {
		db = db.Not(aggORed, aggORedValues...)
	}
	return db
}

func appendORedQuery(aggQuery, query string) string {
	if len(aggQuery) > 0 {
		aggQuery += " or "
	}
	aggQuery += query
	return aggQuery
}
