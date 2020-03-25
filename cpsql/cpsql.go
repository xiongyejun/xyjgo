// compare sql
// 获取多表对比的sql
package cpsql

import (
	"errors"
	"strings"
)

// 多表对比
type compareData struct {
	tables          []string // 选择对比的表格
	fieldsCondition []string // 用来对比的条件列——图号、名称等
	fieldsData      []string // 对比输出的列 ——数量、单价等
	sqlOther        string   // 其他需要了解的，比如获取表格的时候，需要去除空白等等
}

func (me *compareData) GetSql() (ret string, err error) {
	if len(me.tables) < 2 {
		err = errors.New("请设置2个以上的表格。")
		return
	}

	if len(me.fieldsCondition) == 0 {
		err = errors.New("请设置对比条件字段，如型号等。")
		return
	}

	if len(me.fieldsData) == 0 {
		err = errors.New("请设置对比输出的列，如数量、单价等。")
		return
	}

	if me.sqlOther != "" {
		if !strings.HasPrefix(me.sqlOther, "where ") {
			err = errors.New("设置其他where条件，这个条件是针对每一表的，请以where xx开头。\n" + me.sqlOther)
			return
		}
	}
	strBase := "\n(" + me.getAllTablesConditionFieldsItems() + ") A\n"

	return me.leftJoin(strBase), nil
}

func (me *compareData) leftJoin(strBase string) (ret string) {
	iAlias := 'A'
	var tableAlias1, tableAlias2 string

	for i := range me.tables {
		tableAlias1 = string(iAlias)
		tableAlias2 = string(iAlias + 1)

		strBase = "select " + tableAlias1 + ".*," + me.getDataFiles(i, tableAlias2) +
			" from \n" + strBase + "\n left join " + me.tables[i] + " " + tableAlias2 +
			" on " + me.getOnstr(tableAlias1, tableAlias2)

		iAlias += 2
		if i < len(me.tables)-1 {
			strBase = "(" + strBase + ") " + string(iAlias)
		}
	}

	return strBase
}

// leftJoin的时候，table需要获取的字段
// Table1.F1 As Table1F1,Table1.F2 As Table1F2
func (me *compareData) getDataFiles(tableIndex int, tableAlias string) (ret string) {
	var arr []string = make([]string, len(me.fieldsData))

	// table名称存在xx.xx这种，替换为_
	// 目的是构建1个新的列不出错
	tmp := strings.Replace(me.tables[tableIndex], ".", "_", -1)
	for i := range me.fieldsData {
		// 没有的字段好像被设置了NULL，用f标识
		arr[i] = tableAlias + "." + me.fieldsData[i] + " as f" + tmp + me.fieldsData[i]
	}
	ret = strings.Join(arr, ",")
	return
}

// B.F1 = A.F1 And B.F2 = A.F2
func (me *compareData) getOnstr(tableAlias1, tableAlias2 string) (ret string) {
	var arr []string = make([]string, len(me.fieldsCondition))

	for i := range me.fieldsCondition {
		arr[i] = tableAlias1 + "." + me.fieldsCondition[i] + "=" + tableAlias2 + "." + me.fieldsCondition[i]
	}
	ret = strings.Join(arr, " and ")
	return
}

// 获取fieldsCondition字段，所有表格的不重复项目，作为后面left join的基准
func (me *compareData) getAllTablesConditionFieldsItems() (ret string) {
	str := strings.Join(me.fieldsCondition, ",")
	var arr []string = make([]string, len(me.tables))
	for i := range arr {
		arr[i] = "select " + str + " from " + me.tables[i]
		if me.sqlOther != "" {
			arr[i] += (" " + me.sqlOther)
		}
	}
	ret = strings.Join(arr, "\n union \n")
	return
}

// 添加对比的表格
func (me *compareData) AddTables(tablename ...string) {
	me.tables = append(me.tables, tablename...)
}

// 添加对比作为条件的字段
func (me *compareData) AddConditionFields(fields ...string) {
	me.fieldsCondition = append(me.fieldsCondition, fields...)
}

// 添加对比输出的字段
func (me *compareData) AddDataFields(fields ...string) {
	me.fieldsData = append(me.fieldsData, fields...)
}

// 设置其他where条件，这个条件是针对每一表的
func (me *compareData) SetOtherCondition(strwhere string) {
	me.sqlOther = strings.ToLower(strwhere)
}

func New() (ret *compareData) {
	return new(compareData)
}
