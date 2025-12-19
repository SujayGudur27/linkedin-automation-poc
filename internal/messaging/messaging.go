package messaging

import (
	"log"

	"github.com/go-rod/rod"
	"linkedin/internal/config"
)

func SendMessage(page *rod.Page, cfg config.Config, profileURL string) {
	log.Println("Preparing message for:", profileURL)

	if !cfg.AutomationEnabled {
		log.Println("[SAFE MODE] Would send message")
		return
	}

	log.Println("Automation enabled (not used in POC)")
}
