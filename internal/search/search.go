package search

import (
	"log"

	"github.com/go-rod/rod"
)

func SearchProfiles(page *rod.Page, keyword string) ([]string, error) {
	log.Println("Searching profiles for:", keyword)

	page.MustNavigate(
		"https://www.linkedin.com/search/results/people/?keywords=" + keyword,
	)

	// POC: returning dummy URLs
	results := []string{
		"https://www.linkedin.com/in/sample-profile-1",
		"https://www.linkedin.com/in/sample-profile-2",
	}

	return results, nil
}
