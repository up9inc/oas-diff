package acceptanceTests

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/up9inc/oas-diff/differentiator"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

const (
	OAS_SCHEMA_FILE  = "../validator/oas31.json"
	FILE1            = "data/simple.json"
	FILE2            = "data/simple2.json"
	FILE_LOOSE1      = "data/simple_loose.json"
	FILE_LOOSE2      = "data/simple_loose2.json"
	FILE_HEADERS     = "data/headers.json"
	FILE_HEADERS2    = "data/headers2.json"
	FILE_RESPONSES   = "data/responses.json"
	FILE_RESPONSES2  = "data/responses2.json"
	FILE_OPERATIONS  = "data/operations.json"
	FILE_OPERATIONS2 = "data/operations2.json"
	FILE_SERVERS     = "data/servers.json"
	FILE_SERVERS2    = "data/servers2.json"
)

type DiffSuite struct {
	suite.Suite

	absPath string

	jsonFile1 file.JsonFile
	jsonFile2 file.JsonFile

	vall validator.Validator
	diff differentiator.Differentiator
}

func (d *DiffSuite) SetupTest() {
	var err error
	d.absPath, err = filepath.Abs("./")
	if err != nil {
		d.T().Error(err)
		return
	}

	// must be set on each test because of the options
	d.diff = nil
}

func TestDiffSuite(t *testing.T) {
	suite.Run(t, new(DiffSuite))
}

func validateDependencies(d *DiffSuite) {
	if d.jsonFile1 == nil {
		d.T().Error("jsonFile1 is nil, you must initialize jsonFile1 for each test")
		return
	}

	if d.jsonFile2 == nil {
		d.T().Error("jsonFile2 is nil, you must initialize jsonFile2 for each test")
		return
	}

	if d.vall == nil {
		d.T().Error("vall is nil, you must initialize vall for each test")
		return
	}

	if d.diff == nil {
		d.T().Error("diff is nil, you must initialize diff for each test")
		return
	}
}

func validateExecutionStatus(d *DiffSuite, output *differentiator.ChangelogOutput) {
	d.Assert().NotNil(d, "d is nil")
	d.Assert().NotNil(output, "output is nil")

	d.Assert().NotNil(output.ExecutionStatus, "executionStatus is nil")
	d.Assert().Equal(output.ExecutionStatus.BaseFilePath, d.jsonFile1.GetPath())
	d.Assert().Equal(output.ExecutionStatus.SecondFilePath, d.jsonFile2.GetPath())
	d.Assert().Greater(len(output.ExecutionStatus.StartTime), 1)
	d.Assert().Greater(len(output.ExecutionStatus.ExecutionTime), 1)
}

func validateChangeMapOutput(d *DiffSuite, output *differentiator.ChangelogOutput) {
	d.Assert().NotNil(d, "d is nil")
	d.Assert().NotNil(output, "output is nil")

	d.Assert().NotNil(output.Changelog, "changeMap is nil")
	d.Assert().Len(output.Changelog, 4, "changeMap len should be 4")
	d.Assert().NotNil(output.Changelog[model.OAS_INFO_KEY], fmt.Sprintf("failed to find changeMap key '%s'", model.OAS_INFO_KEY))
	d.Assert().NotNil(output.Changelog[model.OAS_SERVERS_KEY], fmt.Sprintf("failed to find changeMap key '%s'", model.OAS_SERVERS_KEY))
	d.Assert().NotNil(output.Changelog[model.OAS_PATHS_KEY], fmt.Sprintf("failed to find changeMap key '%s'", model.OAS_PATHS_KEY))
	d.Assert().NotNil(output.Changelog[model.OAS_WEBHOOKS_KEY], fmt.Sprintf("failed to find changeMap key '%s'", model.OAS_WEBHOOKS_KEY))
}

func SimpleDiff(d *DiffSuite, opts differentiator.DifferentiatorOptions) {
	validateDependencies(d)

	assert := d.Assert()

	_, err := d.jsonFile1.Read()
	assert.NoError(err)

	_, err = d.jsonFile2.Read()
	assert.NoError(err)

	output, err := d.diff.Diff(d.jsonFile1, d.jsonFile2)
	assert.NoError(err, fmt.Sprintf("diff error: %v", err))
	// ExecutionStatus
	validateExecutionStatus(d, output)
	// changeMap
	validateChangeMapOutput(d, output)

	// aux vars
	index := -1
	identifier := ""
	property := ""

	// info
	info := output.Changelog[model.OAS_INFO_KEY]
	assert.Len(info, 2, "info should have 2 changes")

	// info[0]
	index = 0
	assert.Equal("update", info[index].Type)
	assert.Len(info[index].Path, 1)
	assert.Equal([]string{"title"}, info[index].Path)
	assert.Equal("Simple example", info[index].From)
	assert.Equal("Simple example 2", info[index].To)

	// info[1]
	index = 1
	assert.Equal("update", info[index].Type)
	assert.Len(info[index].Path, 1)
	assert.Equal([]string{"version"}, info[index].Path)
	assert.Equal("1.0.0", info[index].From)
	assert.Equal("1.1.0", info[index].To)

	// servers
	servers := output.Changelog[model.OAS_SERVERS_KEY]
	assert.Len(servers, 4, "servers should have 4 changes")

	// servers[0] -> position 0
	index = 0
	identifier = "https://test.com"
	assert.Equal("delete", servers[index].Type)
	if opts.IncludeFilePath {
		assert.Len(servers[index].Path, 3)
		assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_SERVERS_KEY, identifier}, servers[index].Path)
	} else {
		assert.Len(servers[index].Path, 1)
		assert.Equal([]string{identifier}, servers[index].Path)
	}
	assert.Equal(differentiator.Identifier{
		model.Servers{}.GetIdentifierName(): identifier,
	}, servers[index].Identifier)
	assert.Equal(model.Server{
		URL:         identifier,
		Description: "some description",
	}, servers[index].From)
	assert.Equal(nil, servers[index].To)

	// servers[1] -> position 1
	index = 1
	identifier = "http://refael.shipping.sock-shop"
	property = "description"
	assert.Equal("update", servers[index].Type)
	if opts.IncludeFilePath {
		assert.Len(servers[index].Path, 4)
		assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_SERVERS_KEY, identifier, property}, servers[index].Path)
	} else {
		assert.Len(servers[index].Path, 2)
		assert.Equal([]string{identifier, property}, servers[index].Path)
	}
	assert.Equal(differentiator.Identifier{
		model.Servers{}.GetIdentifierName(): identifier,
	}, servers[index].Identifier)
	assert.Equal("refael salt bae", servers[index].From)
	assert.Equal("refael up9-demo-link all", servers[index].To)

	// servers[2] -> position 1
	index = 2
	identifier = "http://gustavo.shipping.sock-shop"
	assert.Equal("create", servers[index].Type)
	if opts.IncludeFilePath {
		assert.Len(servers[index].Path, 3)
		assert.Equal([]string{d.jsonFile2.GetPath(), model.OAS_SERVERS_KEY, identifier}, servers[index].Path)
	} else {
		assert.Len(servers[index].Path, 1)
		assert.Equal([]string{identifier}, servers[index].Path)
	}
	assert.Equal(differentiator.Identifier{
		model.Servers{}.GetIdentifierName(): identifier,
	}, servers[index].Identifier)
	assert.Equal(nil, servers[index].From)
	assert.Equal(model.Server{
		URL:         identifier,
		Description: "gustavo up9-demo-link all",
	}, servers[index].To)

	// servers[3] -> position 2
	index = 3
	identifier = "https://test2.com"
	assert.Equal("create", servers[index].Type)
	if opts.IncludeFilePath {
		assert.Len(servers[index].Path, 3)
		assert.Equal([]string{d.jsonFile2.GetPath(), model.OAS_SERVERS_KEY, identifier}, servers[index].Path)
	} else {
		assert.Len(servers[index].Path, 1)
		assert.Equal([]string{identifier}, servers[index].Path)
	}
	assert.Equal(differentiator.Identifier{
		model.Servers{}.GetIdentifierName(): identifier,
	}, servers[index].Identifier)
	assert.Equal(nil, servers[index].From)
	assert.Equal(model.Server{
		URL:         identifier,
		Description: "some description 2",
	}, servers[index].To)

	// paths
	paths := output.Changelog[model.OAS_PATHS_KEY]
	assert.Len(paths, 3, "paths should have 3 changes")

	// paths[0]
	index = 0
	identifier = "accept"
	basePath := []string{"/users", "get", "parameters"}
	assert.Equal("delete", paths[index].Type)
	if opts.IncludeFilePath {
		assert.Len(paths[index].Path, 6)
		assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], identifier}, paths[index].Path)
	} else {
		assert.Len(paths[index].Path, 4)
		assert.Equal([]string{basePath[0], basePath[1], basePath[2], identifier}, paths[index].Path)
	}
	assert.Equal(differentiator.Identifier{
		model.Parameters{}.GetIdentifierName(): identifier,
	}, paths[index].Identifier)
	assert.Equal(model.Parameter{
		Name: identifier,
		In:   "header",
		Schema: &model.Schema{
			Type: "string",
		},
	}, paths[index].From)
	assert.Equal(nil, paths[index].To)

	// paths[1]
	index = 1
	identifier = "id"
	basePath = []string{"/users", "get", "parameters", identifier, "schema", "pattern"}
	assert.Equal("update", paths[index].Type)
	if opts.IncludeFilePath {
		assert.Len(paths[index].Path, 8)
		assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3], basePath[4], basePath[5]}, paths[index].Path)
	} else {
		assert.Len(paths[index].Path, 6)
		assert.Equal(basePath, paths[index].Path)
	}
	assert.Equal(differentiator.Identifier{
		model.Parameters{}.GetIdentifierName(): identifier,
	}, paths[index].Identifier)
	assert.Equal(".+(_|-|\\.).+", paths[index].From)
	assert.Equal(".+(_|-ABC-|\\.).+", paths[index].To)

	// paths[2]
	index = 2
	identifier = "id"
	basePath = []string{"/users", "get", "parameters", identifier, "example"}
	assert.Equal("update", paths[index].Type)
	if opts.IncludeFilePath {
		assert.Len(paths[index].Path, 7)
		assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3], basePath[4]}, paths[index].Path)
	} else {
		assert.Len(paths[index].Path, 5)
		assert.Equal(basePath, paths[index].Path)
	}
	assert.Equal(differentiator.Identifier{
		model.Parameters{}.GetIdentifierName(): identifier,
	}, paths[index].Identifier)
	assert.Equal("some-uuid-maybe", paths[index].From)
	assert.Equal("custom uuid", paths[index].To)
}

func SimpleDiffLoose(d *DiffSuite, opts differentiator.DifferentiatorOptions) {
	validateDependencies(d)

	assert := d.Assert()

	assert.True(opts.Loose, "--loose flag should be true")

	_, err := d.jsonFile1.Read()
	assert.NoError(err)

	_, err = d.jsonFile2.Read()
	assert.NoError(err)

	output, err := d.diff.Diff(d.jsonFile1, d.jsonFile2)
	assert.NoError(err, fmt.Sprintf("diff error: %v", err))
	// ExecutionStatus
	validateExecutionStatus(d, output)
	// changeMap
	validateChangeMapOutput(d, output)

	// aux vars
	index := -1
	identifier := ""
	property := ""

	// info
	info := output.Changelog[model.OAS_INFO_KEY]

	if opts.ExcludeDescriptions {
		assert.Len(info, 1, "info should have 1 change")

		// info[0]
		index = 0
		assert.Equal("update", info[index].Type)
		assert.Len(info[index].Path, 1)
		assert.Equal([]string{"version"}, info[index].Path)
		assert.Equal("1.0.0", info[index].From)
		assert.Equal("1.1.0", info[index].To)
	} else {
		assert.Len(info, 2, "info should have 2 changes")

		// info[0]
		index = 0
		assert.Equal("update", info[index].Type)
		assert.Len(info[index].Path, 1)
		assert.Equal([]string{"description"}, info[index].Path)
		assert.Equal("", info[index].From)
		assert.Equal("new desc", info[index].To)

		// info[1]
		index = 1
		assert.Equal("update", info[index].Type)
		assert.Len(info[index].Path, 1)
		assert.Equal([]string{"version"}, info[index].Path)
		assert.Equal("1.0.0", info[index].From)
		assert.Equal("1.1.0", info[index].To)
	}

	// servers
	if !opts.ExcludeDescriptions {
		servers := output.Changelog[model.OAS_SERVERS_KEY]
		assert.Len(servers, 1, "servers should have 1 change")

		// servers[0] -> position 0
		index = 0
		identifier = "some url"
		property = "description"
		assert.Equal("update", servers[index].Type)
		if opts.IncludeFilePath {
			assert.Len(servers[index].Path, 4)
			assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_SERVERS_KEY, identifier, property}, servers[index].Path)
		} else {
			assert.Len(servers[index].Path, 2)
			assert.Equal([]string{identifier, property}, servers[index].Path)
		}
		assert.Equal(differentiator.Identifier{
			model.Servers{}.GetIdentifierName(): identifier,
		}, servers[index].Identifier)
		assert.Equal("SOME DESC", servers[index].From)
		assert.Equal("some desc updated", servers[index].To)
	}

	// paths
	paths := output.Changelog[model.OAS_PATHS_KEY]
	assert.Len(paths, 1, "paths should have 1 change")

	// paths[0]
	index = 0
	identifier = "accept"
	basePath := []string{"/users", "get", "parameters", identifier}
	property = "required"
	assert.Equal("update", paths[index].Type)
	if opts.IncludeFilePath {
		assert.Len(paths[index].Path, 7)
		assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3], property}, paths[index].Path)
	} else {
		assert.Len(paths[index].Path, 5)
		assert.Equal([]string{basePath[0], basePath[1], basePath[2], basePath[3], property}, paths[index].Path)
	}
	assert.Equal(differentiator.Identifier{
		model.Parameters{}.GetIdentifierName(): identifier,
	}, paths[index].Identifier)
	assert.Equal(false, paths[index].From)
	assert.Equal(true, paths[index].To)
}

func HeadersDiff(d *DiffSuite, opts differentiator.DifferentiatorOptions) {
	validateDependencies(d)

	assert := d.Assert()

	_, err := d.jsonFile1.Read()
	assert.NoError(err)

	_, err = d.jsonFile2.Read()
	assert.NoError(err)

	output, err := d.diff.Diff(d.jsonFile1, d.jsonFile2)
	assert.NoError(err, fmt.Sprintf("diff error: %v", err))
	// ExecutionStatus
	validateExecutionStatus(d, output)
	// changeMap
	validateChangeMapOutput(d, output)

	// aux vars
	index := -1

	// paths
	paths := output.Changelog[model.OAS_PATHS_KEY]
	if opts.Loose {
		assert.Len(paths, 1, "paths should have 1 change")

	} else {
		assert.Len(paths, 2, "paths should have 2 changes")
	}

	if opts.Loose {
		// paths[0]
		index = 0
		basePath := []string{"/example", "get", "responses", "200", "headers", "x-rate-limit", "description"}
		assert.Equal("update", paths[index].Type)
		if opts.IncludeFilePath {
			assert.Len(paths[index].Path, 9)
			assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3], basePath[4], basePath[5], basePath[6]}, paths[index].Path)
		} else {
			assert.Len(paths[index].Path, 7)
			assert.Equal(basePath, paths[index].Path)
		}
		assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
		assert.Equal("The number of allowed requests in the current period", paths[index].From)
		assert.Equal("new desc", paths[index].To)
	} else {
		// paths[0]
		index = 0
		key := "X-Rate-Limit"
		basePath := []string{"/example", "get", "responses", "200", "headers", key}
		assert.Equal("delete", paths[index].Type)
		if opts.IncludeFilePath {
			assert.Len(paths[index].Path, 8)
			assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3], basePath[4], basePath[5]}, paths[index].Path)
		} else {
			assert.Len(paths[index].Path, 6)
			assert.Equal(basePath, paths[index].Path)
		}
		assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
		assert.Equal(model.Header{
			Description: "The number of allowed requests in the current period",
			Schema: &model.Schema{
				Type: "integer",
			},
		}, paths[index].From)
		assert.Equal(nil, paths[index].To)

		// paths[1]
		index = 1
		basePath[len(basePath)-1] = strings.ToLower(key)
		assert.Equal("create", paths[index].Type)
		if opts.IncludeFilePath {
			assert.Len(paths[index].Path, 8)
			assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3], basePath[4], basePath[5]}, paths[index].Path)
		} else {
			assert.Len(paths[index].Path, 6)
			assert.Equal(basePath, paths[index].Path)
		}
		assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
		assert.Equal(nil, paths[index].From)
		assert.Equal(model.Header{
			Description: "new desc",
			Schema: &model.Schema{
				Type: "integer",
			},
		}, paths[index].To)
	}
}

func ResponsesDiff(d *DiffSuite, opts differentiator.DifferentiatorOptions) {
	validateDependencies(d)

	assert := d.Assert()

	_, err := d.jsonFile1.Read()
	assert.NoError(err)

	_, err = d.jsonFile2.Read()
	assert.NoError(err)

	output, err := d.diff.Diff(d.jsonFile1, d.jsonFile2)
	assert.NoError(err, fmt.Sprintf("diff error: %v", err))
	// ExecutionStatus
	validateExecutionStatus(d, output)
	// changeMap
	validateChangeMapOutput(d, output)

	// aux vars
	index := -1

	// paths
	paths := output.Changelog[model.OAS_PATHS_KEY]
	if opts.Loose {
		assert.Len(paths, 2, "paths should have 2 changes")
	} else {
		assert.Len(paths, 8, "paths should have 4 changes")
	}

	// paths[0]
	index = 0
	basePath := []string{"/example", "get", "responses", "200", "description"}
	assert.Equal("update", paths[index].Type)
	if opts.IncludeFilePath {
		assert.Len(paths[index].Path, 7)
		assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3], basePath[4]}, paths[index].Path)
	} else {
		assert.Len(paths[index].Path, 5)
		assert.Equal(basePath, paths[index].Path)
	}
	assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
	assert.Equal("A simple string response", paths[index].From)
	assert.Equal("the success response", paths[index].To)

	if opts.Loose {
		// update the index for the last path
		index = 1

	} else {
		index = 1
		basePath = []string{"/example", "get", "responses", "200", "content", "application/json"}
		assert.Equal("delete", paths[index].Type)
		if opts.IncludeFilePath {
			assert.Len(paths[index].Path, 8)
			assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3], basePath[4], basePath[5]}, paths[index].Path)
		} else {
			assert.Len(paths[index].Path, 6)
			assert.Equal(basePath, paths[index].Path)
		}
		assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
		assert.Equal(model.MediaType{
			Schema: &model.Schema{
				Type: "object",
			},
		}, paths[index].From)
		assert.Equal(nil, paths[index].To)

		index = 2
		basePath = []string{"/example", "get", "responses", "200", "content", "application/x-binary", "encoding", "base64"}
		assert.Equal("delete", paths[index].Type)
		if opts.IncludeFilePath {
			assert.Len(paths[index].Path, 10)
			assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3], basePath[4], basePath[5], basePath[6], basePath[7]}, paths[index].Path)
		} else {
			assert.Len(paths[index].Path, 8)
			assert.Equal(basePath, paths[index].Path)
		}
		assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
		assert.Equal(model.Encoding{
			ContentType: "base64",
		}, paths[index].From)
		assert.Equal(nil, paths[index].To)

		index = 3
		basePath = []string{"/example", "get", "responses", "200", "content", "application/x-binary", "encoding", "BASE64"}
		assert.Equal("create", paths[index].Type)
		if opts.IncludeFilePath {
			assert.Len(paths[index].Path, 10)
			assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3], basePath[4], basePath[5], basePath[6], basePath[7]}, paths[index].Path)
		} else {
			assert.Len(paths[index].Path, 8)
			assert.Equal(basePath, paths[index].Path)
		}
		assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
		assert.Equal(nil, paths[index].From)
		assert.Equal(model.Encoding{
			ContentType: "base64",
		}, paths[index].To)

		index = 4
		basePath = []string{"/example", "get", "responses", "200", "content", "Application/JSON"}
		assert.Equal("create", paths[index].Type)
		if opts.IncludeFilePath {
			assert.Len(paths[index].Path, 8)
			assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3], basePath[4], basePath[5]}, paths[index].Path)
		} else {
			assert.Len(paths[index].Path, 6)
			assert.Equal(basePath, paths[index].Path)
		}
		assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
		assert.Equal(nil, paths[index].From)
		assert.Equal(model.MediaType{
			Schema: &model.Schema{
				Type: "object",
			},
		}, paths[index].To)

		index = 5
		basePath = []string{"/example", "get", "responses", "200", "links", "address"}
		assert.Equal("delete", paths[index].Type)
		if opts.IncludeFilePath {
			assert.Len(paths[index].Path, 8)
			assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3], basePath[4], basePath[5]}, paths[index].Path)
		} else {
			assert.Len(paths[index].Path, 6)
			assert.Equal(basePath, paths[index].Path)
		}
		assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
		assert.Equal(model.Link{
			OperationID: "some-id",
		}, paths[index].From)
		assert.Equal(nil, paths[index].To)

		index = 6
		basePath = []string{"/example", "get", "responses", "200", "links", "Address"}
		assert.Equal("create", paths[index].Type)
		if opts.IncludeFilePath {
			assert.Len(paths[index].Path, 8)
			assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3], basePath[4], basePath[5]}, paths[index].Path)
		} else {
			assert.Len(paths[index].Path, 6)
			assert.Equal(basePath, paths[index].Path)
		}
		assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
		assert.Equal(nil, paths[index].From)
		assert.Equal(model.Link{
			OperationID: "some-id",
		}, paths[index].To)

		// update the index for the last path
		index = 7
	}

	basePath = []string{"/example", "get", "responses", "default"}
	assert.Equal("create", paths[index].Type)
	if opts.IncludeFilePath {
		assert.Len(paths[index].Path, 6)
		assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3]}, paths[index].Path)
	} else {
		assert.Len(paths[index].Path, 4)
		assert.Equal(basePath, paths[index].Path)
	}
	assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
	assert.Equal(nil, paths[index].From)
	assert.Equal(model.Response{
		Description: "the default response",
	}, paths[index].To)

}

func OperationsDiff(d *DiffSuite, opts differentiator.DifferentiatorOptions) {
	validateDependencies(d)

	assert := d.Assert()

	_, err := d.jsonFile1.Read()
	assert.NoError(err)

	_, err = d.jsonFile2.Read()
	assert.NoError(err)

	output, err := d.diff.Diff(d.jsonFile1, d.jsonFile2)
	assert.NoError(err, fmt.Sprintf("diff error: %v", err))
	// ExecutionStatus
	validateExecutionStatus(d, output)
	// changeMap
	validateChangeMapOutput(d, output)

	// aux vars
	index := -1

	// paths
	paths := output.Changelog[model.OAS_PATHS_KEY]
	if opts.Loose {
		assert.Len(paths, 1, "paths should have 1 change")
	} else {
		assert.Len(paths, 4, "paths should have 4 changes")
	}

	if opts.Loose {
		// paths[0]
		index = 0
		basePath := []string{"/example", "get"}
		assert.Equal("update", paths[index].Type)
		if opts.IncludeFilePath {
			assert.Len(paths[index].Path, 4)
			assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1]}, paths[index].Path)
		} else {
			assert.Len(paths[index].Path, 2)
			assert.Equal(basePath, paths[index].Path)
		}
		assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
		assert.Equal(nil, paths[index].From)
		assert.Equal(&model.Operation{}, paths[index].To)
	} else {
		// paths[0]
		index = 0
		basePath := []string{"/example"}
		assert.Equal("delete", paths[index].Type)
		if opts.IncludeFilePath {
			assert.Len(paths[index].Path, 3)
			assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0]}, paths[index].Path)
		} else {
			assert.Len(paths[index].Path, 1)
			assert.Equal(basePath, paths[index].Path)
		}
		assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
		assert.Equal(model.PathItem{
			Options: &model.Operation{},
			Patch:   &model.Operation{},
			Put:     &model.Operation{},
		}, paths[index].From)
		assert.Equal(nil, paths[index].To)

		// paths[1]
		index = 1
		basePath = []string{"/login", "post", "callbacks", "/logincallback"}
		assert.Equal("delete", paths[index].Type)
		if opts.IncludeFilePath {
			assert.Len(paths[index].Path, 6)
			assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3]}, paths[index].Path)
		} else {
			assert.Len(paths[index].Path, 4)
			assert.Equal(basePath, paths[index].Path)
		}
		assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
		assert.Equal(model.PathItem{
			Post: &model.Operation{},
		}, paths[index].From)
		assert.Equal(nil, paths[index].To)

		// paths[2]
		index = 2
		basePath = []string{"/login", "post", "callbacks", "/LoginCallback"}
		assert.Equal("create", paths[index].Type)
		if opts.IncludeFilePath {
			assert.Len(paths[index].Path, 6)
			assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3]}, paths[index].Path)
		} else {
			assert.Len(paths[index].Path, 4)
			assert.Equal(basePath, paths[index].Path)
		}
		assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
		assert.Equal(nil, paths[index].From)
		assert.Equal(model.PathItem{
			Post: &model.Operation{},
		}, paths[index].To)

		// paths[3]
		index = 3
		basePath = []string{"/Example"}
		assert.Equal("create", paths[index].Type)
		if opts.IncludeFilePath {
			assert.Len(paths[index].Path, 3)
			assert.Equal([]string{d.jsonFile1.GetPath(), model.OAS_PATHS_KEY, basePath[0]}, paths[index].Path)
		} else {
			assert.Len(paths[index].Path, 1)
			assert.Equal(basePath, paths[index].Path)
		}
		assert.Equal(differentiator.Identifier(differentiator.Identifier(nil)), paths[index].Identifier)
		assert.Equal(nil, paths[index].From)
		assert.Equal(model.PathItem{
			Get:     &model.Operation{},
			Options: &model.Operation{},
			Patch:   &model.Operation{},
			Put:     &model.Operation{},
		}, paths[index].To)

	}

}

func (d *DiffSuite) TestSimpleDiffWithFullFilePath() {
	d.jsonFile1 = file.NewJsonFile(FILE1)
	d.jsonFile2 = file.NewJsonFile(FILE2)

	d.vall = validator.NewValidator()
	err := d.vall.InitOAS31Schema(OAS_SCHEMA_FILE)
	if err != nil {
		d.T().Error(err)
		return
	}

	opts := differentiator.DifferentiatorOptions{
		IncludeFilePath: true,
	}
	d.diff = differentiator.NewDifferentiator(d.vall, opts)
	SimpleDiff(d, opts)
}

func (d *DiffSuite) TestSimpleDiff() {
	d.jsonFile1 = file.NewJsonFile(FILE1)
	d.jsonFile2 = file.NewJsonFile(FILE2)

	d.vall = validator.NewValidator()
	err := d.vall.InitOAS31Schema(OAS_SCHEMA_FILE)
	if err != nil {
		d.T().Error(err)
		return
	}

	opts := differentiator.DifferentiatorOptions{
		IncludeFilePath: false,
	}
	d.diff = differentiator.NewDifferentiator(d.vall, opts)
	SimpleDiff(d, opts)
}

func (d *DiffSuite) TestSimpleLooseDiff() {
	d.jsonFile1 = file.NewJsonFile(FILE_LOOSE1)
	d.jsonFile2 = file.NewJsonFile(FILE_LOOSE2)

	d.vall = validator.NewValidator()
	err := d.vall.InitOAS31Schema(OAS_SCHEMA_FILE)
	if err != nil {
		d.T().Error(err)
		return
	}

	opts := differentiator.DifferentiatorOptions{
		IncludeFilePath: false,
		Loose:           true,
	}
	d.diff = differentiator.NewDifferentiator(d.vall, opts)
	SimpleDiffLoose(d, opts)
}

func (d *DiffSuite) TestSimpleLooseDiffWithFullFilePath() {
	d.jsonFile1 = file.NewJsonFile(FILE_LOOSE1)
	d.jsonFile2 = file.NewJsonFile(FILE_LOOSE2)

	d.vall = validator.NewValidator()
	err := d.vall.InitOAS31Schema(OAS_SCHEMA_FILE)
	if err != nil {
		d.T().Error(err)
		return
	}

	opts := differentiator.DifferentiatorOptions{
		IncludeFilePath: true,
		Loose:           true,
	}
	d.diff = differentiator.NewDifferentiator(d.vall, opts)
	SimpleDiffLoose(d, opts)
}

func (d *DiffSuite) TestSimpleLooseDiffWithExcludeDescriptions() {
	d.jsonFile1 = file.NewJsonFile(FILE_LOOSE1)
	d.jsonFile2 = file.NewJsonFile(FILE_LOOSE2)

	d.vall = validator.NewValidator()
	err := d.vall.InitOAS31Schema(OAS_SCHEMA_FILE)
	if err != nil {
		d.T().Error(err)
		return
	}

	opts := differentiator.DifferentiatorOptions{
		Loose:               true,
		ExcludeDescriptions: true,
	}
	d.diff = differentiator.NewDifferentiator(d.vall, opts)
	SimpleDiffLoose(d, opts)
}

func (d *DiffSuite) TestHeadersDiff() {
	d.jsonFile1 = file.NewJsonFile(FILE_HEADERS)
	d.jsonFile2 = file.NewJsonFile(FILE_HEADERS2)

	d.vall = validator.NewValidator()
	err := d.vall.InitOAS31Schema(OAS_SCHEMA_FILE)
	if err != nil {
		d.T().Error(err)
		return
	}

	opts := differentiator.DifferentiatorOptions{
		IncludeFilePath: false,
	}
	d.diff = differentiator.NewDifferentiator(d.vall, opts)
	HeadersDiff(d, opts)
}

func (d *DiffSuite) TestHeadersLooseDiff() {
	d.jsonFile1 = file.NewJsonFile(FILE_HEADERS)
	d.jsonFile2 = file.NewJsonFile(FILE_HEADERS2)

	d.vall = validator.NewValidator()
	err := d.vall.InitOAS31Schema(OAS_SCHEMA_FILE)
	if err != nil {
		d.T().Error(err)
		return
	}

	opts := differentiator.DifferentiatorOptions{
		IncludeFilePath: false,
		Loose:           true,
	}
	d.diff = differentiator.NewDifferentiator(d.vall, opts)
	HeadersDiff(d, opts)
}

func (d *DiffSuite) TestResponsesDiff() {
	d.jsonFile1 = file.NewJsonFile(FILE_RESPONSES)
	d.jsonFile2 = file.NewJsonFile(FILE_RESPONSES2)

	d.vall = validator.NewValidator()
	err := d.vall.InitOAS31Schema(OAS_SCHEMA_FILE)
	if err != nil {
		d.T().Error(err)
		return
	}

	opts := differentiator.DifferentiatorOptions{
		IncludeFilePath: false,
	}
	d.diff = differentiator.NewDifferentiator(d.vall, opts)
	ResponsesDiff(d, opts)
}

func (d *DiffSuite) TestResponsesLooseDiff() {
	d.jsonFile1 = file.NewJsonFile(FILE_RESPONSES)
	d.jsonFile2 = file.NewJsonFile(FILE_RESPONSES2)

	d.vall = validator.NewValidator()
	err := d.vall.InitOAS31Schema(OAS_SCHEMA_FILE)
	if err != nil {
		d.T().Error(err)
		return
	}

	opts := differentiator.DifferentiatorOptions{
		IncludeFilePath: false,
		Loose:           true,
	}
	d.diff = differentiator.NewDifferentiator(d.vall, opts)
	ResponsesDiff(d, opts)
}

func (d *DiffSuite) TestOperationsDiff() {
	d.jsonFile1 = file.NewJsonFile(FILE_OPERATIONS)
	d.jsonFile2 = file.NewJsonFile(FILE_OPERATIONS2)

	d.vall = validator.NewValidator()
	err := d.vall.InitOAS31Schema(OAS_SCHEMA_FILE)
	if err != nil {
		d.T().Error(err)
		return
	}

	opts := differentiator.DifferentiatorOptions{
		IncludeFilePath: false,
	}
	d.diff = differentiator.NewDifferentiator(d.vall, opts)
	OperationsDiff(d, opts)
}

func (d *DiffSuite) TestOperationsLooseDiff() {
	d.jsonFile1 = file.NewJsonFile(FILE_OPERATIONS)
	d.jsonFile2 = file.NewJsonFile(FILE_OPERATIONS2)

	d.vall = validator.NewValidator()
	err := d.vall.InitOAS31Schema(OAS_SCHEMA_FILE)
	if err != nil {
		d.T().Error(err)
		return
	}

	opts := differentiator.DifferentiatorOptions{
		IncludeFilePath: false,
		Loose:           true,
	}
	d.diff = differentiator.NewDifferentiator(d.vall, opts)
	OperationsDiff(d, opts)
}
