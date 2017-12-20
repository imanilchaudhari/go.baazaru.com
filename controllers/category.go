package controllers

import (

	//"log"
	"net/http"

	"baazaru.com/components/view"
)

// Displays the list if avalable categories
func CategoryListGet(w http.ResponseWriter, r *http.Request) {

	v := view.New(r)
	v.Name = "category"
	v.Render(w)
}
