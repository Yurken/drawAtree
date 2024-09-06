package main

import "drawAtree/web"

func main() {
	listenTo := "127.0.0.1:12345"
	web.WebServer(listenTo)
}
