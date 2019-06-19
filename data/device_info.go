package data

import "time"

type DeviceSession struct {
	SessionId   string
	SessionTime time.Time
	UseCache    bool
}

type DeviceInfo struct {
	ClientIpAddress string
	FirmwareVersion string
	GprsConnection  string
	Groups          []int
	SystemIpAddress string
	SystemState     string
	TTMSessionId    string
}

type DeviceConfiguration struct {
	TransmitterId  string
	CentralId      string
	ConnectionType string
	UseCache       bool
}

type DeviceStatus struct {
	CommandStatus string
	SystemState   string
	UseCache      bool
}
