// cui: http request/response tui
// Copyright 2022 Mario Finelli
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"net/http"
)

var requestKinds = []string{"Form Data", "JSON", "Raw"}

var httpMethods = []string{
	http.MethodDelete,
	http.MethodHead,
	http.MethodGet,
	http.MethodOptions,
	http.MethodPatch,
	http.MethodPost,
	http.MethodPut,
}

// a subset from: https://en.wikipedia.org/wiki/List_of_HTTP_header_fields
var commonHeaderKeys = []string{
	"Accept",
	"Accept-Charset",
	"Accept-Encoding",
	"Accept-Language",
	"Access-Control-Request-Method",
	"Access-Control-Request-Headers",
	"Authorization",
	"Cache-Control",
	"Connection",
	"Content-Encoding",
	"Content-Length",
	"Content-MD5",
	"Content-Type",
	"Cookie",
	"Date",
	"Expect",
	"Forwarded",
	"From",
	"Host",
	"If-Match",
	"If-Modified-Since",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
	"Max-Forwards",
	"Origin",
	"Pragma",
	"Prefer",
	"Proxy-Authorization",
	"Range",
	"Referer",
	"Transfer-Encoding",
	"User-Agent",
	"Via",
	"X-Requested-With",
	"X-Forwarded-For",
	"X-Forwarded-Host",
	"X-Forwarded-Proto",
	"X-Csrf-Token",
	"X-Request-ID",
}

var commonHeaderContentTypes = []string{
	"application/javascript",
	"application/json",
	"application/x-www-form-urlencoded",
	"application/xml",
	"text/css",
	"text/csv",
	"text/html",
	"text/plain",
	"text/xml",
}
