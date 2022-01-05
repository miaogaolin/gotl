package parser

import (
	"fmt"
	"path/filepath"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/miaogaolin/gotl/common/ddl-parser/gen"
)

func (p *Parser) FromContent(bytes []byte) (ret []*Table, err error) {
	defer func() {
		p := recover()
		if p != nil {
			switch e := p.(type) {
			case error:
				err = e
			default:
				err = fmt.Errorf("%+v", p)
			}
		}
	}()

	prefix := filepath.Base("")
	p.prefix = prefix
	inputStream := antlr.NewInputStream(string(bytes))
	caseChangingStream := newCaseChangingStream(inputStream, true)
	lexer := gen.NewMySqlLexer(caseChangingStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	mysqlParser := gen.NewMySqlParser(tokens)
	mysqlParser.RemoveErrorListeners()
	mysqlParser.AddErrorListener(p)

	visitor := &visitor{
		prefix: prefix,
		debug:  p.debug,
		logger: p.logger,
	}
	v := mysqlParser.Root().Accept(visitor)
	if v == nil {
		return empty, nil
	}

	createTables, ok := v.([]*CreateTable)
	if !ok {
		return empty, nil
	}

	for _, e := range createTables {
		ret = append(ret, e.Convert())
	}

	return
}
