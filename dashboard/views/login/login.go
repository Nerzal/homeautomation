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
		SetInnerHTML("Password:")

	passwordInput := input.New(input.PasswordInput).
		SetId("password").
		SetName("password")

	loginButton := input.New(input.ButtonInput).
		SetValue("Login").
		AddEventListener("click", js.FuncOf(s.onLogin))

	form.AppendChildrenBr(userNameLabel, userInput, passwordLabel, passwordInput, loginButton)

	content.AppendChild(form)

	s.userInput = userInput
	s.passwordInput = passwordInput
}

func (s *Service) onLogin(this js.Value, args []js.Value) interface{} {
	userInput := input.FromElement(s.userInput).Value()
	passwordInput := input.FromElement(s.passwordInput).Value()

	println("user:", userInput)
	println("password:", passwordInput)

	if userInput != userName {
		handleInvalidCredentials()
		return nil
	}

	if passwordInput != password {
		handleInvalidCredentials()
		return nil
	}

	go func() {
		s.events <- Event{UserName: userInput}
	}()

	return nil
}

func handleInvalidCredentials() {
	tinydom.GetWindow().Alert("Invalid username or password")
}
