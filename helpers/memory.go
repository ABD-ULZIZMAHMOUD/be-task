package helpers

var Countries map[string]string


func LoadCountries() {
	Countries = make(map[string]string)
	Countries["Cameroon"] = "\\(237\\) ?[2368]\\d{7,8}$"
	Countries["Ethiopia"] = "\\(251\\) ?[1-59]\\d{8}$"
	Countries["Morocco"] = "\\(212\\) ?[5-9]\\d{8}$"
	Countries["Mozambique"] = "\\(258\\) ?[28]\\d{7,8}$"
	Countries["Uganda"] = "\\(256\\) ?\\d{9}$"
}

func GetRegexByCountry(countryName string) string {
	return Countries[countryName]
}
