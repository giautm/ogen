package gen

import (
	"golang.org/x/xerrors"

	"github.com/ogen-go/ogen"
)

func (g *Generator) generateComponents() error {
	if err := g.generateComponentRequestBodies(); err != nil {
		return xerrors.Errorf("requestBodies: %w", err)
	}
	if err := g.generateComponentResponses(); err != nil {
		return xerrors.Errorf("responses: %w", err)
	}
	return nil
}

func (g *Generator) generateComponentRequestBodies() error {
	refs := make(map[string]ogen.RequestBody)
	for name, body := range g.spec.Components.RequestBodies {
		if body.Ref != "" {
			refs[name] = body
			continue
		}

		rbody, err := g.generateRequestBody(name, &body)
		if err != nil {
			return xerrors.Errorf("%s: %w", name, err)
		}

		g.requestBodies[name] = rbody
	}

	for name, body := range refs {
		rbody, err := g.generateRequestBody(name, &body)
		if err != nil {
			return xerrors.Errorf("%s: %w", name, err)
		}

		g.requestBodies[name] = rbody
	}

	return nil
}

func (g *Generator) generateComponentResponses() error {
	refs := make(map[string]ogen.Response)
	for name, resp := range g.spec.Components.Responses {
		if resp.Ref != "" {
			refs[name] = resp
			continue
		}

		r, err := g.generateResponse(name, resp)
		if err != nil {
			return xerrors.Errorf("%s: %w", name, err)
		}

		g.responses[name] = r
	}

	for name, resp := range refs {
		r, err := g.generateResponse(name, resp)
		if err != nil {
			return xerrors.Errorf("%s: %w", name, err)
		}

		g.responses[name] = r
	}

	return nil
}
