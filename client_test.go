package daitem

import (
	"log/slog"
	"testing"
)

func TestNewClient(t *testing.T) {
	options := &Options{
		Username:   "",
		Password:   "",
		MasterCode: "",
	}

	client := NewClient(options)

	systems, err := client.ListSystems()
	if err != nil {
		t.Fatal(err)
	}
	current := systems[0]

	info, err := client.GetSystemInfo(current.Id)
	slog.Info("System Info", slog.Any("info", info))

	configuration, err := client.GetSystemConfiguration(current.Id)
	slog.Info("System configuration", slog.Any("configuration", configuration))

	connect, err := client.SystemConnect(current.Id)
	slog.Info("System connection", slog.Any("connection", connect))

	state, err := client.GetSystemState(current.Id)
	slog.Info("System state", slog.Any("state", state))

	//_, err = client.SystemSendCommand(current.Id, false)
	//if err != nil {
	//	t.Fatal(err)
	//}

	transmitterConnected, err := client.IsTransmitterConnect(configuration.TransmitterId)
	slog.Info("Transmitter connected", slog.Any("connected", transmitterConnected))
}
