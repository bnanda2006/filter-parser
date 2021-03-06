package filter

import (
	"fmt"
	"strings"
)

type Path struct {
	URIPrefix       string
	AttributeName   string
	SubAttribute    string
	ValueExpression Expression
}

// Expression is a type to assign to implemented expressions.
type Expression interface{}

// AttributeExpression is an Expression with a name, operator and value.
type AttributeExpression struct {
	Expression
	AttributePath   AttributePath
	CompareOperator Token
	CompareValue    string
}

type AttributePath struct {
	URIPrefix     string
	AttributeName string
	SubAttribute  string
}

type ValuePath struct {
	Expression
	URIPrefix       string
	AttributeName   string
	ValueExpression Expression
}

// UnaryExpression is an Expression with a token bound to a (child) expression X.
type UnaryExpression struct {
	Expression
	CompareOperator Token
	X               Expression
}

// BinaryExpression is an Expression with a token bound to two (child) expressions X and Y.
type BinaryExpression struct {
	Expression
	X               Expression
	CompareOperator Token
	Y               Expression
}

func splitURIPrefix(attrName string) (string, string) {
	var uriPrefix string
	uriParts := strings.Split(attrName, ":")
	if l := len(uriParts); l > 1 {
		uriPrefix = strings.Join(uriParts[:l-1], ":")
		attrName = uriParts[l-1]
	}
	return uriPrefix, attrName
}

func (path Path) String() string {
	attrName := path.AttributeName
	if path.URIPrefix != "" {
		attrName = fmt.Sprintf("%s:%s", path.URIPrefix, attrName)
	}
	if path.SubAttribute == "" {
		if path.ValueExpression != nil {
			return fmt.Sprintf("%s[%s]", attrName, path.ValueExpression)
		}
		return attrName
	}
	if path.ValueExpression == nil {
		return fmt.Sprintf("%s.%s", attrName, path.SubAttribute)
	}
	return fmt.Sprintf("%s[%s].%s", attrName, path.ValueExpression, path.SubAttribute)
}

func (expression AttributeExpression) String() string {
	return fmt.Sprintf("'%s %s %s'", expression.AttributePath, expression.CompareOperator, expression.CompareValue)
}

func (attributePath AttributePath) String() string {
	attrName := attributePath.AttributeName
	if attributePath.URIPrefix != "" {
		attrName = fmt.Sprintf("%s:%s", attributePath.URIPrefix, attrName)
	}
	if attributePath.SubAttribute != "" {
		return fmt.Sprintf("%s.%s", attrName, attributePath.SubAttribute)
	}
	return attrName
}

func (valuePath ValuePath) String() string {
	attrName := valuePath.AttributeName
	if valuePath.URIPrefix != "" {
		attrName = fmt.Sprintf("%s:%s", valuePath.URIPrefix, attrName)
	}
	return fmt.Sprintf("%s[%s]", attrName, valuePath.ValueExpression)
}

func (expression UnaryExpression) String() string {
	return fmt.Sprintf("%s %s", expression.CompareOperator, expression.X)
}

func (expression BinaryExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", expression.X, expression.CompareOperator, expression.Y)
}
