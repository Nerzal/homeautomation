package login

import (
	"syscall/js"

	"github.com/Nerzal/tinydom"
	"github.com/Nerzal/tinydom/elements/img"
	"github.com/Nerzal/tinydom/elements/input"
)

var doc = tinydom.GetDocument()

const userName = "Nerzal"
const password = "1234"

// Service holds user data
type Service struct {
	userInput     *tinydom.Element
	passwordInput *tinydom.Element
	events        chan Event
}

// New creates a new instance of Login
func New(events chan Event) *Service {
	return &Service{
		events: events,
	}
}

func (s *Service) RenderLogin() {
	content := doc.GetElementById("content")
	content.RemoveAllChildNodes()

	loginComponent := s.createLoginComponent()

	content.AppendChildren(loginComponent)

	doc.GetElementById("username").Focus()

}

func (s *Service) onLogin(this js.Value, args []js.Value) interface{} {
	userInput := input.FromElement(s.userInput).Value()
	passwordInput := input.FromElement(s.passwordInput).Value()

	println("user:", userInput)
	println("password:", passwordInput)

	// TODO: Let the Server(API) validate the credentials

	if userInput != userName {
		handleInvalidCredentials()
		return nil
	}

	if passwordInput != password {
		handleInvalidCredentials()
		return nil
	}

	// Goroutine is needed, as blocking operations like this are not allowed inside of async javascript handlers
	// This event handler is called from javascript -> gluecode -> Go
	go func() {
		s.events <- Event{UserName: userInput}
	}()

	return nil
}

func handleInvalidCredentials() {
	println("login failed: invalid credentials provided")

	userInput := doc.GetElementById("username")
	err := userInput.AppendClass("invalid-input")
	if err != nil {
		println("failed to append failed-input class:", err.Error())
	}

	passwordInput := doc.GetElementById("password")
	err = passwordInput.AppendClass("invalid-input")
	if err != nil {
		println("failed to append failed-input class:", err.Error())
	}

	invalidInputText := doc.GetElementById("invalid-input-text")
	if !invalidInputText.IsNull() {
		return
	}

	submitComponent := doc.GetElementById("submit-container")

	hintText := doc.CreateElement("p").
		SetInnerHTML("invalid username or password!").
		SetClass("invalid-input-text", "center").
		SetId("invalid-input-text")

	submitComponent.AppendChild(hintText)

}

func (s *Service) createLoginComponent() *tinydom.Element {
	loginComponent := doc.CreateElement("div").
		SetClass("login").
		SetId("login-component")

	header := doc.CreateElement("div").
		SetClass("header")

	title := doc.CreateElement("p").
		SetInnerHTML("TinyGo Wasm Homeautomation Dashboard Login")

	noobyGamesImg := img.New("assets/noobygames.png", "noobygames logo")
	tinyGoImg := img.New("assets/tinygo-logo-wasm.png", "tinygo wasm logo")

	header.AppendChildren(tinyGoImg.Element, title, noobyGamesImg.Element)

	userInput := input.New(input.TextInput).
		SetAutofocus(true).
		SetId("username").
		SetName("username").
		SetAttribute("placeholder", "Username")

	passwordInput := input.New(input.PasswordInput).
		SetId("password").
		SetName("password").
		SetAttribute("placeholder", "Password")

	submitContainer := doc.CreateElement("div").
		SetId("submit-container").
		SetClass("submit-container")

	loginButton := doc.
		CreateElement("button").
		SetAttribute("type", "button").
		SetInnerHTML("Sign In").
		AddEventListener("click", js.FuncOf(s.onLogin))

	submitContainer.AppendChild(loginButton)

	loginComponent.AppendChildren(header, userInput, passwordInput, submitContainer)

	s.userInput = userInput
	s.passwordInput = passwordInput

	return loginComponent
}
