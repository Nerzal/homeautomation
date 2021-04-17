package login

import (
	"syscall/js"

	"github.com/Nerzal/tinydom"
	"github.com/Nerzal/tinydom/elements/input"
	"github.com/Nerzal/tinydom/elements/label"
)

var doc = tinydom.GetDocument()

func RenderLogin() {
	content := doc.CreateElement("content")
	content.RemoveAllChildNodes()

	form := doc.CreateElement("div")

	userNameLabel := label.New().
		SetFor("username").
		SetInnerHTML("Name:")

	userInput := input.New(input.TextInput).
		SetAutofocus(true).
		SetId("username").
		SetName("username")

	passwordLabel := label.New().
		SetFor("password").
		SetInnerHTML("Name:")

	passwordInput := input.New(input.PasswordInput).
		SetId("password").
		SetName("password")

	loginButton := input.New(input.ButtonInput).
		SetId("login-button").
		SetInnerHTML("Login").
		AddEventListener("click", js.FuncOf(OnLogin))

	form.AppendChildrenBr(userNameLabel, userInput, passwordLabel, passwordInput, loginButton)

	content.AppendChild(form)
}

func OnLogin(this js.Value, args []js.Value) interface{} {
	println("login button clicked")

	return nil
}
