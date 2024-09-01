package middlweare

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mnsdojo/custom-router-go/internal/router"
)

const (
	Reset  = "\033[0m"
	Green  = "\033[32m"
	Cyan   = "\033[36m"
	Yellow = "\033[33m"
)

func LoggerMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		method := Green + r.Method + Reset
		path := Cyan + r.URL.Path + Reset
		fmt.Printf("Started %s %s\n", method, path)
		next(w, r)
		duration := time.Since(start)
		fmt.Printf("Completed %s %s in %s\n", method, path, Yellow+duration.String()+Reset)
	}
}
