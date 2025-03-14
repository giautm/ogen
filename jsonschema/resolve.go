package jsonschema

import (
	"context"

	"github.com/go-faster/errors"
	yaml "github.com/go-faster/yamlx"

	"github.com/ogen-go/ogen/internal/jsonpointer"
	"github.com/ogen-go/ogen/internal/location"
)

// ReferenceResolver resolves JSON schema references.
type ReferenceResolver interface {
	ResolveReference(ref string) (*RawSchema, error)
}

type resolver struct {
	ReferenceResolver
	file location.File
}

func (p *Parser) getResolver(loc string) (r resolver, rerr error) {
	r, ok := p.schemas[loc]
	if ok {
		return r, nil
	}

	raw, err := p.external.Get(context.TODO(), loc)
	if err != nil {
		return r, errors.Wrapf(err, "get %q", loc)
	}

	file := location.NewFile(loc, loc, raw)
	defer func() {
		if rerr != nil {
			rerr = &location.Error{
				File: file,
				Err:  rerr,
			}
		}
	}()

	var node yaml.Node
	if err := yaml.Unmarshal(raw, &node); err != nil {
		return r, errors.Wrap(err, "unmarshal")
	}

	r = resolver{
		ReferenceResolver: NewRootResolver(&node),
		file:              file,
	}
	p.schemas[loc] = r

	return r, nil
}

func (p *Parser) resolve(ref string, ctx *jsonpointer.ResolveCtx) (*Schema, error) {
	key, err := ctx.Key(ref)
	if err != nil {
		return nil, err
	}

	if s, ok := p.refcache[key]; ok {
		return s, nil
	}

	r, err := p.getResolver(key.Loc)
	if err != nil {
		return nil, err
	}

	if err := ctx.AddKey(key, r.file); err != nil {
		return nil, err
	}
	defer func() {
		// Drop the resolved ref to prevent false-positive infinite recursion detection.
		ctx.Delete(key)
	}()

	raw, err := r.ResolveReference(key.Ref)
	if err != nil {
		return nil, err
	}

	return p.parse1(raw, ctx, func(s *Schema) *Schema {
		s.Ref = ref
		p.refcache[key] = s
		return p.extendInfo(raw, s, p.file(ctx))
	})
}
