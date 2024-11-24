# Daitem

API wrapper to connect and manager your personal Daitem Allarm system

## Configuration
You need to know:
- username
- password
- masterCode of your central

# How to user
```go
client := daitem.NewClient(&Options{
    Username:   "",
    Password:   "",
    MasterCode: "",
})

// List all user systems and find the systemId
systems, _ _= client.ListSystems()

// Get current state
state, _ := client.GetSystemState(systems[0].Id)
// state.State == 'off' || 'on'

// Activate or deactivate
client.SystemSendCommand(systems[0].Id, true)
```

For all commands see `client_test.go`