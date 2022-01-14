package main

import "github.com/martini-contrib/cors"

const MAX_WORDS = 5
const MIN_WORDS_COUNT = 2
const WORDS_QUEUE_POLL_INTERVAL = 10

var CORS_OPTIONS = &cors.Options{
	AllowOrigins:     []string{"http://statika-shmatika.fvds.ru"},
	AllowMethods:     []string{"GET", "POST"},
	AllowHeaders:     []string{"X-Client-Id"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
}
