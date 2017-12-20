package controllers

import (

	//"log"
	"net/http"

	"baazaru.com/components/view"
)

// Displays the list if avalable categories
func GetItems(w http.ResponseWriter, r *http.Request) {

	v := view.New(r)
	v.Name = "items"
	v.Render(w)
}
