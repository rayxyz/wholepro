package main

// The request and response types should be annotated sufficiently for all
// transports we intend to use.

type request struct {
	A int64 `json:"a"`
	B int64 `json:"b"`
}

type response struct {
	V int64 `json:"v"`
}
