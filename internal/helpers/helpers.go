package helpers

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/chingsley/bookings/internal/config"
)

var app *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func LogError(err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
}

func IsAuthenticated(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "user_id")
	return exists
}

func StrToDate(dateString string, w http.ResponseWriter) time.Time {
	layout := "2006-01-02"
	dateValue, err := time.Parse(layout, dateString)
	if err != nil {
		ServerError(w, errors.New("error while converting string to date"))
	}
	return dateValue
}

func DateToStr(dateValue time.Time) string {
	return dateValue.Format("2006-01-02")
}
