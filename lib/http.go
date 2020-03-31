package lib

import (
	"fmt"
	"net/http"
	
	"github.com/laplaceon/cfbypass"
)

const UserAgent = "Mozilla/5.0 (Windows NT 10.0; rv:50.0) Gecko/20100101 Firefox/50.0"

var HTTPClient *http.Client

func init() {
	HTTPClient = &http.Client{}
}

func GetCFCookies(url string) []*http.Cookie {
	fmt.Println(cfbypass.GetTokens(url, UserAgent, ""))

	return []*http.Cookie{}
}