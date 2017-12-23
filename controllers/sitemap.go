package controllers

import (
	"net/http"

	"baazaru.com/components/view"
)

// Displays the sitemap page
func SitemapIndex(w http.ResponseWriter, r *http.Request) {
	// Display the view
	// w.Header().Set("Content-Type", "application/xml")
	v := view.New(r)
	v.Name = "sitemap_index"
	v.Render(w)
}
