not-a-web-framework
===================

A baseline for creating web-services in Go (a bit on steroids, but still pretty small). It's not a framework, not even a library. Steal this code instead of making yet another framework out of it! Don't even `go get` this package. Instead - git clone and merge it with your code base or build your next service on top of it.

The inspiration comes from the blog post [Build Your Own Framework in Go](http://nicolasmerouze.com/build-web-framework-golang/).

## Which packages are used?

 * Alice (`github.com/justinas/alice`) for middleware (http.Handler chaining)
 * Gorilla Context (`github.com/gorilla/context`) for sharing data between middlewares
 * HttpRouter (`github.com/julienschmidt/httprouter`) for pattern-based request routing
