package utils

import "regexp"

func IsValidateUrl(url string) bool {
	urlRegex := `^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?(\?[;&a-z\d%_.~+=-]*)?(\#[-a-z\d_]*)?$`
	re := regexp.MustCompile(urlRegex)
	return re.MatchString(url)
}
