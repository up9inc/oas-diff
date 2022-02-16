package differentiator

import (
	"fmt"

	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

// make sure internalDiff struct implements the InternalDiff interface
//var _ InternalDiff = (*internalDiff)(nil)

type InternalDiff interface {
	Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator) (*internalDiff, error)
}

type internalDiff struct {
	key    string
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
	return lib.Diff(a, b, lib.DisableStructValues(), lib.SliceOrdering(false))
}

func (i *internalDiff) handleChanges(changes lib.Changelog) error {
	for _, c := range changes {
		path := fmt.Sprintf("%s.%s", i.key, c.Path[len(c.Path)-1])
		i.changelog = append(i.changelog,
			&changelog{
				Type: c.Type,
				Path: path,
				From: c.From,
				To:   c.To,
			},
		)
	}

	return nil
}

func (i *internalDiff) handleArrayChanges(data, data2 model.Array, changes lib.Changelog) error {
	for _, c := range changes {
		lastPath := c.Path[len(c.Path)-1]
		penultPath := lastPath
		if len(c.Path) > 1 {
			penultPath = c.Path[len(c.Path)-2]
		}
		path := fmt.Sprintf("%s.%s", i.key, lastPath)

		if c.Type == "create" || c.Type == "delete" {
			// create will display the path as the new element url value
			// url is the identifier for the servers array, let's get the index of the new element based on the last path
			// we have to figure out if it was created on the file1 or file2

			// file1
			index, _, err := data.SearchByIdentifier(lastPath)
			if err != nil {
				return err
			}
			if index != -1 {
				path = fmt.Sprintf("%s#%s.%d", i.filePath, i.key, index)

			} else {
				// file2
				index, _, err := data2.SearchByIdentifier(lastPath)
				if err != nil {
					return err
				}
				if index != -1 {
					path = fmt.Sprintf("%s#%s.%d", i.filePath2, i.key, index)
				}
			}
		}

		if c.Type == "update" {
			// url is the identifier for the servers array, let's get the index of the new element based on the penult path
			// we have to figure out if it was created on the file1 or file2

			// file1
			index, _, err := data.SearchByIdentifier(penultPath)
			if err != nil {
				return err
			}
			if index != -1 {
				path = fmt.Sprintf("%s#%s.%d.%s", i.filePath, i.key, index, lastPath)

			} else {
				// file2
				index, _, err := data2.SearchByIdentifier(penultPath)
				if err != nil {
					return err
				}
				if index != -1 {
					path = fmt.Sprintf("%s#%s.%d.%s", i.filePath2, i.key, index, lastPath)
				}
			}
		}

		i.changelog = append(i.changelog,
			&changelog{
				Type: c.Type,
				Path: path,
				From: c.From,
				To:   c.To,
			},
		)
	}

	return nil
}
