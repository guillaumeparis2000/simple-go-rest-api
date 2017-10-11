package router

import (
	"net/http"
	"github.com/guillaumeparis2000/rest-api/handlers"
)

type Route struct {
	Name 		string
	Method 		string
	Pattern 	string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		handlers.TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		handlers.TodoShow,
	},
	Route{
		"TodoDelete",
		"DELETE",
		"/todos/{todoId}",
		handlers.TodoDelete,
	},
	Route{
		"TodoUpdate",
		"PUT",
		"/todos/{todoId}",
		handlers.TodoUpdate,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		handlers.TodoCreate,
	},
}