package auth

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/go-rod/rod"
	"linkedin/internal/config"
)

func Login(page *rod.Page, cfg config.Config) error {
	email := os.Getenv("LINKEDIN_EMAIL")

	if !isTestEmail(email) {
		return errors.New("real LinkedIn accounts are blocked for safety")
	}

	log.Println("Login allowed for test account:", email)

	page.MustNavigate("https://www.linkedin.com/login")

	return nil
}

func isTestEmail(email string) bool {
	return strings.HasSuffix(email, "@test.com") ||
		strings.HasSuffix(email, "@dummy.com") ||
		strings.HasSuffix(email, "@example.com")
}
