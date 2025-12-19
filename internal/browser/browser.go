package browser

import (
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func Init() *rod.Browser {
	log.Println("Launching browser")

	u := launcher.New().
		Headless(false).
		Append("--disable-blink-features=AutomationControlled").
		MustLaunch()

	return rod.New().ControlURL(u).MustConnect()
}
