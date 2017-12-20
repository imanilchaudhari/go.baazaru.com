package logrequest

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"baazaru.com/components/logger"
	"baazaru.com/components/session"
	"baazaru.com/models"
)

// ConsoleTarget will log the HTTP requests
func ConsoleTarget(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// DbTarget will log the HTTP requests
func DbTarget(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !strings.Contains(r.URL.Path, "/api/v1/verify") {
			// Get session
			sess := session.Instance(r)

			user_id := uint64(0)

			// If the user is logged in
			if sess.Values["id"] != nil {
				user_id = uint64(sess.Values["id"].(uint32))
			}

			err := models.TrackRequestURL(user_id, r)
			if err != nil {
				log.Println(err)
			}
		}

		next.ServeHTTP(w, r)
	})
}

// FileTarget will log the HTTP requests
func FileTarget(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set directory path to save log files
		logger.Path("./runtime/log")

		// Set log level by constant
		logger.Level(logger.DEBUG)

		// Set log level by string
		logger.LevelAsString("debug")

		// Set function to prepare format log line
		logger.Format(func(level int, line string, message string) string {
			return fmt.Sprintf("Only message: %s", message)
		})

		// Set log file size limit
		logger.SizeLimit(2 * logger.MB)

		ip, err := models.GetRemoteIP(r)
		if err != nil {
			log.Println(err)
		}

		logger.Debug("REQUEST IP :- " + ip + " URI " + r.URL.RequestURI())

		// Set log output to stdout
		logger.Stdout(true)

		next.ServeHTTP(w, r)
	})
}
