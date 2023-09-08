package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandlerRoot)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/api", server.AddMiddleware(HandlerHome, CheckAuth(), Logging()))
	server.Listen()
}
