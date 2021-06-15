package dashboard

import (
	"syscall/js"

	"github.com/Nerzal/tinydom"
	"github.com/Nerzal/tinydom/elements/a"
	"github.com/Nerzal/tinydom/elements/li"
	"github.com/Nerzal/tinydom/elements/nav"
)

var doc = tinydom.GetDocument()

// Service provides functionality to render the dashboard
type Service struct {
}

// New creates a new instance of Service
func New() *Service {
	js.Global().Set("handleMessage", js.FuncOf(handleMessage))

	return &Service{}
}

func (s *Service) RenderDashboard() {
	js.Global().
		Get("ConnectToMQTT").
		Invoke()

	content := doc.GetElementById("content")
	content.RemoveAllChildNodes()

	content.
		Style().
		SetHeight("100%")

	dashboardLink := a.New("#", "Dashboard")
	livingRoomLink := a.New("#", "Livingroom")
	bedroomLink := a.New("#", "Bedroom")

	homeItem := li.New()
	homeItem.AppendChild(dashboardLink.Element)

	livingRoomItem := li.New()
	livingRoomItem.AppendChild(livingRoomLink.Element)

	bedroomItem := li.New()
	bedroomItem.AppendChild(bedroomLink.Element)

	navigation := nav.New().
		AppendListItem(homeItem).
		AppendListItem(livingRoomItem).
		AppendListItem(bedroomItem).
		SetClass("sidebar")

	overview := doc.CreateElement("div").
		SetId("overview").
		SetClass("overview")

	overviewHeader := doc.CreateElement("h1").
		SetInnerHTML("Overview")

	turnOnButton := doc.
		CreateElement("button").
		SetAttribute("type", "button").
		SetInnerHTML("On").
		AddEventListener("click", js.FuncOf(s.turnOn))

	turnOffButton := doc.
		CreateElement("button").
		SetAttribute("type", "button").
		SetInnerHTML("Off").
		AddEventListener("click", js.FuncOf(s.turnOff))

	buttonDescriptor := doc.CreateElement("h2").SetInnerHTML("Bedroom Lights")

	overview.AppendChildren(overviewHeader, doc.CreateElement("br"), buttonDescriptor, turnOnButton, turnOffButton)

	content.AppendChildren(navigation, overview)
}

func (s *Service) turnOn(this js.Value, args []js.Value) interface{} {
	println("bedroom: turnOn button pressed")

	js.Global().Get("publish").Invoke("/noobygames/homeautomation/home/bedroom/light/on", "on", 2)

	return nil
}

func (s *Service) turnOff(this js.Value, args []js.Value) interface{} {
	println("bedroom: turnOff button pressed")

	js.Global().Get("publish").Invoke("/noobygames/homeautomation/home/bedroom/light/off", "off", 2)

	return nil
}

func (s *Service) onLogin(this js.Value, args []js.Value) interface{} {
	return nil
}

func handleMessage(this js.Value, args []js.Value) interface{} {
	message := args[0].String()
	println("mqtt message arrived:", message)

	return nil
}
