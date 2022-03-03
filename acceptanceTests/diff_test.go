package acceptanceTests

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/up9inc/oas-diff/differentiator"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

const (
	OAS_SCHEMA_FILE = "../validator/oas31.json"
	FILE1           = "data/simple.json"
	FILE2           = "data/simple2.json"
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

	d.jsonFile1 = file.NewJsonFile(FILE1)
	d.jsonFile2 = file.NewJsonFile(FILE2)

	d.vall = validator.NewValidator()
	err = d.vall.InitOAS31Schema(OAS_SCHEMA_FILE)
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

func SimpleDiff(d *DiffSuite, opts differentiator.DifferentiatorOptions) {
	if d.diff == nil {
		d.T().Error("diff is nil, you must initialize diff for each test")
		return
	}

	assert := d.Assert()

	_, err := d.jsonFile1.Read()
	assert.NoError(err)

	_, err = d.jsonFile2.Read()
	assert.NoError(err)

	output, err := d.diff.Diff(d.jsonFile1, d.jsonFile2)
	assert.NoError(err, fmt.Sprintf("diff error: %v", err))
	assert.NotNil(output, "changeMap is nil")
	assert.Len(output, 3, "changeMap len should be 3")
	assert.NotNil(output[model.OAS_INFO_KEY], fmt.Sprintf("failed to find changeMap key '%s'", model.OAS_INFO_KEY))
	assert.NotNil(output[model.OAS_SERVERS_KEY], fmt.Sprintf("failed to find changeMap key '%s'", model.OAS_SERVERS_KEY))
	assert.NotNil(output[model.OAS_PATHS_KEY], fmt.Sprintf("failed to find changeMap key '%s'", model.OAS_PATHS_KEY))

	// aux vars
	index := -1
	identifier := ""
	property := ""

	// info
	info := output[model.OAS_INFO_KEY]
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
	servers := output[model.OAS_SERVERS_KEY]
	assert.Len(servers, 4, "servers should have 4 changes")

	// servers[0] -> position 0
	index = 0
	identifier = "https://test.com"
	assert.Equal("delete", servers[index].Type)
	if opts.IncludeFilePath {
		assert.Len(servers[index].Path, 3)
		assert.Equal([]string{fmt.Sprintf("%s/%s", d.absPath, FILE1), model.OAS_SERVERS_KEY, identifier}, servers[index].Path)
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
		assert.Equal([]string{fmt.Sprintf("%s/%s", d.absPath, FILE1), model.OAS_SERVERS_KEY, identifier, property}, servers[index].Path)
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
		assert.Equal([]string{fmt.Sprintf("%s/%s", d.absPath, FILE2), model.OAS_SERVERS_KEY, identifier}, servers[index].Path)
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
		assert.Equal([]string{fmt.Sprintf("%s/%s", d.absPath, FILE2), model.OAS_SERVERS_KEY, identifier}, servers[index].Path)
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
	paths := output[model.OAS_PATHS_KEY]
	assert.Len(paths, 3, "paths should have 3 changes")

	// paths[0]
	index = 0
	identifier = "accept"
	basePath := []string{"/users", "get", "parameters"}
	assert.Equal("delete", paths[index].Type)
	if opts.IncludeFilePath {
		assert.Len(paths[index].Path, 6)
		assert.Equal([]string{fmt.Sprintf("%s/%s", d.absPath, FILE1), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], identifier}, paths[index].Path)
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
		assert.Equal([]string{fmt.Sprintf("%s/%s", d.absPath, FILE1), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3], basePath[4], basePath[5]}, paths[index].Path)
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
		assert.Equal([]string{fmt.Sprintf("%s/%s", d.absPath, FILE1), model.OAS_PATHS_KEY, basePath[0], basePath[1], basePath[2], basePath[3], basePath[4]}, paths[index].Path)
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

func (d *DiffSuite) TestSimpleDiffWithFullFilePath() {
	opts := differentiator.DifferentiatorOptions{
		IncludeFilePath: true,
	}
	d.diff = differentiator.NewDifferentiator(d.vall, opts)
	SimpleDiff(d, opts)
}

func (d *DiffSuite) TestSimpleDiffWithoutFilePath() {
	opts := differentiator.DifferentiatorOptions{
		IncludeFilePath: false,
	}
	d.diff = differentiator.NewDifferentiator(d.vall, opts)
	SimpleDiff(d, opts)
}
