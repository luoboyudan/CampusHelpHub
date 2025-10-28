package main

import "campushelphub/setup"

func main() {
	app := setup.InitializeApp()
	app.Engine.Run(app.Config.Server.Addr)
}
