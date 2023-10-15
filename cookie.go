
// setting a cookie

cookie := &http.Cookie{
	Name:     "secret cookie",
	Value:    "age: 23",
	Path:     "/",
	MaxAge:   3600, // Cookie will expire after 1 hour (in seconds)
	HttpOnly: true, // Cookie cannot be accessed via JavaScript
	Secure:   true, // Cookie can only be sent over HTTPS
}

// Set the cookie in the response
http.SetCookie(w, cookie)

// keep your cookies as small as possible
// don't store sensitive data in cookies
// use http-only cookies
// set expiration dates on your cookies
// set the path on your cookies
// set the domain on your cookies
// set the secure flag on your cookies
// set the same-site flag on your cookies
		

