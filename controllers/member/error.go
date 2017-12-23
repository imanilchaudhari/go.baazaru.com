package member

import (

	//"log"
	"net/http"

	"baazaru.com/components/view"
)

// Error401 handles 401 - Unauthorized
func Error401(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	//fmt.Fprint(w, "Unauthorized 401")

	// Display the view
	v := view.New(r)
	v.Name = "error_401"
	v.Render(w)
}

// Error404 handles 404 - Page Not Found
func Error404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	//fmt.Fprint(w, "Not Found 404")

	// Display the view
	v := view.New(r)
	v.Name = "error_404"
	v.Render(w)
}

// Error500 handles 500 - Internal Server Error
func Error500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	//fmt.Fprint(w, "Internal Server Error 500")

	// Display the view
	v := view.New(r)
	v.Name = "error_500"
	v.Render(w)
}
