package route

import (
	"net/http"

	"baazaru.com/controllers"
	"baazaru.com/route/middleware/acl"
	hr "baazaru.com/route/middleware/httprouterwrapper"
	"baazaru.com/route/middleware/logrequest"
	//"baazaru.com/route/middleware/pprofhandler"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// Load the HTTP routes and middleware
func LoadHTTPS() http.Handler {
	return middleware(routes())
}

// Load the HTTPS routes and middleware
func LoadHTTP() http.Handler {
	return middleware(routes())

	// Uncomment this and comment out the line above to always redirect to HTTPS
	//return http.HandlerFunc(redirectToHTTPS)
}

func redirectToHTTPS(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host, http.StatusMovedPermanently)
}

// *****************************************************************************
// Routes
// *****************************************************************************

func routes() *httprouter.Router {
	r := httprouter.New()

	// Set 404 handler
	r.NotFound = alice.
		New().
		ThenFunc(controllers.Error404)

	// Serve static files, no directory browsing
	r.GET("/web/*filepath", hr.Handler(alice.
		New().
		ThenFunc(controllers.Static)))

	// Home page and public pages
	r.GET("/", hr.Handler(alice.
		New().
		ThenFunc(controllers.Index)))
	r.GET("/about", hr.Handler(alice.
		New().
		ThenFunc(controllers.AboutGET)))
	r.GET("/terms", hr.Handler(alice.
		New().
		ThenFunc(controllers.TermsGET)))
	r.GET("/contact", hr.Handler(alice.
		New().
		ThenFunc(controllers.ContactGET)))
	r.POST("/contact", hr.Handler(alice.
		New().
		ThenFunc(controllers.ContactPOST)))
	r.GET("/verify", hr.Handler(alice.
		New().
		ThenFunc(controllers.VerifyUsernameGET)))
	r.GET("/public/:site/:username", hr.Handler(alice.
		New().
		ThenFunc(controllers.PublicUsernameGET)))
	r.GET("/image/:userid/:picid", hr.Handler(alice.
		New().
		ThenFunc(controllers.WatermarkImagesGET)))

	// Login and logout
	r.GET("/login", hr.Handler(alice.
		New(acl.DisallowAuth).
		ThenFunc(controllers.LoginGET)))
	r.POST("/login", hr.Handler(alice.
		New(acl.DisallowAuth).
		ThenFunc(controllers.LoginPOST)))
	r.GET("/logout", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.Logout)))

	// Register
	r.GET("/register", hr.Handler(alice.
		New(acl.DisallowAuth).
		ThenFunc(controllers.RegisterGET)))
	r.POST("/register", hr.Handler(alice.
		New(acl.DisallowAuth).
		ThenFunc(controllers.RegisterPOST)))

	// Email verification
	r.GET("/emailverification/:token", hr.Handler(alice.
		New().
		ThenFunc(controllers.EmailVerificationGET)))

	// User Pages
	r.GET("/profile", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.UserProfileGET)))
	r.GET("/profile/initial", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.InitialPhotoGET)))
	r.POST("/profile/initial", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.PhotoPOST)))
	r.GET("/profile/initial/delete/:picid", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.InitialPhotoDeleteGET)))
	r.GET("/profile/photo/upload", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.PhotoUploadGET)))
	r.POST("/profile/photo/upload", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.PhotoPOST)))
	r.GET("/profile/photo/delete/:picid", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.PhotoDeleteGET)))
	r.GET("/profile/site", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.UserSiteGET)))
	r.POST("/profile/site", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.UserSitePOST)))
	r.GET("/profile/photo/download/:picid", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.PhotoDownloadGET)))
	r.POST("/profile/information", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.UserInformationPOST)))
	r.GET("/profile/information", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.UserInformationGET)))
	r.GET("/profile/email", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.UserEmailGET)))
	r.POST("/profile/email", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.UserEmailPOST)))
	r.GET("/profile/password", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.UserPasswordGET)))
	r.POST("/profile/password", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(controllers.UserPasswordPOST)))

	// Admin Pages
	r.GET("/admin", hr.Handler(alice.
		New(acl.AllowOnlyAdministrator).
		ThenFunc(controllers.AdminGET)))
	r.GET("/admin/user/:userid", hr.Handler(alice.
		New(acl.AllowOnlyAdministrator).
		ThenFunc(controllers.AdminUserGET)))
	r.GET("/admin/all", hr.Handler(alice.
		New(acl.AllowOnlyAdministrator).
		ThenFunc(controllers.AdminAllGET)))
	r.GET("/admin/approve/:userid/:picid", hr.Handler(alice.
		New(acl.AllowOnlyAdministrator).
		ThenFunc(controllers.AdminApproveGET)))
	r.GET("/admin/reject/:userid/:picid", hr.Handler(alice.
		New(acl.AllowOnlyAdministrator).
		ThenFunc(controllers.AdminRejectGET)))
	r.GET("/admin/unverify/:userid/:picid", hr.Handler(alice.
		New(acl.AllowOnlyAdministrator).
		ThenFunc(controllers.AdminUnverifyGET)))

	// API
	r.GET("/api/v1/verify/:site/:username", hr.Handler(alice.
		New().
		ThenFunc(controllers.APIVerifyUserGET)))
	r.GET("/api/v1/request/token", hr.Handler(alice.
		New().
		ThenFunc(controllers.APIRegisterChromeGET)))

	// Cron
	// TODO This should not be publicly accessible
	r.GET("/cron/notifyemailexpire", hr.Handler(alice.
		New().
		ThenFunc(controllers.CronNotifyEmailExpireGET)))

	// Enable Pprof
	/*r.GET("/debug/pprof/*pprof", hr.Handler(alice.
	New(acl.DisallowAnon).
	ThenFunc(pprofhandler.Handler)))*/

	return r
}

// *****************************************************************************
// Middleware
// *****************************************************************************

func middleware(h http.Handler) http.Handler {
	// Log every request
	h = logrequest.Database(h)

	// Clear handler for Gorilla Context
	h = context.ClearHandler(h)

	return h
}
