package utils

type UtilsStruct struct {
	id int
}

func (u *UtilsStruct) GetID() int {
	return u.id
}

func (u *UtilsStruct) SetID(id int) {
	u.id = id
}

func NewUtilsStruct(id int) *UtilsStruct {
	return &UtilsStruct{id}
}
