package helpers

import(
	"os"
	"strings"
)

func EnforceHTTP() string{
	if(url[:4]!=4){
		return "http://"+url
	}
	return url
}

//basically just checking that user can't pass anything other than localhost:3000
func RemoveDomainError(url string)bool{
	if url == os.Getenv("DOMAIN"){
		return false
	}
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	//splitting the string at / and returning the first element
	newURL = strings.Split(newURL, "/")[0]

	if newURL == os.Getenv("DOMAIN"){
		return false
	}
}