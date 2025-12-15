package main

func main() {
	//mock_http_server.StartHttpServer()
	Ptr(2)
	Ptr("123ABC")
	Ptr(true)
	Ptr(12.3)
}

func Ptr[T any](v T) *T {
	return &v
}
