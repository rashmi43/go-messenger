package middleware

import (
	"log"
        "fmt"
	"strings"
	"net/http"
)

func MiddlewareFirst(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("MiddlewareFirst - Before Handler")
                var request []string
                // Loop through headers
                // Add the request string
 		url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
 		request = append(request, url)
		 // Add the host
 		request = append(request, fmt.Sprintf("Host: %v", r.Host))
 		for name, headers := range r.Header {
   			name = strings.ToLower(name)
   			for _, h := range headers {
     				request = append(request, fmt.Sprintf("%v: %v", name, h))
   			}
 		}
		// Tracability , monitor where request are coming from
                fmt.Println(request)
		next.ServeHTTP(w, r)
		log.Println("MiddlewareFirst - After Handler")
	})
}
