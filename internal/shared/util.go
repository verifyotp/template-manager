package shared

import (
	"regexp"
	"strings"
	"time"
)

func GenerateSlug(name string) string {
	// Convert the name to lowercase
	slug := strings.ToLower(name)

	// Remove non-alphanumeric characters
	reg := regexp.MustCompile("[^a-z0-9]+")
	slug = reg.ReplaceAllString(slug, "-")

	// Remove leading and trailing hyphens
	slug = strings.Trim(slug, "-")

	return time.Now().Format("2006-01-02<>15:04:05.00") + "-" + slug
}
