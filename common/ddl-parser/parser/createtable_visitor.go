/*
 * MIT License
 *
 * Copyright (c) 2021 zeromicro
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 */

package parser

import (
	"strings"

	"github.com/miaogaolin/gotl/common/ddl-parser/gen"
)

type CreateTable struct {
	// Name describes the literal of table
	Name        string
	Columns     []*ColumnDeclaration
	Constraints []*TableConstraint
}

type ColumnDeclaration struct {
	Name             string
	ColumnDefinition *ColumnDefinition
}

type createDefinition struct {
	ColumnDeclaration *ColumnDeclaration
	TableConstraint   *TableConstraint
}

// visitCreateTable visits a parse tree produced by MySqlParser#createTable.
func (v *visitor) visitCreateTable(ctx gen.ICreateTableContext) *CreateTable {
	v.trace("VisitCreateTable")
	switch tx := ctx.(type) {
	case *gen.CopyCreateTableContext:
		v.panicWithExpr(tx.GetStart(),
			"Unsupported creating a table by copying from another table",
		)
	case *gen.QueryCreateTableContext:
		v.panicWithExpr(tx.GetStart(),
			"Unsupported creating a table by querying from another table",
		)
	case *gen.ColumnCreateTableContext:
		return v.visitColumnCreateTable(tx)
	}

	v.panicWithExpr(ctx.GetStart(), "Unknown creating a table")
	return nil
}

// visitColumnCreateTable visits a parse tree produced by MySqlParser#columnCreateTable.
func (v *visitor) visitColumnCreateTable(ctx *gen.ColumnCreateTableContext) *CreateTable {
	v.trace("VisitColumnCreateTable")
	var ret CreateTable
	tableName := ctx.TableName().GetText()
	tableName = strings.Trim(tableName, "`")
	tableName = strings.Trim(tableName, "'")
	replacer := strings.NewReplacer("\r", "", "\n", "")
	tableName = replacer.Replace(tableName)
	ret.Name = tableName
	if ctx.CreateDefinitions() != nil {
		if createDefinitionsContext, ok := ctx.CreateDefinitions().(*gen.CreateDefinitionsContext); ok {
			definitions := v.visitCreateDefinitions(createDefinitionsContext)
			v.convertCreateDefinition(definitions, &ret)
		}
	}

	return &ret
}

// visitCreateDefinitions visits a parse tree produced by MySqlParser#createDefinitions.
func (v *visitor) visitCreateDefinitions(ctx *gen.CreateDefinitionsContext) []*createDefinition {
	v.trace("VisitCreateDefinitions")
	var ret []*createDefinition
	for _, e := range ctx.AllCreateDefinition() {
		data := v.VisitCreateDefinition(e)
		if data == nil {
			continue
		}

		switch r := data.(type) {
		case *ColumnDeclaration:
			ret = append(ret, &createDefinition{
				ColumnDeclaration: r,
			})
		case *TableConstraint:
			ret = append(ret, &createDefinition{
				TableConstraint: r,
			})
		}
	}
	return ret
}

// VisitCreateDefinition visits a parse tree produced by MySqlParser#createDefinition.
func (v *visitor) VisitCreateDefinition(ctx gen.ICreateDefinitionContext) interface{} {
	v.trace("VisitCreateDefinition")
	switch tx := ctx.(type) {
	case *gen.ColumnDeclarationContext:
		var ret ColumnDeclaration
		ret.Name = v.visitUid(tx.Uid())
		iDefinitionContext := tx.ColumnDefinition()
		definitionContext, ok := iDefinitionContext.(*gen.ColumnDefinitionContext)
		if ok {
			out := v.VisitColumnDefinition(definitionContext)
			if cd, ok := out.(*ColumnDefinition); ok {
				ret.ColumnDefinition = cd
			}
		}

		return &ret
	case *gen.ConstraintDeclarationContext:
		if tx.TableConstraint() != nil {
			return v.visitTableConstraint(tx.TableConstraint())
		}
	}

	return nil
}

func (v *visitor) convertCreateDefinition(list []*createDefinition, table *CreateTable) {
	for _, e := range list {
		if e.ColumnDeclaration != nil {
			table.Columns = append(table.Columns, e.ColumnDeclaration)
		}
		if e.TableConstraint != nil {
			table.Constraints = append(table.Constraints, e.TableConstraint)
		}
	}
}

type Table struct {
	Name        string
	Columns     []*Column
	Constraints []*TableConstraint
}

type Column struct {
	Name       string
	DataType   DataType
	Constraint *ColumnConstraint
}

func (c *CreateTable) Convert() *Table {
	var ret Table
	ret.Name = c.Name
	for _, e := range c.Columns {
		definition := e.ColumnDefinition
		var data Column
		data.Name = e.Name
		if definition != nil {
			data.DataType = definition.DataType
			data.Constraint = definition.ColumnConstraint
		}
		ret.Columns = append(ret.Columns, &data)
	}

	ret.Constraints = c.Constraints
	return &ret
}
