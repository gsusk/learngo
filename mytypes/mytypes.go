package mytypes

type MyInt int
type MyString string

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (i MyString) Len() int {
	return len(i)
}
