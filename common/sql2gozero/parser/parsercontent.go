package parser

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/miaogaolin/gotl/common/ddl-parser/parser"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
)

// Parse parses ddl into golang structure
func ParseContent(content []byte, database string) ([]*Table, error) {
	p := parser.NewParser()
	tables, err := p.FromContent(content)
	if err != nil {
		return nil, err
	}

	indexNameGen := func(column ...string) string {
		return strings.Join(column, "_")
	}

	prefix := filepath.Base("")
	var list []*Table
	for _, e := range tables {
		columns := e.Columns

		var (
			primaryColumnSet = collection.NewSet()

			primaryColumn string
			uniqueKeyMap  = make(map[string][]string)
			normalKeyMap  = make(map[string][]string)
		)

		for _, column := range columns {
			if column.Constraint != nil {
				if column.Constraint.Primary {
					primaryColumnSet.AddStr(column.Name)
				}

				if column.Constraint.Unique {
					indexName := indexNameGen(column.Name, "unique")
					uniqueKeyMap[indexName] = []string{column.Name}
				}

				if column.Constraint.Key {
					indexName := indexNameGen(column.Name, "idx")
					uniqueKeyMap[indexName] = []string{column.Name}
				}
			}
		}

		for _, e := range e.Constraints {
			if len(e.ColumnPrimaryKey) > 1 {
				return nil, fmt.Errorf("%s: unexpected join primary key", prefix)
			}

			if len(e.ColumnPrimaryKey) == 1 {
				primaryColumn = e.ColumnPrimaryKey[0]
				primaryColumnSet.AddStr(e.ColumnPrimaryKey[0])
			}

			if len(e.ColumnUniqueKey) > 0 {
				list := append([]string(nil), e.ColumnUniqueKey...)
				list = append(list, "unique")
				indexName := indexNameGen(list...)
				uniqueKeyMap[indexName] = e.ColumnUniqueKey
			}
		}

		if primaryColumnSet.Count() > 1 {
			return nil, fmt.Errorf("%s: unexpected join primary key", prefix)
		}

		primaryKey, fieldM, err := convertColumns(columns, primaryColumn)
		if err != nil {
			return nil, err
		}

		var fields []*Field
		// sort
		for _, c := range columns {
			field, ok := fieldM[c.Name]
			if ok {
				fields = append(fields, field)
			}
		}

		var (
			uniqueIndex = make(map[string][]*Field)
			normalIndex = make(map[string][]*Field)
		)

		for indexName, each := range uniqueKeyMap {
			for _, columnName := range each {
				uniqueIndex[indexName] = append(uniqueIndex[indexName], fieldM[columnName])
			}
		}

		for indexName, each := range normalKeyMap {
			for _, columnName := range each {
				normalIndex[indexName] = append(normalIndex[indexName], fieldM[columnName])
			}
		}

		checkDuplicateUniqueIndex(uniqueIndex, e.Name)

		list = append(list, &Table{
			Name:        stringx.From(e.Name),
			Db:          stringx.From(database),
			PrimaryKey:  primaryKey,
			UniqueIndex: uniqueIndex,
			Fields:      fields,
		})
	}

	return list, nil
}
