package main

import "snippetbox.justinlanghorst.com/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
