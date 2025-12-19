package profile

import (
	"log"

	"github.com/go-rod/rod"
	"linkedin/internal/config"
)

func SendConnection(page *rod.Page, cfg config.Config, profileURL string) {
	log.Println("Visiting profile:", profileURL)

	if !cfg.AutomationEnabled {
		log.Println("[SAFE MODE] Would send connection request")
		return
	}

	log.Println("Automation enabled (not used in POC)")
}
