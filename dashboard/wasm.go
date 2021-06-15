package main

import (
	"github.com/Nerzal/homeautomation/dashboard/views/dashboard"
	"github.com/Nerzal/homeautomation/dashboard/views/login"
)

func main() {
	dashboardService := dashboard.New()

	loginEvents := make(chan login.Event, 1)
	loginService := login.New(loginEvents)
	loginService.RenderLogin()

	loginEvent := <-loginEvents

	println("New user logged in:", loginEvent.UserName)

	dashboardService.RenderDashboard()
	select {}
}
