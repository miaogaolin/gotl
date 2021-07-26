package gen

import (
	"github.com/miaogaolin/printlove-go/common/sql2gozero/parser"
	"github.com/tal-tech/go-zero/tools/goctl/config"
)

// NewDefaultGenerator creates an instance for defaultGenerator
func NewGenerator(pkg string, cfg *config.Config, opt ...Option) (*defaultGenerator, error) {

	generator := &defaultGenerator{cfg: cfg, pkg: pkg}
	var optionList []Option
	optionList = append(optionList, newDefaultOption())
	optionList = append(optionList, opt...)
	for _, fn := range optionList {
		fn(generator)
	}

	return generator, nil
}

// ret1: key-table name,value-code
func (g *defaultGenerator) GenFromDDContent(content []byte, withCache bool, database string) (map[string]string, error) {
	m := make(map[string]string)
	tables, err := parser.ParseContent(content, database)
	if err != nil {
		return nil, err
	}

	for _, e := range tables {
		code, err := g.genModel(*e, withCache)
		if err != nil {
			return nil, err
		}

		m[e.Name.Source()] = code
	}

	return m, nil
}
