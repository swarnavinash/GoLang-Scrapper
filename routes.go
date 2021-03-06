package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		IndexHandler,
	},
	Route{
		"Amazon",
		"GET",
		"/movie/amazon/",
		AmazonHandler,
	},
	Route{
		"AmazonScrape",
		"GET",
		"/movie/amazon/{amazon_id}",
		AmazonScrappingHandler,
	},
}
