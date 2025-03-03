package typed

import (
	"github.com/tendermint/starport/starport/pkg/field"
	"github.com/tendermint/starport/starport/pkg/multiformatname"
)

// Options ...
type Options struct {
	AppName    string
	AppPath    string
	ModuleName string
	ModulePath string
	OwnerName  string
	TypeName   multiformatname.Name
	MsgSigner  multiformatname.Name
	Fields     field.Fields
	Indexes    field.Fields
	NoMessage  bool
	IsIBC      bool
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
