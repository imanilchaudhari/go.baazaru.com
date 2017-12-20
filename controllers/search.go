package controllers

import (

	//"log"
	"net/http"

	"baazaru.com/components/view"
)

// Displays the list if avalable categories
func SearchListGet(w http.ResponseWriter, r *http.Request) {

	v := view.New(r)
	v.Name = "search"
	v.Render(w)
}
