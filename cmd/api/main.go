// main.go
package main

import "gin_app/routes"

func main() {
	router := routes.AppRoutes()
	router.Run()
}
