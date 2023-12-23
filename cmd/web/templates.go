package main

import "github.com/waeekron/snippetbox/internal/models"

// Holding structure for any dynamic data passed down to a template
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
