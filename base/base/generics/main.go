package main

func main() {
	//mock_http_server.StartHttpServer()
	Ptr(2)
	Ptr(true)
}

func Ptr[T any](v T) *T {
	return &v
}
