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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/miaogaolin/gotl/common/ddl-parser/gen"
)

func TestVisitor_VisitSqlStatements(t *testing.T) {
	p := NewParser(WithDebugMode(true))
	accept := func(p *gen.MySqlParser, visitor *visitor) interface{} {
		root := p.Root()
		return root.Accept(visitor)
	}

	t.Run("empty", func(t *testing.T) {
		_, err := p.testMysqlSyntax("test.sql", accept, ``)
		assert.Nil(t, err)
	})

	t.Run("createDatabase", func(t *testing.T) {
		ret, err := p.testMysqlSyntax("test.sql", accept, "create database user")
		assert.Nil(t, err)
		assert.Equal(t, []*CreateTable(nil), ret)
	})

	t.Run("createSingleTable", func(t *testing.T) {
		ret, err := p.testMysqlSyntax("test.sql", accept, `
			create table if not exists user(
				id bigint(11) primary key not null default 0 comment '主键ID'
			)
		`)
		tables, ok := ret.([]*CreateTable)
		assert.True(t, ok)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(tables))
		assertCreateTableEqual(t, &CreateTable{
			Name: "user",
			Columns: []*ColumnDeclaration{
				{
					Name: "id",
					ColumnDefinition: &ColumnDefinition{
						DataType: &NormalDataType{tp: BigInt},
						ColumnConstraint: &ColumnConstraint{
							NotNull:         true,
							HasDefaultValue: true,
							AutoIncrement:   false,
							Primary:         true,
							Comment:         "主键ID",
						},
					},
				},
			},
		}, tables[0])
	})

	t.Run("createMultipleTables", func(t *testing.T) {
		ret, err := p.testMysqlSyntax("test.sql", accept, `
			-- user
			create table if not exists user(
				id bigint(11) primary key not null default 0 comment '主键ID'
			)
			
			-- student
			create table if not exists student(
				id bigint(11) primary key not null default 0 comment '主键ID',
				name varchar(10) key not null default '' comment '学生姓名'
			)
		`)
		tables, ok := ret.([]*CreateTable)
		assert.True(t, ok)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(tables))
		userTable := tables[0].Convert()
		studentTable := tables[1].Convert()
		assert.NotNil(t, userTable)
		assert.NotNil(t, studentTable)
		assert.Equal(t, &Table{
			Name: "user",
			Columns: []*Column{
				{
					Name:     "id",
					DataType: &NormalDataType{tp: BigInt},
					Constraint: &ColumnConstraint{
						NotNull:         true,
						HasDefaultValue: true,
						Primary:         true,
						Comment:         "主键ID",
					},
				},
			},
		}, userTable)
		assert.Equal(t, &Table{
			Name: "student",
			Columns: []*Column{
				{
					Name:     "id",
					DataType: &NormalDataType{tp: BigInt},
					Constraint: &ColumnConstraint{
						NotNull:         true,
						HasDefaultValue: true,
						Primary:         true,
						Comment:         "主键ID",
					},
				},
				{
					Name:     "name",
					DataType: &NormalDataType{tp: VarChar},
					Constraint: &ColumnConstraint{
						NotNull:         true,
						HasDefaultValue: true,
						Key:             true,
						Comment:         "学生姓名",
					},
				},
			},
		}, studentTable)
	})

	t.Run("ddlWithOtherSql", func(t *testing.T) {
		ret, err := p.testMysqlSyntax("test.sql", accept, `
			-- ddl create table
			create table if not exists user(
				id bigint(11) primary key not null comment 'id'
			)
			-- ddl create database
			create database foo;
			-- dml select
			select * from bar;
			-- dml update
			update foo set bar = 'test';
			-- dml insert
			insert into foo ('id','name') values ('1','bar');
		`)
		assert.Nil(t, err)
		assert.NotNil(t, ret)
		tables, ok := ret.([]*CreateTable)
		assert.True(t, ok)
		assert.Equal(t, 1, len(tables))
		assert.Equal(t, "user", tables[0].Name)
	})
}
