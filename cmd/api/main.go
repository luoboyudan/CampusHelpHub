package main

import "campushelphub/setup"

func main() {
	app := setup.InitializeApp()
	defer app.Stop()
	app.Run()
}
