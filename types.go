package main

import (
	"fmt"
)

type MyType interface {
	encode() []byte
	decode() ([]byte, error)
	TypeId() uint32
}



type MyType_32 struct {
  m_int32 uint32
}

func (b_32 *MyType_32)encode() []byte{

  return nil
}
func (b_32 *MyType_32)decode() ([]byte, error){

  return nil,nil
}
func (b_32 *MyType_32)TypeId() uint32 {

  return 
}

var _ MyType = (*MyType_32)(nil)



func SetType(typeId uint32, t *MyType) {

	return
}

func MakeType(TypeId int32) (t *MyType) {

	return
}

func Read(path string) (t MyType, err error) {

	return
}

func Write(path string, t MyType) (err error) {

	return
}

func main() {
	fmt.Println("Srinivas Lingutla")
	fmt.Println("CS 485 - HW1")
}


