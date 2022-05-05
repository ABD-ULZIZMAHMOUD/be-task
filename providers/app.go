package providers

import (
	"be-task/helpers"
	"os"
)

func Run() {
	/**
	* init gin frame work
	* run default middleware like CROS
	 */
	r := Gin()
	/**
	* setup routes with default group
	 */
	Routing(r)
	/**
	* start app on port to change port just edit APP_PORT in
	* env file like :8000
	 */
	if os.Getenv("APP_ENV") == "local" {
		err := r.Run(":" + os.Getenv("APP_PORT_LOCAL"))
		helpers.CheckError(err)
	} else {
		err := r.Run(":" + os.Getenv("APP_PORT_PRODUCTION"))
		helpers.CheckError(err)
	}
}
