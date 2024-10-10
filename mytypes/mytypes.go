package mytypes

import (
	"strings"
)

type MyInt int
type MyString string
type MyBuilder struct {
	Contents strings.Builder
}

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (i MyString) Len() int {
	return len(i)
}

func (i MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

type StringUppercaser struct {
	Contents strings.Builder
}

func (t StringUppercaser) ToUpper() string {
	return strings.ToUpper(t.Contents.String())
}

func Double(x *int) {
	*x *= 2
}

func (n *MyInt) Double() {
	*n *= 2
}
