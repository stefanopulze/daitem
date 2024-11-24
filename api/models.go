package api

const (
	CommandOK      = "CMD_OK"
	SystemStateOff = "off"
)

type System struct {
	Id     int    `json:"id"`
	Vendor string `json:"vendor"`
	Name   string `json:"name"`
	Role   int    `json:"role"`
}

type SystemState struct {
	Message       string  `json:"message"`
	State         string  `json:"systemState"`
	Groups        []Group `json:"groups"`
	Defaults      string  `json:"defaults"`
	CommandStatus string  `json:"commandStatus"`
}

type Group struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type CentralVersions struct {
	Box         string `json:"box"`
	BoxRadio    string `json:"boxRadio"`
	PlugKnx     string `json:"plugKnx"`
	RawVersions string `json:"rawVersions"`
}

type SystemConnect struct {
	TtmSessionId      string          `json:"ttmSessionId"`
	SystemState       string          `json:"systemState"`
	GroupList         []Group         `json:"groupList"`
	Status            string          `json:"status"`
	Versions          CentralVersions `json:"versions"`
	ConnectedUserType string          `json:"connectedUserType"`
}

type SystemInfo struct {
	Central      Central       `json:"central"`
	Transmitters []Transmitter `json:"transmitters"`
	Sensors      []Sensor      `json:"sensors"`
	Sirens       []Siren       `json:"sirens"`
	Commands     []Command     `json:"commands"`
	//ReadingDate  time.Time     `json:"readingDate"`
}

type Firmwares struct {
	Principal         Firmware `json:"principal"`
	Radio             Firmware `json:"radio"`
	TransmitterModule Firmware `json:"transmitterModule"`
	ComfortModule     Firmware `json:"comfortModule"`
}

func (f Firmwares) UpdatableFirmwares() bool {
	return f.Principal.UpdatableFirmware || f.Radio.UpdatableFirmware || f.TransmitterModule.UpdatableFirmware || f.ComfortModule.UpdatableFirmware
}

type Firmware struct {
	CurrentVersion    string `json:"currentVersion"`
	Type              string `json:"type"`
	UpdatableFirmware bool   `json:"updatableFirmware"`
}

type Central struct {
	PlugRTC           bool      `json:"plugRTC"`
	PlugGSM           bool      `json:"plugGSM"`
	PlugADSL          bool      `json:"plugADSL"`
	HasIO             bool      `json:"hasIO"`
	ParameterGsmSaved bool      `json:"parameterGsmSaved"`
	HasPlug           bool      `json:"hasPlug"`
	Firmwares         Firmwares `json:"firmwares"`
	Anomalies         struct {
		MainPowerSupplyAlert          bool `json:"mainPowerSupplyAlert"`
		SecondaryPowerSupplyAlert     bool `json:"secondaryPowerSupplyAlert"`
		DefaultMediaAlert             bool `json:"defaultMediaAlert"`
		AutoprotectionMechanicalAlert bool `json:"autoprotectionMechanicalAlert"`
		AutoprotectionWiredAlert      bool `json:"autoprotectionWiredAlert"`
	} `json:"anomalies"`
	CanInhibit bool `json:"canInhibit"`
}

func (c Central) HasAnomalies() bool {
	return c.Anomalies.MainPowerSupplyAlert ||
		c.Anomalies.SecondaryPowerSupplyAlert ||
		c.Anomalies.DefaultMediaAlert ||
		c.Anomalies.AutoprotectionMechanicalAlert ||
		c.Anomalies.AutoprotectionWiredAlert
}

type Transmitter struct {
	RefCode    string `json:"refCode"`
	Type       int    `json:"type"`
	Index      int    `json:"index"`
	Serial     string `json:"serial"`
	Label      string `json:"label"`
	Inhibited  bool   `json:"inhibited"`
	CanInhibit bool   `json:"canInhibit"`
	IsPlug     bool   `json:"isPlug"`
	//Firmwares  []Firmware `json:"firmwares"`
	Anomalies struct {
		MediaADSLAlert                bool `json:"mediaADSLAlert"`
		MediaGSMAlert                 bool `json:"mediaGSMAlert"`
		MediaRTCAlert                 bool `json:"mediaRTCAlert"`
		OutOfOrderAlert               bool `json:"outOfOrderAlert"`
		MainPowerSupplyAlert          bool `json:"mainPowerSupplyAlert"`
		SecondaryPowerSupplyAlert     bool `json:"secondaryPowerSupplyAlert"`
		AutoprotectionMechanicalAlert bool `json:"autoprotectionMechanicalAlert"`
		RadioAlert                    bool `json:"radioAlert"`
	} `json:"anomalies"`
	IsBox bool `json:"isBox"`
}

func (t Transmitter) HasAnomalies() bool {
	return t.Anomalies.MediaADSLAlert ||
		t.Anomalies.MediaGSMAlert ||
		t.Anomalies.MediaRTCAlert ||
		t.Anomalies.OutOfOrderAlert ||
		t.Anomalies.MainPowerSupplyAlert ||
		t.Anomalies.SecondaryPowerSupplyAlert ||
		t.Anomalies.AutoprotectionMechanicalAlert ||
		t.Anomalies.RadioAlert
}

type Sensor struct {
	Uid        string `json:"uid"`
	RefCode    string `json:"refCode"`
	Group      int    `json:"group"`
	Type       int    `json:"type"`
	Index      int    `json:"index"`
	Label      string `json:"label"`
	Serial     string `json:"serial"`
	Subtype    int    `json:"subtype"`
	Inhibited  bool   `json:"inhibited"`
	IsVideo    bool   `json:"isVideo"`
	CanInhibit bool   `json:"canInhibit"`
	Anomalies  struct {
		PowerSupplyAlert              bool `json:"powerSupplyAlert"`
		AutoprotectionMechanicalAlert bool `json:"autoprotectionMechanicalAlert"`
		SensorAlert                   bool `json:"sensorAlert"`
		LoopAlert                     bool `json:"loopAlert"`
		MaskAlert                     bool `json:"maskAlert"`
		RadioAlert                    bool `json:"radioAlert"`
	} `json:"anomalies"`
}

func (s Sensor) HasAnomalies() bool {
	return s.Anomalies.PowerSupplyAlert ||
		s.Anomalies.AutoprotectionMechanicalAlert ||
		s.Anomalies.SensorAlert ||
		s.Anomalies.LoopAlert ||
		s.Anomalies.MaskAlert ||
		s.Anomalies.RadioAlert
}

type Siren struct {
	RefCode    string `json:"refCode"`
	Type       int    `json:"type"`
	Index      int    `json:"index"`
	Serial     string `json:"serial"`
	Label      string `json:"label"`
	Inhibited  bool   `json:"inhibited"`
	CanInhibit bool   `json:"canInhibit"`
	Anomalies  struct {
		PowerSupplyAlert              bool `json:"powerSupplyAlert"`
		AutoprotectionMechanicalAlert bool `json:"autoprotectionMechanicalAlert"`
		RadioAlert                    bool `json:"radioAlert"`
	} `json:"anomalies"`
}

func (s Siren) HasAnomalies() bool {
	return s.Anomalies.PowerSupplyAlert ||
		s.Anomalies.AutoprotectionMechanicalAlert ||
		s.Anomalies.RadioAlert
}

type Command struct {
	RefCode    string `json:"refCode"`
	SubType    int    `json:"subType"`
	Type       int    `json:"type"`
	Index      int    `json:"index"`
	Serial     string `json:"serial"`
	Label      string `json:"label"`
	Inhibited  bool   `json:"inhibited"`
	CanInhibit bool   `json:"canInhibit"`
	Anomalies  struct {
		PowerSupplyAlert              bool `json:"powerSupplyAlert"`
		AutoprotectionMechanicalAlert bool `json:"autoprotectionMechanicalAlert"`
		RadioAlert                    bool `json:"radioAlert"`
	} `json:"anomalies"`
}

func (c Command) HasAnomalies() bool {
	return c.Anomalies.PowerSupplyAlert ||
		c.Anomalies.AutoprotectionMechanicalAlert ||
		c.Anomalies.RadioAlert
}

type Configuration struct {
	TransmitterId        string `json:"transmitterId"`
	CentralId            string `json:"centralId"`
	InstallationComplete bool   `json:"installationComplete"`
	Name                 string `json:"name"`
	Role                 int    `json:"role"`
	Id                   int    `json:"id"`
	Standalone           bool   `json:"standalone"`
	GprsPhone            string `json:"gprsPhone"`
}
