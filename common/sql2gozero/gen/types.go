package gen

import (
	"github.com/miaogaolin/gotl/common/sql2gozero/template"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func genTypes(table Table, methods string, withCache bool) (string, error) {
	fields := table.Fields
	fieldsString, err := genFields(table, fields)
	if err != nil {
		return "", err
	}

	text, err := pathx.LoadTemplate(category, typesTemplateFile, template.Types)
	if err != nil {
		return "", err
	}

	output, err := util.With("types").
		Parse(text).
		Execute(map[string]interface{}{
			"withCache":             withCache,
			"method":                methods,
			"upperStartCamelObject": table.Name.ToCamel(),
			"fields":                fieldsString,
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
