package handlers

import (
	"fmt"
	"github.com/raunchers/workAutomation/internal/render"
	"github.com/raunchers/workAutomation/pkg/models"
	"net/http"
)

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

	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")

	render.RenTicketTemplate(w, "OffRTR.page.tmpl", info)
}

func PostOfflineASTC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing ASTC ticket...")
	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")

	fmt.Printf("%v\n", info.Alarm)
	fmt.Printf("%v\n", info.HexID)
	fmt.Printf("%v\n", info.TicketNumber)

	render.RenTicketTemplate(w, "OffASTC.page.tmpl", info)
}

func PostOfflineWAN2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing WAN 2 ticket...")
	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")

	fmt.Printf("%v\n", info.Alarm)
	fmt.Printf("%v\n", info.HexID)
	fmt.Printf("%v\n", info.TicketNumber)
}

func PostOfflineWAN3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing WAN 3 ticket...")
	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")

	fmt.Printf("%v\n", info.Alarm)
	fmt.Printf("%v\n", info.HexID)
	fmt.Printf("%v\n", info.TicketNumber)
}

func PostOfflineIPSEC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing IPSEC ticket...")
	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")

	fmt.Printf("%v\n", info.Alarm)
	fmt.Printf("%v\n", info.HexID)
	fmt.Printf("%v\n", info.TicketNumber)
}

func PostBouncingRTR(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")

	fmt.Printf("%v\n", info.Alarm)
	fmt.Printf("%v\n", info.HexID)
	fmt.Printf("%v\n", info.TicketNumber)

	render.RenTicketTemplate(w, "bounceRTR.page.tmpl", info)
}

func PostBouncingASTC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing Bouncing ASTC ticket...")
	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")

	fmt.Printf("%v\n", info.Alarm)
	fmt.Printf("%v\n", info.HexID)
	fmt.Printf("%v\n", info.TicketNumber)

	render.RenTicketTemplate(w, "bounceASTC.page.tmpl", info)
}

func PostBouncingWAN2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing Bouncing WAN 2 ticket...")
	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")

	fmt.Printf("%v\n", info.Alarm)
	fmt.Printf("%v\n", info.HexID)
	fmt.Printf("%v\n", info.TicketNumber)
}

func PostBouncingWAN3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing Bouncing WAN 3 ticket...")
	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")

	fmt.Printf("%v\n", info.Alarm)
	fmt.Printf("%v\n", info.HexID)
	fmt.Printf("%v\n", info.TicketNumber)
}

func PostBouncingIPSEC(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing Bouncing IPSEC ticket...")
	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")

	fmt.Printf("%v\n", info.Alarm)
	fmt.Printf("%v\n", info.HexID)
	fmt.Printf("%v\n", info.TicketNumber)
}
