package BooleanCodeEnum

/*
 @author king 409060350@qq.com
*/

type BooleanCode struct {
	Code []byte
	Name string
}

var bytes00 []byte = []byte{0x00}
var FALSE = BooleanCode{Code: bytes00, Name: "false"}
var bytes01 []byte = []byte{0x01}
var TRUE = BooleanCode{Code: bytes01, Name: "true"}
