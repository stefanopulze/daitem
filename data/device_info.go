package data

type DeviceSession struct {
	SessionId string
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
}

type DeviceStatus struct {
	CommandStatus string
	SystemState   string
}
