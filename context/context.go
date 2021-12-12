package context

import (
	"github.com/stefanopulze/daitem/storage"
	"log"
	"time"
)

var contextExpirationOffset *time.Duration

const (
	SessionId      = "SessionId"
	SessionTime    = "SessionTime"
	CentralId      = "CentralId"
	TransmitterId  = "TransmitterId"
	ConnectionType = "ConnectionType"
	TTMSessionId   = "TTMSessionId"
)

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
	SessionTime    time.Time
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

	return ctx.SessionTime.Add(*contextExpirationOffset).Before(time.Now())
}

func (ctx *Context) Merge(source *Context) {
	if len(ctx.CentralId) == 0 {
		ctx.CentralId = source.CentralId
	}

	if len(ctx.TransmitterId) == 0 {
		ctx.TransmitterId = source.TransmitterId
	}

	ctx.SessionId = source.SessionId
	ctx.TTMSessionId = source.TTMSessionId
	ctx.ConnectionType = source.ConnectionType
	ctx.SessionTime = source.SessionTime
}

// IsValid Check if context is NOT expired
func (ctx *Context) IsValid() bool {
	return !ctx.IsExpired()
}

// Load context from storage
func Load(storage storage.Storage) (*Context, error) {
	ctx := Context{}

	if bytes, e := storage.Read(CentralId); e == nil {
		ctx.CentralId = string(bytes)
	}
	if bytes, e := storage.Read(TransmitterId); e == nil {
		ctx.TransmitterId = string(bytes)
	}
	if bytes, e := storage.Read(SessionId); e == nil {
		ctx.SessionId = string(bytes)
	}
	if bytes, e := storage.Read(TTMSessionId); e == nil {
		ctx.TTMSessionId = string(bytes)
	}
	if bytes, e := storage.Read(ConnectionType); e == nil {
		ctx.ConnectionType = string(bytes)
	}
	if bytes, e := storage.Read(SessionTime); e == nil {
		if date, err := time.Parse(time.RFC3339, string(bytes)); err == nil {
			ctx.SessionTime = date
			log.Printf("Restore session from: %v", date.Format(time.RFC3339))
		}
	}

	return &ctx, nil
}
