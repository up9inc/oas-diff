package differentiator

import (
	"fmt"
	"strings"

	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

type InternalDiff interface {
	Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator, opts *DifferentiatorOptions) error
}

type internalDiff struct {
	key    string
	opts   *DifferentiatorOptions
	schema *schema

	filePath  string
	filePath2 string

	changelog []*changelog
}

func NewInternalDiff(key string) *internalDiff {
	return &internalDiff{
		key:       key,
		schema:    NewSchema(key),
		changelog: make([]*changelog, 0),
	}
}

func (i *internalDiff) diff(a, b interface{}) (lib.Changelog, error) {
	return lib.Diff(a, b, lib.StructMapKeySupport(), lib.DisableStructValues(), lib.SliceOrdering(false))
}

// TODO: Include source file information in path
func (i *internalDiff) handleChange(change lib.Change) {
	// TODO: Exclusion logic here

	path := strings.Join(change.Path, ".")
	i.changelog = append(i.changelog,
		&changelog{
			Type: change.Type,
			Path: path,
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
	path := fmt.Sprintf("%s.%s", i.key, lastPath)
	index := -1

	// data -> file1 -> base
	// data2 -> file2

	if change.Type == "create" || change.Type == "delete" {
		// create will display the path as the new element url value
		// url is the identifier for the servers array, let's get the index of the new element based on the last path
		// file1 is always the base file
		// creation is always from file2
		// deletion is always from file1

		var filePath string

		if change.Type == "create" {
			filePath = i.filePath2
			index, err = data2.SearchByIdentifier(lastPath)
		} else {
			filePath = i.filePath
			index, err = data.SearchByIdentifier(lastPath)
		}

		if err != nil {
			return err
		}
		if index != -1 {
			path = i.buildArrayPath(change.Type, change.Path, filePath, index)
			//path = fmt.Sprintf("%s#%s.%d", filePath, i.key, index)

		}
	}

	// TODO: Find the source file/index of the updated element property
	// ISSUE: The identifier will always be present on both files, we need more info than just the identifier to find the source of the update
	if change.Type == "update" {
		// url is the identifier for the servers array, let's get the index of the new element based on the penult path
		// we have to figure out if it was updated from file1 or file2

		// file1
		index, err := data.SearchByIdentifier(penultPath)
		if err != nil {
			return err
		}
		if index != -1 {
			path = i.buildArrayPath(change.Type, change.Path, i.filePath, index)
			//path = fmt.Sprintf("%s#%s.%d.%s", i.filePath, i.key, index, lastPath)

		} else {
			// file2
			index, err := data2.SearchByIdentifier(penultPath)
			if err != nil {
				return err
			}
			if index != -1 {
				path = i.buildArrayPath(change.Type, change.Path, i.filePath2, index)
				//path = fmt.Sprintf("%s#%s.%d.%s", i.filePath2, i.key, index, lastPath)
			}
		}
	}

	i.changelog = append(i.changelog,
		&changelog{
			Type: change.Type,
			Path: path,
			From: change.From,
			To:   change.To,
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

func (i *internalDiff) buildArrayPath(operation string, path []string, filePath string, index int) string {
	var auxPath string
	var result string

	// len == 2 -> array property and the identifier value
	if len(path) > 2 {
		// ignore the last path, the last path is the array identifier value
		for i := 0; i < len(path)-1; i++ {
			if i == 0 {
				auxPath = path[i]
				continue
			}
			auxPath = fmt.Sprintf("%s.%s", auxPath, path[i])
		}
	}

	// path len == 1 and the path is the identifier value
	if len(auxPath) == 0 {
		if i.opts.IncludeFilePath {
			result = fmt.Sprintf("%s#%s.%d", filePath, i.key, index)
		} else {
			result = fmt.Sprintf("%s.%d", i.key, index)
		}
	} else {
		// auxPath
		if i.opts.IncludeFilePath {
			result = fmt.Sprintf("%s#%s.%s.%d", filePath, i.key, auxPath, index)
		} else {
			result = fmt.Sprintf("%s.%s.%d", i.key, auxPath, index)
		}
	}

	// for update we need to include the updated property name on the end
	if operation == "update" {
		result = fmt.Sprintf("%s.%s", result, path[len(path)-1])
	}

	return result
}
