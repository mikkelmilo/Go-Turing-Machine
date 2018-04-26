package parser

import "strconv"

type TMLError struct {
	line   int
	column int
	msg    string
}

func (s TMLError) String() string {
	line_str := strconv.Itoa(s.line)
	column_str := strconv.Itoa(s.column)
	return s.msg + " in line " + line_str + ", column " + column_str
}

func (s TMLError) Error() string {
	return s.String()
}
