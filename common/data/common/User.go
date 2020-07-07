package common

import (
	"github.com/mikeqiao/newworld/data"
	"strconv"
	"fmt"
)

type User struct {
	vip	uint32
	name	string
	uid	uint64
	level	uint32
	prefix string
	update *data.UpdateMod
}

func NewUser(uid uint64, prefix string, update *data.UpdateMod) *User{
	d := new(User)
	d.uid= uid
	d.prefix= prefix
	if nil != update{
		d.update= update
	}else{
		d.update= new(data.UpdateMod)
		table := "User_" + fmt.Sprint(d.uid)
		d.update.Init(table)
	}
	return d
}

func(this *User) Setvip(value uint32){
	this.update.AddData(this.prefix+"vip", value)
	this.vip = value
}

func(this *User) Getvip() uint32{
	return this.vip
}

func(this *User) Setname(value string){
	this.update.AddData(this.prefix+"name", value)
	this.name = value
}

func(this *User) Getname() string{
	return this.name
}

func(this *User) Setuid(value uint64){
	this.update.AddData(this.prefix+"uid", value)
	this.uid = value
}

func(this *User) Getuid() uint64{
	return this.uid
}

func(this *User) Setlevel(value uint32){
	this.update.AddData(this.prefix+"level", value)
	this.level = value
}

func(this *User) Getlevel() uint32{
	return this.level
}

func (this *User)InitDataParam(ks []string, d string) {
	if nil == this.update{
		return
	}
		if len(ks) <= 0 {
			return
		}
		tkey := ks[0]
		switch tkey {
		case "vip":
		dd, _:=strconv.ParseUint(d,10,64)
		dv:=uint32(dd)
			this.vip= dv
		case "name":
		dv:=d
			this.name= dv
		case "uid":
		dv, _:=strconv.ParseUint(d,10,64)
			this.uid= dv
		case "level":
		dd, _:=strconv.ParseUint(d,10,64)
		dv:=uint32(dd)
			this.level= dv
		}
}

func(this *User) Destroy(){
	this.update.DelData(this.prefix + "vip")
	this.update.DelData(this.prefix + "name")
	this.update.DelData(this.prefix + "uid")
	this.update.DelData(this.prefix + "level")
}

