package expr

import (
	"fmt"
	"go-chat/internal/entities"
	"go-chat/pkg/logging"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
	"strings"
	"time"
)

var logger = logging.Get()

func And(expressions ...string) string {
	if len(expressions) == 1 {
		return expressions[0]
	}
	result := expressions[0]
	for i := 1; i < len(expressions); i++ {
		result += fmt.Sprintf(" AND %s", expressions[i])
	}
	return result
}

type commonQueryExpression struct {
	column     string
	useBetween bool
	from       string
	to         string

	useCompare   bool
	sign         string
	compareValue string

	useIn    bool
	inValues []string
}

func (c *commonQueryExpression) Build(builder clause.Builder) {
	stmt, ok := builder.(*gorm.Statement)
	if !ok {
		return
	}

	if c.useBetween {
		stmt.WriteString(fmt.Sprintf("%s BETWEEN %s AND %s", c.column, c.from, c.to))
	}
	if c.useCompare {
		stmt.WriteString(fmt.Sprintf("%s %s %s", c.column, c.sign, c.compareValue))
	}
	if c.useIn {
		list := strings.Join(c.inValues, ", ")
		stmt.WriteString(fmt.Sprintf("%s IN (%s)", c.column, list))
	}
}

func CommonQuery(column string) *commonQueryExpression {
	return &commonQueryExpression{column: column}
}

func toString(a any) string {
	switch a.(type) {
	case int, int32, int64, int16, int8, float32, float64:
		return fmt.Sprintf("%v", a)
	case time.Time:
		return fmt.Sprintf("'%v'", a.(time.Time).Format(time.RFC3339))
	default:
		return fmt.Sprintf("'%v'", a)
	}
}

func (c *commonQueryExpression) Between(from, to any) *commonQueryExpression {
	c.useBetween = true

	c.from = toString(from)
	c.to = toString(to)
	return c
}

func (c *commonQueryExpression) Eq(value any) *commonQueryExpression {
	c.useCompare = true
	c.sign = "="
	c.compareValue = toString(value)
	return c
}

func (c *commonQueryExpression) In(values []string) *commonQueryExpression {
	c.useIn = true
	c.inValues = values
	return c
}

func quoted(value string) string {
	return fmt.Sprintf("'%s'", value)
}

func doubleQuoted(value string) string {
	return fmt.Sprintf("\"%s\"", value)
}

type joinQuery struct {
	left  entities.Table
	right entities.Table

	expression string
	query      string
}

func (j *joinQuery) On(expression string) string {
	j.expression = expression
	j.build()
	return j.query
}

func (j *joinQuery) Eq(left, right string) string {
	j.expression = fmt.Sprintf("%s = %s", left, right)
	j.build()
	return j.query
}

func (j *joinQuery) To(left entities.Table) string {
	j.left = left
	j.build()
	return j.query
}

func columnWithAliasedTable(table entities.Table, column string) string {
	return fmt.Sprintf("%s.%s", doubleQuoted(reflect.TypeOf(table).Name()), doubleQuoted(column))
}

func (j *joinQuery) build() {
	joinTable := tableToAliasedTable(j.right)

	if j.left != nil {
		var foreignKey, references string
		tRight := reflect.TypeOf(j.right)
		rightName := typeName(tRight)

		t := reflect.TypeOf(j.left)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if typeName(field.Type) == rightName {
				gormTag := t.Field(i).Tag.Get("gorm")
				for _, tag := range strings.Split(gormTag, ";") {
					split := strings.Split(tag, ":")
					if len(split) == 2 {
						key, value := split[0], split[1]
						if key == "foreignKey" {
							foreignKey = value
						} else if key == "references" {
							references = value
						}
					}
				}

			}
		}

		if foreignKey == "" {
			logger.Errorf("Can't find foreignKey in gorm tag")
			return
		}
		if references == "" {
			logger.Errorf("Can't find references in gorm tag")
			return
		}

		field, _ := tRight.FieldByName(references)
		referencesTag := field.Tag.Get("db")

		field, _ = t.FieldByName(foreignKey)
		foreignKeyTag := field.Tag.Get("db")

		j.expression = fmt.Sprintf("%s = %s", columnWithAliasedTable(j.right, referencesTag), columnWithAliasedTable(j.left, foreignKeyTag))
	}

	j.query = fmt.Sprintf("JOIN %s ON %s", joinTable, j.expression)
}

func typeName(t reflect.Type) string {
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t.Name()
}

func Join(table entities.Table) *joinQuery {
	return &joinQuery{right: table}
}

func tableToAliasedTable(table entities.Table) string {
	alias := reflect.TypeOf(table).Name()
	return fmt.Sprintf("%s AS %s", doubleQuoted(table.TableName()), doubleQuoted(alias))
}
