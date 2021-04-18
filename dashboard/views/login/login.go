package login

import (
	"syscall/js"

	"github.com/Nerzal/tinydom"
	"github.com/Nerzal/tinydom/elements/input"
	"github.com/Nerzal/tinydom/elements/label"
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

	headerComponent := doc.CreateElement("div").
		SetClass("header")

	loginComponent := s.createLoginComponent()

	content.AppendChildren(headerComponent, loginComponent)

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

	loginComponent := doc.GetElementById("login-component")

	hintText := doc.CreateElement("p").
		SetInnerHTML("invalid username or password!").
		SetClass("invalid-input-text", "center").
		SetId("invalid-input-text")

	loginComponent.AppendChildBr(hintText)

}

func (s *Service) createLoginComponent() *tinydom.Element {
	loginComponent := doc.CreateElement("div").
		SetClass("login-component").
		SetId("login-component")

	userNameLabel := label.New().
		SetFor("username").
		SetInnerHTML("Name:").
		SetClass("center", "label")

	userInput := input.New(input.TextInput).
		SetAutofocus(true).
		SetId("username").
		SetName("username").
		SetClass("center")

	passwordLabel := label.New().
		SetFor("password").
		SetInnerHTML("Password:").
		SetClass("center", "label")

	passwordInput := input.New(input.PasswordInput).
		SetId("password").
		SetName("password").
		SetClass("center")

	loginButton := input.New(input.ButtonInput).
		SetValue("Sign In").
		AddEventListener("click", js.FuncOf(s.onLogin)).
		SetClass("center", "button")

	loginComponent.AppendChildrenBr(userNameLabel, userInput, passwordLabel, passwordInput, loginButton)

	s.userInput = userInput
	s.passwordInput = passwordInput

	return loginComponent
}
