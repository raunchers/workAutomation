package handlers

import (
	"github.com/raunchers/workAutomation/internal/render"
	"github.com/raunchers/workAutomation/pkg/models"
	"net/http"
)

// Home renders landing page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "home.page.tmpl")
}

// OfflineRTR Renders page for an offline router
func OfflineRTR(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "offlineRTR.page.tmpl")
}

// OfflineASTC render page for an offline ASTC
func OfflineASTC(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "offlineASTC.page.tmpl")
}

// OfflineWAN2 render page for an offline WAN 2
func OfflineWAN2(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "offlineWAN2.page.tmpl")
}

// OfflineWAN3 render page for an offline WAN 3
func OfflineWAN3(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "offlineWAN3.page.tmpl")
}

// OfflineIPSEC render page for an offline IPSEC tunnel
func OfflineIPSEC(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "offlineIPSEC.page.tmpl")
}

// BouncingRTR render page for a bouncing router
func BouncingRTR(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "bouncingRTR.page.tmpl")
}

// BouncingASTC render page for a bouncing ASTC
func BouncingASTC(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "bouncingASTC.page.tmpl")
}

// BouncingWAN2 render page for a bouncing WAN 2
func BouncingWAN2(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "bouncingWAN2.page.tmpl")
}

// BouncingWAN3 render page for a bouncing WAN 3
func BouncingWAN3(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "bouncingWAN3.page.tmpl")
}

// BouncingIPSEC render page for a bouncing IPSEC tunnel
func BouncingIPSEC(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "bouncingIPSEC.page.tmpl")
}

// RestoredRTRASTC render page for a router / ASTC that restored
func RestoredRTRASTC(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "restoredRTRASTC.page.tmpl")
}

// RestoredWANIPSEC render page for a WAN / IPSEC connection that restored
func RestoredWANIPSEC(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "restoredWANIPSEC.page.tmpl")
}

// Wobbly render page for a wobbly connection alarm
func Wobbly(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "wobbly.page.tmpl")
}

// CPUAlarm render page for a CPU alarm
func CPUAlarm(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "cpuAlarm.page.tmpl")
}

// TempAlarm render page for a Temperature alarm
func TempAlarm(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "tempAlarm.page.tmpl")
}

// CPUClear render page for a CPU ticket that cleared
func CPUClear(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "cpu.page.tmpl")
}

// TempClear render page for a temperature ticket that cleared
func TempClear(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "temp.page.tmpl")
}

func ATTTemplate(w http.ResponseWriter, r *http.Request) {
	render.RenTemplate(w, "at&tTemplate.page.tmpl")
}

// PostATTTemplate parse data from AT&T template page, render the AT&T ticket with the correct info
func PostATTTemplate(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.TicketNumber = r.FormValue("TicketNumber")
	info.ICCID = r.FormValue("ICCID")
	info.Address = r.FormValue("Address")
	info.Region = r.FormValue("Region")
	info.IssueStartDateTime = r.FormValue("IssueStartDateTime")
	info.Description = r.FormValue("Description")

	render.RenTicketTemplate(w, "at&tTicket.page.tmpl", info)
}

// PostTempAlarm parse data from temp alarm page, render temp alarm ticket template with correct info
func PostTempAlarm(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")
	info.DRMInfo = r.FormValue("DRMInfo")

	render.RenTicketTemplate(w, "tempHigh.page.tmpl", info)
}

// PostcpuClear parse data from cpu cleared page, render temp alarm ticket template with correct info
func PostcpuClear(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.DRMInfo = r.FormValue("DRMInfo")

	render.RenTicketTemplate(w, "cpuClear.page.tmpl", info)
}

// PosttempClear parse data from temp cleared page, render temp alarm ticket template with correct info
func PosttempClear(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.DRMInfo = r.FormValue("DRMInfo")

	render.RenTicketTemplate(w, "tempClear.page.tmpl", info)
}

// PostCPUAlarm parse data from temp alarm page, render temp alarm ticket template with correct info
func PostCPUAlarm(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")
	info.DRMInfo = r.FormValue("DRMInfo")

	render.RenTicketTemplate(w, "cpuHigh.page.tmpl", info)
}

// PostWobbly parse data from wobbly alarm page, render wobbly alarm ticket template with correct info
func PostWobbly(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.DRMInfo = r.FormValue("DRMInfo")

	render.RenTicketTemplate(w, "wobblyConnection.page.tmpl", info)
}

// PostOfflineRTR parse data from offline router alarm page, render offline router ticket template with correct info
func PostOfflineRTR(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")
	info.DRMInfo = r.FormValue("DRMInfo")

	render.RenTicketTemplate(w, "OffRTR.page.tmpl", info)
}

// PostOfflineASTC parse data from ASTC alarm page, render ASTC ticket template with correct info
func PostOfflineASTC(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")

	render.RenTicketTemplate(w, "OffASTC.page.tmpl", info)
}

// PostOfflineWAN2 parse data from WAN 2 alarm page, render WAN 2 ticket template with correct info
func PostOfflineWAN2(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")
	info.DRMInfo = r.FormValue("DRMInfo")

	render.RenTicketTemplate(w, "OffWAN2.page.tmpl", info)
}

// PostOfflineWAN3 parse data from WAN 3 alarm page, render WAN 3 ticket template with correct info
func PostOfflineWAN3(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")
	info.DRMInfo = r.FormValue("DRMInfo")

	render.RenTicketTemplate(w, "OffWAN3.page.tmpl", info)
}

// PostOfflineIPSEC parse data from IPSEC tunnel alarm page, render IPSEC tunnel ticket template with correct info
func PostOfflineIPSEC(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")
	info.DRMInfo = r.FormValue("DRMInfo")

	render.RenTicketTemplate(w, "OffIPSEC.page.tmpl", info)
}

// PostBouncingRTR parse data from bouncing RTR alarm page, render bouncing RTR ticket template with correct info
func PostBouncingRTR(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")

	render.RenTicketTemplate(w, "bounceRTR.page.tmpl", info)
}

// PostBouncingASTC parse data from bouncing ASTC alarm page, render bouncing ASTC ticket template with correct info
func PostBouncingASTC(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")

	render.RenTicketTemplate(w, "bounceASTC.page.tmpl", info)
}

// PostBouncingWAN2 parse data from bouncing WAN 2 alarm page, render bouncing WAN 2 ticket template with correct info
func PostBouncingWAN2(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")
	info.DRMInfo = r.FormValue("DRMInfo")

	render.RenTicketTemplate(w, "bounceWAN2.page.tmpl", info)
}

// PostBouncingWAN3 parse data from bouncing WAN 3 alarm page, render bouncing WAN 3 ticket template with correct info
func PostBouncingWAN3(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")
	info.DRMInfo = r.FormValue("DRMInfo")

	render.RenTicketTemplate(w, "bounceWAN3.page.tmpl", info)
}

// PostBouncingIPSEC parse data from bouncing IPSEC tunnel alarm page, render bouncing IPSEC tunnel ticket template with correct info
func PostBouncingIPSEC(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.Alarm = r.FormValue("Alarm")
	info.HexID = r.FormValue("HexID")
	info.TicketNumber = r.FormValue("TicketNumber")
	info.DRMInfo = r.FormValue("DRMInfo")

	render.RenTicketTemplate(w, "bounceIPSEC.page.tmpl", info)
}

// PostRestoredRTR parse data from RTR cleared page, render RTR ticket template with correct info
func PostRestoredRTR(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.DRMInfo = r.FormValue("DRMInfo")

	render.RenTicketTemplate(w, "restoreRTRASTC.page.tmpl", info)
}

// PostRestoredASTC parse data from ASTC cleared page, render ASTC ticket template with correct info
func PostRestoredASTC(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.DRMInfo = r.FormValue("DRMInfo")

	render.RenTicketTemplate(w, "restoreRTRASTC.page.tmpl", info)
}

// PostRestoredWANIPSEC parse data from WAN/IPSEC cleared page, render WAN/IPSEC ticket template with correct info
func PostRestoredWANIPSEC(w http.ResponseWriter, r *http.Request) {

	var info models.Update

	info.DRMInfo = r.FormValue("DRMInfo")

	render.RenTicketTemplate(w, "restoreWANIPSEC.page.tmpl", info)
}
