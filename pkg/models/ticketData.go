package models

type Update struct {
	Equipment    string // Equipment in alarm, I.E: RTR, ASTC, WAN etc
	Alarm        string // The alarm, I.E: NYCDOT-RTR-41BE - PING - CRITICAL: etc.
	HexID        string // Site's Hex ID
	TicketNumber string // The trouble ticket number
	DRMInfo      string // Output from DRM
}
