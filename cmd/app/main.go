package main

import (
	"log"

	"linkedin/internal/auth"
	"linkedin/internal/browser"
	"linkedin/internal/config"
	"linkedin/internal/search"
)

func main() {
	log.Println("Starting LinkedIn Automation POC")

	cfg := config.Load()

	b := browser.Init()
	page := b.MustPage("https://www.linkedin.com")

	err := auth.Login(page, cfg)
	if err != nil {
		log.Println("Login blocked:", err)
		return
	}

	profiles, _ := search.SearchProfiles(page, "software engineer")
	log.Println("Profiles found:", len(profiles))

	log.Println("POC run completed safely ")
}
