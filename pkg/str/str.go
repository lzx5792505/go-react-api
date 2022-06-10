package str

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

// Plural 转为复数 user -> users
func Plural(str string) string {
	return pluralize.NewClient().Plural(str)
}

// Singular 转为单数 users -> user
func Singular(str string) string {
	return pluralize.NewClient().Singular(str)
}

// Snake 转为 snake_case，如 TopicComment -> topic_comment
func Snake(str string) string {
	return strcase.ToSnake(str)
}

// Camel 转为 CamelCase，如 topic_comment -> TopicComment
func Camel(str string) string {
	return strcase.ToCamel(str)
}

// LowerCamel 转为 lowerCamelCase，如 TopicComment -> topicComment
func LowerCamel(str string) string {
	return strcase.ToLowerCamel(str)
}
