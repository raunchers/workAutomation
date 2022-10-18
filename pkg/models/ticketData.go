package models

type Update struct {
	Equipment          string // Equipment in alarm, I.E: RTR, ASTC, WAN etc
	Alarm              string // The alarm, I.E: NYCDOT-RTR-41BE - PING - CRITICAL: etc.
	HexID              string // Site's Hex ID
	TicketNumber       string // The trouble ticket number
	DRMInfo            string // Output from DRM
	ICCID              string // ICCID of down Cell
	IMSI               string // IMSI of down cell for T-Mobile tickets
	SignalStrength     string // Signal Strength of T-Mobile SIM
	SignalQuality      string // Signal Quality of T-Mobile SIM
	Latency            string // Latency of down / bouncing T-Mobile connection
	ZipCode            string // ZipCode where the kit is located
	Address            string // Cross street where the kit is located
	Region             string // Borough the kit is located in, I:E Brox, Staten Island, etc ...
	IssueStartDateTime string // Date and time of the start of the issue.
	LastConnect        string // last time the device successfully connected
	Description        string // Extra details for the AT&T and T-Mobile tickets
}
