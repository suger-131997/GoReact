package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseRecorder struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (r *responseRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}
	n, err := r.ResponseWriter.Write(b)
	r.bytes += n
	return n, err
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()

		rr := &responseRecorder{ResponseWriter: w}
		next.ServeHTTP(rr, req)

		dur := time.Since(start)

		log.Printf(
			`http request method=%s path=%s status=%d bytes=%d remote=%s ua=%q dur=%s`,
			req.Method,
			req.URL.Path,
			rr.status,
			rr.bytes,
			req.RemoteAddr,
			req.UserAgent(),
			dur,
		)
	})
}
