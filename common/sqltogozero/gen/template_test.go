package gen

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/miaogaolin/printlove-go/common/sqltogozero/template"
	"github.com/tal-tech/go-zero/tools/goctl/util"
)

func TestGenTemplates(t *testing.T) {
	err := util.InitTemplates(category, templates)
	assert.Nil(t, err)
	dir, err := util.GetTemplateDir(category)
	assert.Nil(t, err)
	file := filepath.Join(dir, "model-new.tpl")
	data, err := ioutil.ReadFile(file)
	assert.Nil(t, err)
	assert.Equal(t, string(data), template.New)
}

func TestRevertTemplate(t *testing.T) {
	name := "model-new.tpl"
	err := util.InitTemplates(category, templates)
	assert.Nil(t, err)

	dir, err := util.GetTemplateDir(category)
	assert.Nil(t, err)

	file := filepath.Join(dir, name)
	data, err := ioutil.ReadFile(file)
	assert.Nil(t, err)

	modifyData := string(data) + "modify"
	err = util.CreateTemplate(category, name, modifyData)
	assert.Nil(t, err)

	data, err = ioutil.ReadFile(file)
	assert.Nil(t, err)

	assert.Equal(t, string(data), modifyData)

	assert.Nil(t, RevertTemplate(name))

	data, err = ioutil.ReadFile(file)
	assert.Nil(t, err)
	assert.Equal(t, template.New, string(data))
}

func TestClean(t *testing.T) {
	name := "model-new.tpl"
	err := util.InitTemplates(category, templates)
	assert.Nil(t, err)

	assert.Nil(t, Clean())

	dir, err := util.GetTemplateDir(category)
	assert.Nil(t, err)

	file := filepath.Join(dir, name)
	_, err = ioutil.ReadFile(file)
	assert.NotNil(t, err)
}

func TestUpdate(t *testing.T) {
	name := "model-new.tpl"
	err := util.InitTemplates(category, templates)
	assert.Nil(t, err)

	dir, err := util.GetTemplateDir(category)
	assert.Nil(t, err)

	file := filepath.Join(dir, name)
	data, err := ioutil.ReadFile(file)
	assert.Nil(t, err)

	modifyData := string(data) + "modify"
	err = util.CreateTemplate(category, name, modifyData)
	assert.Nil(t, err)

	data, err = ioutil.ReadFile(file)
	assert.Nil(t, err)

	assert.Equal(t, string(data), modifyData)

	assert.Nil(t, Update())

	data, err = ioutil.ReadFile(file)
	assert.Nil(t, err)
	assert.Equal(t, template.New, string(data))
}
