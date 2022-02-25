package differentiator

import (
	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

type InternalDiff interface {
	InternalDiff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator, opts DifferentiatorOptions) error
}

type internalDiff struct {
	opts   DifferentiatorOptions
	key    string
	schema *schema

	filePath  string
	filePath2 string

	changelog []*changelog
}

type identifier map[string]string

func NewInternalDiff(key string) *internalDiff {
	return &internalDiff{
		key:       key,
		schema:    NewSchema(key),
		changelog: make([]*changelog, 0),
	}
}

func (i *internalDiff) diff(a, b interface{}) (lib.Changelog, error) {
	return lib.Diff(a, b, lib.CustomValueDiffers(NewStringDiffer(i.opts)), lib.StructMapKeySupport(), lib.DisableStructValues(), lib.SliceOrdering(false))
}

// TODO: Include source file information in path - GJSON query?
func (i *internalDiff) handleChange(change lib.Change) {
	// TODO: Exclusion logic here

	//path := strings.Join(change.Path, ".")
	i.changelog = append(i.changelog,
		&changelog{
			Type: change.Type,
			Path: change.Path,
			From: change.From,
			To:   change.To,
		},
	)
}

func (i *internalDiff) handleChanges(changes lib.Changelog) {
	for _, c := range changes {
		i.handleChange(c)
	}
}

func (i *internalDiff) handleArrayChange(data, data2 model.Array, change lib.Change) (err error) {
	lastPath := change.Path[len(change.Path)-1]
	penultPath := lastPath
	if len(change.Path) > 1 {
		penultPath = change.Path[len(change.Path)-2]
	}
	var identifierValue string

	if change.Type == "create" || change.Type == "delete" {
		// create will display the path as the new element url value
		// the last path value is the identifier value of the array
		// file1 is always the base file
		// creation is always from file2
		// deletion is always from file1

		identifierValue = lastPath
		filePath := i.filePath

		if change.Type == "create" {
			filePath = i.filePath2

		}
		change.Path = i.buildArrayPath(change.Path, filePath)
	}

	// TODO: Find the source file/index of the updated element property - GJSON query?
	// ISSUE: The identifier will always be present on both files, we need more info than just the identifier to find the source of the update
	if change.Type == "update" {
		// the last path value is the property name that was updated
		// the penult path value is the identifier value of the array
		// we have to figure out if it was updated from file1 or file2

		// TODO: for now let's just assume file1 as the source of the update - when --include-file-path
		identifierValue = penultPath
		filePath := i.filePath
		change.Path = i.buildArrayPath(change.Path, filePath)
	}

	i.changelog = append(i.changelog,
		&changelog{
			Type:       change.Type,
			Path:       change.Path,
			Identifier: identifier{data.GetIdentifierName(): identifierValue},
			From:       change.From,
			To:         change.To,
		},
	)

	return nil
}

func (i *internalDiff) handleArrayChanges(data, data2 model.Array, changes lib.Changelog) error {
	for _, c := range changes {
		err := i.handleArrayChange(data, data2, c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *internalDiff) buildArrayPath(path []string, filePath string) []string {
	if i.opts.IncludeFilePath {
		// TODO: We really need to include the key?
		auxPath := []string{filePath, i.key}
		auxPath = append(auxPath, path...)
		return auxPath
	}
	return path
}
