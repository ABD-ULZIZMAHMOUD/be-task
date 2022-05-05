package main

import (
	_ "be-task/docs"
	"be-task/helpers"
	"be-task/providers"
)

// @title jumia BE task
// @version 1.0
// @description The APIs of BE task
// @termsOfService http://swagger.io/terms/
// @BasePath /
func main() {
	//connect to database
	providers.InitDatabase()
	//load countries to memory
	helpers.LoadCountries()
	//run gin server
	providers.Run()
}
