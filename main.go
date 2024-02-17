package main

func main() {
	server := &APIServer{":3000"}
	server.Run()
}
