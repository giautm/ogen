package gen

import (
	"github.com/ogen-go/ogen/gen/ir"
)

// genctx is a generation context.
type genctx struct {
	global *tstorage // readonly
	local  *tstorage
}

func (g *genctx) saveType(t *ir.Type) error {
	return g.local.saveType(t)
}

func (g *genctx) saveRef(ref string, t *ir.Type) error {
	return g.local.saveRef(ref, t)
}

func (g *genctx) lookupRef(ref string) (*ir.Type, bool) {
	if t, ok := g.global.refs[ref]; ok {
		return t, true
	}
	if t, ok := g.local.refs[ref]; ok {
		return t, true
	}
	return nil, false
}

func (g *genctx) saveResponse(ref string, r *ir.Response) error {
	return g.local.saveResponse(ref, r)
}

func (g *genctx) saveWType(ref string, t *ir.Type) error {
	return g.local.saveWType(ref, t)
}

func (g *genctx) lookupResponse(ref string) (*ir.Response, bool) {
	if r, ok := g.global.responses[ref]; ok {
		return r, true
	}
	if r, ok := g.local.responses[ref]; ok {
		return r, true
	}
	return nil, false
}

func (g *genctx) lookupWType(ref string) (*ir.Type, bool) {
	if t, ok := g.global.wtypes[ref]; ok {
		return t, true
	}
	if t, ok := g.local.wtypes[ref]; ok {
		return t, true
	}
	return nil, false
}
