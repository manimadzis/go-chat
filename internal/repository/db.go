package repository

import (
	"fmt"
	"go-chat/internal/domain"
	"go-chat/pkg/logging"
	"reflect"
	"strings"
)

var logger = logging.Get()

type userFields struct {
	ID           string
	Login        string
	PasswordHash string
	IsDeleted    string

	Messages string
	Chats    string
}

type chatFields struct {
	ID   string
	Name string
}

type messageFields struct {
	ID     string
	Text   string
	Time   string
	ChatID string
	UserID string
}

var UserFields = initFieldStruct(domain.User{}, userFields{})
var ChatFields = initFieldStruct(domain.Chat{}, chatFields{})
var MessageFields = initFieldStruct(domain.Message{}, messageFields{})

type structConstraint interface {
	userFields | chatFields | messageFields
}

func dbFieldName(table domain.Table, fieldName string) string {
	t := reflect.TypeOf(table)
	field, ok := t.FieldByName(fieldName)
	if !ok {
		logger.Panicf("No field %s in struct %s", fieldName, table.TableName())
	}

	if tagName := field.Tag.Get("db"); tagName != "" {
		fieldName = tagName
	} else if gormTags := field.Tag.Get("gorm"); gormTags != "" {
		if strings.Index(gormTags, "foreignKey") >= 0 {
			return field.Name
		}
	}

	return fmt.Sprintf("\"%s\".\"%s\"", table.TableName(), fieldName)
}

func initFieldStruct[T structConstraint](table domain.Table, fields T) *T {
	v := reflect.ValueOf(&fields)
	t := reflect.TypeOf(fields)

	for i := 0; i < t.NumField(); i++ {
		fieldName := dbFieldName(table, t.Field(i).Name)
		v.Elem().Field(i).SetString(fieldName)
	}

	return v.Interface().(*T)
}
