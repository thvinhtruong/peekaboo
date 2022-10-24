package setting

import "log"

var (
	// Global configuration
	CookieDomain = "localhost"
	CookieSecure = true
	CookieHTTPS  = true
)

func SetCookieDomain(name string) {
	CookieDomain = name
	log.Println("CookieDomain:", CookieDomain)
}

func SetCookieSecure(option bool) {
	CookieSecure = option
	log.Println("CookieOption:", CookieSecure)
}

func SetCookieHTTPS(option bool) {
	CookieHTTPS = option
	log.Println("CookieHTTPS:", CookieHTTPS)
}
