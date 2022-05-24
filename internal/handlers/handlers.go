package handlers

import (
	"fmt"
	"github.com/raunchers/workAutomation/internal/render"
	"net/http"
)

type Update struct {
	Equipment    string // Equipment in alarm, I.E: RTR, ASTC, WAN etc
	Alarm        string // The alarm, I.E: NYCDOT-RTR-41BE - PING - CRITICAL: etc.
	HexID        string // Site's Hex ID
	TicketNumber string // The trouble ticket number
	DRMInfo      string // Output from DRM
}

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "home.page.tmpl")
}

func OfflineRTR(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "offlineRTR.page.tmpl")
}

func OfflineASTC(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "offlineASTC.page.tmpl")
}

func OfflineWAN2(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "offlineWAN2.page.tmpl")
}

func OfflineWAN3(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "offlineWAN3.page.tmpl")
}

func OfflineIPSEC(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "offlineIPSEC.page.tmpl")
}

func BouncingRTR(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "bouncingRTR.page.tmpl")
}

func BouncingASTC(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "bouncingASTC.page.tmpl")
}

func BouncingWAN2(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "bouncingWAN2.page.tmpl")
}

func BouncingWAN3(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "bouncingWAN3.page.tmpl")
}

func BouncingIPSEC(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "bouncingIPSEC.page.tmpl")
}

func PostOfflineRTR(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing ticket...")
	var info Update

	info.Equipment = r.FormValue("Equipment")
	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")

	fmt.Printf("%v\n", info.Equipment)
	fmt.Printf("%v\n", info.Alarm)
	fmt.Printf("%v\n", info.HexID)
	fmt.Printf("%v\n", info.TicketNumber)
}
