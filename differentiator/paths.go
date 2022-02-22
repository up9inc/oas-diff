package differentiator

import (
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

// make sure we implement the InternalDiff interface
var _ InternalDiff = (*pathsDiff)(nil)

type pathsDiff struct {
	*internalDiff
	data  model.Paths
	data2 model.Paths
}

func NewPathsDiff() *pathsDiff {
	return &pathsDiff{
		internalDiff: NewInternalDiff(model.OAS_PATHS_KEY),
		data:         model.Paths{},
		data2:        model.Paths{},
	}
}

func (p *pathsDiff) Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator) error {
	var err error

	// schema
	err = p.schema.Build(validator)
	if err != nil {
		return err
	}

	// paths1
	p.filePath = jsonFile.GetPath()
	err = p.data.Parse(jsonFile)
	if err != nil {
		return err
	}

	// paths2
	p.filePath2 = jsonFile2.GetPath()
	err = p.data2.Parse(jsonFile2)
	if err != nil {
		return err
	}

	// TODO: Diff segments of the paths instead the entire map
	for k := range p.data {
		// Parameters
		err = p.handleArray(p.data[k].Parameters, p.data2[k].Parameters)
		if err != nil {
			return err
		}
		p.data[k].Parameters = nil
		p.data2[k].Parameters = nil

		// Servers
		err = p.handleArray(p.data[k].Servers, p.data2[k].Servers)
		if err != nil {
			return err
		}
		p.data[k].Servers = nil
		p.data2[k].Servers = nil

		// Operations

		// Connect
		if p.data[k].Connect != nil && p.data2[k].Connect != nil {
			// Connect Parameters
			err = p.handleArray(p.data[k].Connect.Parameters, p.data2[k].Connect.Parameters)
			if err != nil {
				return err
			}
			p.data[k].Connect.Parameters = nil
			p.data2[k].Connect.Parameters = nil
		}

		/* 		Connect
		   		Delete
		   		Get
		   		Head
		   		Options
		   		Patch
		   		Post
		   		Put
		   		Trace        */
	}

	// paths changelog
	changes, err := p.diff(p.data, p.data2)
	if err != nil {
		return err
	}

	// changelogs
	return p.handleChanges(changes)
}

func (p *pathsDiff) handleOperation(ops, ops2 *model.Operation) error {
	opsChanges, err := p.diff(ops, ops2)
	if err != nil {
		return err
	}
	err = p.handleChanges(opsChanges)
	if err != nil {
		return err
	}

	return nil
}

func (p *pathsDiff) handleArray(params, params2 model.Array) error {
	paramsChanges, err := p.diff(params, params2)
	if err != nil {
		return err
	}
	err = p.handleArrayChanges(params, params2, paramsChanges)
	if err != nil {
		return err
	}

	return nil
}
