package main

import "github.com/masilvasql/proxy/project"

func main() {
	users := make(map[string]bool)
	users["Lucas"] = true
	users["João"] = true
	users["Maria"] = false

	proxy := project.NewProxyVideoService(users)

	proxy.PlayVideo("Lucas")
	proxy.PlayVideo("João")
	proxy.PlayVideo("Maria")
}
