package updateUser

import (
	dataUList "common/data/DataUList"
)

type UpdateUser struct {
	data	*dataUList.DataUList
	uid	uint64
}

func NewUpdateUser(uid uint64) *UpdateUser{
	d := new(UpdateUser)
	d.uid= uid
	return d
}

func(this *UpdateUser) Setdata(value *dataUList.DataUList){
	this.data = value
}

func(this *UpdateUser) Getdata() *dataUList.DataUList{
	return this.data
}

func (this *UpdateUser)UpdateData() {
 if nil != this.data {
		this.data.UpdateData()
	}
}

func (this *UpdateUser)Close() {
	this.UpdateData()
}

func(this *UpdateUser) Destroy(){
 if nil != this.data {
		this.data.Destroy()
	}
}

