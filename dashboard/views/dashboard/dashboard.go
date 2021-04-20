package dashboard

import (
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
	return &Service{}
}

func (s *Service) RenderDashboard() {
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

	content.AppendChildren(navigation, overview)
}
