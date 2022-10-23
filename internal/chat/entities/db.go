package entities

import (
	"fmt"
	"go-chat/pkg/logging"
	"reflect"
	"strings"
)

var logger = logging.Get()

type userFields struct {
	ID,
	Login,
	PasswordHash,
	IsDeleted,
	Chats string
}

type chatFields struct {
	ID,
	Name,
	Messages,
	Users string
}

type messageFields struct {
	ID,
	Text,
	Time,
	ChatID,
	UserID,
	User string
}

var UserFields = initFieldStruct(User{}, userFields{})
var ChatFields = initFieldStruct(Chat{}, chatFields{})
var MessageFields = initFieldStruct(Message{}, messageFields{})

type structConstraint interface {
	userFields | chatFields | messageFields
}

func dbFieldName(table Table, fieldName string) string {
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

func initFieldStruct[T structConstraint](table Table, fields T) *T {
	v := reflect.ValueOf(&fields)
	t := reflect.TypeOf(fields)

	for i := 0; i < t.NumField(); i++ {
		fieldName := dbFieldName(table, t.Field(i).Name)
		v.Elem().Field(i).SetString(fieldName)
	}

	return v.Interface().(*T)
}
