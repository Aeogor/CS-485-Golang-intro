package myType

type MyType interface {
	encode() []byte
	decode() ([]byte, error)
	TypeId() uint32
}



type MyType_32 struct {
  ID_32 uint32
  Value int32
}

func (b_32 *MyType_32)encode() []byte{



  return nil
}
func (b_32 *MyType_32)decode() ([]byte, error){

  return nil,nil
}
func (b_32 *MyType_32)TypeId() uint32 {

  return 5
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




