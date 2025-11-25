package main

type Partial interface {
	Less(other interface{}) bool
}


// func (s *string) Less(other interface{}) bool {

// } 

type MyString string

// 使用数组长度来排序
func (s *MyString) Less(other interface{}) bool {
	switch v := other.(type) {
	case *MyString:
		return len(*s) < len(*v)
	case MyString:
		return len(*s) < len(v)
	default:
		return false
	}
}