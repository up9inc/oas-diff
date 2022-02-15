package differentiator

import (
	"fmt"

	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

type serversDiff struct {
	*internalDiff
	data  *model.Servers
	data2 *model.Servers
}

func (s *serversDiff) Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator) (*serversDiff, error) {
	var err error
	s = &serversDiff{
		internalDiff: NewInternalDiff(model.OAS_SERVERS_KEY),
	}

	// schema
	err = s.schema.Build(validator)
	if err != nil {
		return nil, err
	}

	// servers1
	s.data, err = model.ParseServers(jsonFile)
	if err != nil {
		return nil, err
	}

	// servers2
	s.data2, err = model.ParseServers(jsonFile2)
	if err != nil {
		return nil, err
	}

	// servers changelog
	changes, err := lib.Diff(s.data, s.data2, lib.DisableStructValues(), lib.SliceOrdering(false))
	if err != nil {
		return nil, err
	}

	for _, c := range changes {
		lastPath := c.Path[len(c.Path)-1]
		penultPath := lastPath
		if len(c.Path) > 1 {
			penultPath = c.Path[len(c.Path)-2]
		}
		path := fmt.Sprintf("%s.%s", s.key, lastPath)

		if c.Type == "create" || c.Type == "delete" {
			// create will display the path as the new element url value
			// url is the identifier for the servers array, let's get the index of the new element based on the last path
			// we have to figure out if it was created on the file1 or file2

			// file1
			index, _ := s.data.FilterByURL(lastPath)
			if index != -1 {
				path = fmt.Sprintf("%s#%s.%d", jsonFile.GetPath(), s.key, index)

			} else {
				// file2
				index, _ := s.data2.FilterByURL(lastPath)
				if index != -1 {
					path = fmt.Sprintf("%s#%s.%d", jsonFile2.GetPath(), s.key, index)
				}
			}
		}

		if c.Type == "update" {
			// url is the identifier for the servers array, let's get the index of the new element based on the penult path
			// we have to figure out if it was created on the file1 or file2

			// file1
			index, _ := s.data.FilterByURL(penultPath)
			if index != -1 {
				path = fmt.Sprintf("%s#%s.%d.%s", jsonFile.GetPath(), s.key, index, lastPath)

			} else {
				// file2
				index, _ := s.data2.FilterByURL(penultPath)
				if index != -1 {
					path = fmt.Sprintf("%s#%s.%d.%s", jsonFile2.GetPath(), s.key, index, lastPath)
				}
			}
		}

		s.changelog = append(s.changelog,
			&changelog{
				Type: c.Type,
				Path: path,
				From: c.From,
				To:   c.To,
			},
		)
	}

	return s, nil
}
