package differentiator

import (
	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
)

type infoDiff struct {
	key       string
	schema    *schema
	data      *model.Info
	data2     *model.Info
	changelog *changelog
}

func NewInfoDiff() *infoDiff {
	const key = model.OAS_INFO_KEY

	return &infoDiff{
		key:       key,
		schema:    NewSchema(key),
		changelog: NewChangelog(key),
	}
}

func (d *differentiator) infoDiff(jsonFile file.JsonFile, jsonFile2 file.JsonFile) error {
	var err error
	d.info = NewInfoDiff()

	// schema
	err = d.info.schema.Build(d.validator)
	if err != nil {
		return err
	}

	// info1
	d.info.data, err = model.ParseInfo(jsonFile)
	if err != nil {
		return err
	}

	// info2
	d.info.data2, err = model.ParseInfo(jsonFile2)
	if err != nil {
		return err
	}

	// info changelog
	d.info.changelog.Changelog, err = lib.Diff(d.info.data, d.info.data2)
	if err != nil {
		return err
	}

	return nil
}
