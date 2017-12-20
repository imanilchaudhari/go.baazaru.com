package controllers

import (
	"net/http"

	"baazaru.com/components/view"
)

// Displays the about page
func AboutGET(w http.ResponseWriter, r *http.Request) {
	// Display the view
	v := view.New(r)
	v.Name = "about"
	v.Render(w)
}
