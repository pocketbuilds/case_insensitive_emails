package case_insensitive_emails

import (
	"strings"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbuilds/xpb"
)

func init() {
	xpb.Register(&Plugin{})
}

type Plugin struct{}

func (p *Plugin) Name() string {
	return "case_insensitive_emails"
}

var version string

func (p *Plugin) Version() string {
	return version
}

func (p *Plugin) Description() string {
	return "Automatically lowercase emails."
}

func (p *Plugin) Init(app core.App) error {
	app.OnModelCreate().BindFunc(p.lowercaseAuthRecordEmail)
	app.OnModelUpdate().BindFunc(p.lowercaseAuthRecordEmail)
	app.OnRecordAuthWithPasswordRequest().BindFunc(p.requeryAuthRecordWithLowercaseEmail)
	return nil
}

func (p *Plugin) lowercaseAuthRecordEmail(e *core.ModelEvent) error {
	record, ok := e.Model.(*core.Record)
	if ok && record.Collection() != nil && record.Collection().IsAuth() {
		record.SetEmail(strings.ToLower(record.Email()))
	}
	return e.Next()
}

func (p *Plugin) requeryAuthRecordWithLowercaseEmail(e *core.RecordAuthWithPasswordRequestEvent) error {
	if e.Record == nil {
		e.Record, _ = e.App.FindAuthRecordByEmail(e.Collection.Name, strings.ToLower(e.Identity))
	}
	return e.Next()
}
