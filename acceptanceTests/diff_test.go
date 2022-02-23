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

func SimpleDiff(d *DiffSuite, opts *differentiator.DifferentiatorOptions) {
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

	// info
	info := output[model.OAS_INFO_KEY]
	assert.Len(info, 2, "info should have 2 changes")
	// info[0]
	assert.Equal("update", info[0].Type)
	assert.Equal("title", info[0].Path)
	assert.Equal("Simple example", info[0].From)
	assert.Equal("Simple example 2", info[0].To)
	// info[1]
	assert.Equal("update", info[1].Type)
	assert.Equal("version", info[1].Path)
	assert.Equal("1.0.0", info[1].From)
	assert.Equal("1.1.0", info[1].To)

	// servers
	servers := output[model.OAS_SERVERS_KEY]
	assert.Len(servers, 3, "servers should have 3 changes")
	// servers[0]
	assert.Equal("delete", servers[0].Type)
	if opts.IncludeFilePath {
		assert.Equal(fmt.Sprintf("%s/%s#%s.0", d.absPath, FILE1, model.OAS_SERVERS_KEY), servers[0].Path)
	} else {
		assert.Equal(fmt.Sprintf("%s.0", model.OAS_SERVERS_KEY), servers[0].Path)
	}
	assert.Equal(model.Server{
		URL:         "https://test.com",
		Description: "some description",
	}, servers[0].From)
	assert.Equal(nil, servers[0].To)
	// servers[1]
	assert.Equal("create", servers[1].Type)
	if opts.IncludeFilePath {
		assert.Equal(fmt.Sprintf("%s/%s#%s.1", d.absPath, FILE2, model.OAS_SERVERS_KEY), servers[1].Path)
	} else {
		assert.Equal(fmt.Sprintf("%s.1", model.OAS_SERVERS_KEY), servers[1].Path)
	}
	assert.Equal(nil, servers[1].From)
	assert.Equal(model.Server{
		URL:         "http://gustavo.shipping.sock-shop",
		Description: "gustavo up9-demo-link all",
	}, servers[1].To)
	// servers[2]
	assert.Equal("create", servers[2].Type)
	if opts.IncludeFilePath {
		assert.Equal(fmt.Sprintf("%s/%s#%s.2", d.absPath, FILE2, model.OAS_SERVERS_KEY), servers[2].Path)
	} else {
		assert.Equal(fmt.Sprintf("%s.2", model.OAS_SERVERS_KEY), servers[2].Path)
	}
	assert.Equal(nil, servers[2].From)
	assert.Equal(model.Server{
		URL:         "https://test2.com",
		Description: "some description 2",
	}, servers[2].To)

	// paths
	paths := output[model.OAS_PATHS_KEY]
	assert.Len(paths, 1, "paths should have 1 change")
	// paths[0]
	paramName := "accept"
	assert.Equal("delete", paths[0].Type)
	if opts.IncludeFilePath {
		assert.Equal(fmt.Sprintf("%s/%s#%s./users.get.parameters.0", d.absPath, FILE1, model.OAS_PATHS_KEY), paths[0].Path)
	} else {
		assert.Equal(fmt.Sprintf("%s./users.get.parameters.0", model.OAS_PATHS_KEY), paths[0].Path)
	}
	assert.Equal(model.Parameter{
		Name: paramName,
		In:   "header",
		Schema: &model.SchemaRef{
			Ref:   "",
			Value: nil,
		},
	}, paths[0].From)
	assert.Equal(nil, paths[0].To)
}

func (d *DiffSuite) TestSimpleDiffWithFullFilePath() {
	opts := &differentiator.DifferentiatorOptions{
		IncludeFilePath: true,
	}
	d.diff = differentiator.NewDiff(d.vall, opts)
	SimpleDiff(d, opts)
}

func (d *DiffSuite) TestSimpleDiffWithoutFilePath() {
	opts := &differentiator.DifferentiatorOptions{
		IncludeFilePath: false,
	}
	d.diff = differentiator.NewDiff(d.vall, opts)
	SimpleDiff(d, opts)
}
