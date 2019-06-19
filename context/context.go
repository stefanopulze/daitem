package context

import (
	"github.com/stefanopulze/daitem/storage"
	"time"
)

var contextExpirationOffset *time.Duration

// Context contains all sensitive data of the user and the current http session used by
// api to make http call
type Context struct {
	Username       string
	Password       string
	MasterCode     string
	CentralId      string
	TransmitterId  string
	SessionId      string
	TTMSessionId   string
	ConnectionType string
	creationTime   time.Time
}

// Check if context is expired
// Default context are valid for 15 minutes after [creationTime]
// @return true if expired
func (ctx *Context) IsExpired() bool {
	if ctx.SessionId == "" {
		return true
	}

	if contextExpirationOffset == nil {
		d, _ := time.ParseDuration("15m")
		contextExpirationOffset = &d
	}

	return ctx.creationTime.Add(*contextExpirationOffset).After(time.Now())
}

// Check if context is NOT expired
func (ctx *Context) IsValid() bool {
	return !ctx.IsExpired()
}

// Persist context on storage
func (ctx *Context) Save(storage storage.Storage) {

}

// Load context from storage
func Load(storage storage.Storage) (*Context, error) {
	return nil, nil
}
