package resttest

import (
	"io"
	"net/http"
	"net/http/httptest"
)

func request(HttpMethod string, routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	request, _ := http.NewRequest(HttpMethod, url, body)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(routeHandler)

	handler.ServeHTTP(responseRecorder, request)

	return *responseRecorder
}

// Process GET HTTP request with route handler and return response
func RequestGet(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodGet, routeHandler, url, body)
}

// Process POST HTTP request with route handler and return response
func RequestPost(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodPost, routeHandler, url, body)
}

// Process PATCH HTTP request with route handler and return response
func RequestPatch(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodPatch, routeHandler, url, body)
}

// Process PUT HTTP request with route handler and return response
func RequestPut(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodPut, routeHandler, url, body)
}

// Process DELETE HTTP request with route handler and return response
func RequestDelete(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodDelete, routeHandler, url, body)
}

// Process HEAD HTTP request with route handler and return response
func RequestHead(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodHead, routeHandler, url, body)
}

// Process OPTIONS HTTP request with route handler and return response
func RequestOptions(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodOptions, routeHandler, url, body)
}

// Process CONNECT HTTP request with route handler and return response
func RequestConnect(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodConnect, routeHandler, url, body)
}

// Process TRACE HTTP request with route handler and return response
func RequestTrace(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodTrace, routeHandler, url, body)
}
