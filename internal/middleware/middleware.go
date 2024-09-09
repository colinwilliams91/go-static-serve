package middleware

import "net/http"

// CacheControlMiddleware ensures content is not served from browser's cache
func CacheControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// *** DISABLE CACHING *** FOR DEVELOPMENT ONLY? ***
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		next.ServeHTTP(w, r)
	})
}

/* -- Cache-Control --

  no-cache: Instructs the browser that it must revalidate with the server before using a cached copy.
	  		The browser will make a request to the server to check if the content has changed.

  no-store: Directs the browser to avoid storing any part of the request or response in its cache.
  			This means no information is saved between requests.

  must-revalidate: Ensures that if the content becomes stale, the browser must revalidate it with the server before using it.
  			It reinforces the requirement for revalidation.
*/

/* -- Pragma --

  no-cache: This is a legacy HTTP/1.0 header used to prevent caching. While Cache-Control is preferred
  		  in HTTP/1.1 and later, Pragma: no-cache is still included for compatibility with older HTTP/1.0 browsers.
*/

/* -- Expires --

  0: Sets the expiration date of the content to a time in the past, effectively indicating that the content
  		   is already stale and should not be cached. It’s another way to ensure that the browser fetches fresh content.
*/