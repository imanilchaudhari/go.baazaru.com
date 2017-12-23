package controllers

import (

	//"log"
	"net/http"

	"baazaru.com/components/view"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

// Displays the list if avalable categories
func CategoryIndex(w http.ResponseWriter, r *http.Request) {
	var params httprouter.Params
	params = context.Get(r, "params").(httprouter.Params)
	slug := params.ByName("slug")

	v := view.New(r)
	v.Vars["slug"] = slug
	v.Name = "category_index"
	v.Render(w)
}

// Displays the list if avalable categories
func CategoryView(w http.ResponseWriter, r *http.Request) {
	var params httprouter.Params
	params = context.Get(r, "params").(httprouter.Params)
	slug := params.ByName("slug")

	v := view.New(r)
	v.Vars["slug"] = slug
	v.Name = "category_view"
	v.Render(w)
}
