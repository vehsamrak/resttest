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

// RequestGet emulates GET HTTP request with route handler and return response
func RequestGet(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodGet, routeHandler, url, body)
}

// RequestPost emulates POST HTTP request with route handler and return response
func RequestPost(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodPost, routeHandler, url, body)
}

// RequestPatch emulates PATCH HTTP request with route handler and return response
func RequestPatch(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodPatch, routeHandler, url, body)
}

// RequestPut emulates PUT HTTP request with route handler and return response
func RequestPut(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodPut, routeHandler, url, body)
}

// RequestDelete emulates DELETE HTTP request with route handler and return response
func RequestDelete(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodDelete, routeHandler, url, body)
}

// RequestHead emulates HEAD HTTP request with route handler and return response
func RequestHead(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodHead, routeHandler, url, body)
}

// RequestOptions emulates OPTIONS HTTP request with route handler and return response
func RequestOptions(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodOptions, routeHandler, url, body)
}

// RequestConnect emulates CONNECT HTTP request with route handler and return response
func RequestConnect(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodConnect, routeHandler, url, body)
}

// RequestTrace emulates TRACE HTTP request with route handler and return response
func RequestTrace(routeHandler http.HandlerFunc, url string, body io.Reader) httptest.ResponseRecorder {
	return request(http.MethodTrace, routeHandler, url, body)
}
