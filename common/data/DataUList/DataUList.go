package dataUList

import (
	"github.com/mikeqiao/newworld/data"
	"strconv"
	"common/data/common"
	"strings"
	"fmt"
)

type DataUList struct {
	name3	string
	userList	map[uint64]*common.User
	mySData	map[uint64]uint32
	uid	uint64
	lang	string
	name2	string
	prefix string
	update *data.UpdateMod
}

func NewDataUList(uid uint64, prefix string, update *data.UpdateMod) *DataUList{
	d := new(DataUList)
	d.uid= uid
	d.prefix= prefix
	if nil != update{
		d.update= update
	}else{
		d.update= new(data.UpdateMod)
		table := "DataUList_" + fmt.Sprint(d.uid)
		d.update.Init(table)
	}
	return d
}

func(this *DataUList) Setname3(value string){
	this.update.AddData(this.prefix+"name3", value)
	this.name3 = value
}

func(this *DataUList) Getname3() string{
	return this.name3
}

func(this *DataUList) CreateuserListNewData(key uint64)(value *common.User){
	newdata := common.NewUser(this.uid, "this.prefixuserList.", this.update)
	this.userList[key] = newdata
	value = newdata
	return
}

func(this *DataUList) DeluserListData(key uint64){
	if v,ok:=this.userList[key]; ok{
		delete(this.userList, key)
		v.Destroy()
	}
}

func(this *DataUList) GetuserListDataByKey(key uint64) (value *common.User) {
	if v,ok:=this.userList[key]; ok{
		value = v
	}
	return
}

func(this *DataUList) GetuserListDataAll() (d map[uint64]*common.User){
	d = make(map[uint64]*common.User)
	for k, v := range this.userList{
		d[k] = v
	}
	return
}

func(this *DataUList) AddmySDataData(key uint64,value uint32){
 	keystr := fmt.Sprint(key)
	this.update.AddData(this.prefix + "mySData." + keystr, value)
	this.mySData[key] = value
}

func(this *DataUList) DelmySDataData(key uint64){
 	keystr := fmt.Sprint(key)
	this.update.DelData(this.prefix + "mySData." + keystr)
	if _,ok:=this.mySData[key]; ok{
		delete(this.mySData, key)
	}
}

func(this *DataUList) GetmySDataDataByKey(key uint64) (value uint32) {
	if v,ok:=this.mySData[key]; ok{
		value = v
	}
	return
}

func(this *DataUList) GetmySDataDataAll() (d map[uint64]uint32){
	d = make(map[uint64]uint32)
	for k, v := range this.mySData{
		d[k] = v
	}
	return
}

func(this *DataUList) Setuid(value uint64){
	this.update.AddData(this.prefix+"uid", value)
	this.uid = value
}

func(this *DataUList) Getuid() uint64{
	return this.uid
}

func(this *DataUList) Setlang(value string){
	this.update.AddData(this.prefix+"lang", value)
	this.lang = value
}

func(this *DataUList) Getlang() string{
	return this.lang
}

func(this *DataUList) Setname2(value string){
	this.update.AddData(this.prefix+"name2", value)
	this.name2 = value
}

func(this *DataUList) Getname2() string{
	return this.name2
}

func (this *DataUList)InitData() {
	if nil == this.update{
		return
	}
	data:= this.update.GetAllData()
	for k,d := range data {
			ks := strings.Split(k, ".")
		if len(ks) <= 0 {
			continue
		}
		tkey := ks[0]
		switch tkey {
		case "name3":
		dv:=d
			this.name3= dv
		case "userList":
			if nil == this.userList {
				this.userList=make(map[uint64]*common.User)
			}
			if len(ks) > 2 {
				d1 := ks[1]
		dv1, _:=strconv.ParseUint(d1,10,64)
			ts,ok := this.userList[dv1]
			if !ok || nil == ts {
				ts = common.NewUser(this.uid, "userList." + d1 + ".", this.update)
			}
				this.userList[dv1]= ts
			ts.InitDataParam(ks[2:],d)
			}
		case "mySData":
			if nil == this.mySData {
				this.mySData=make(map[uint64]uint32)
			}
			if len(ks) == 2 {
				d1 := ks[1]
		dv1, _:=strconv.ParseUint(d1,10,64)
		dd, _:=strconv.ParseUint(d,10,64)
		dv:=uint32(dd)
				this.mySData[dv1]= dv
			}
		case "uid":
		dv, _:=strconv.ParseUint(d,10,64)
			this.uid= dv
		case "lang":
		dv:=d
			this.lang= dv
		case "name2":
		dv:=d
			this.name2= dv
		}
	}
}

func (this *DataUList)UpdateData() {
	if nil != this.update{
		this.update.Update()
	}
}

func (this *DataUList)Close() {
	this.UpdateData()
}

func(this *DataUList) Destroy(){
	this.update.DelData(this.prefix + "name3")
	for _,v:=range this.userList{
		if nil != v {
			v.Destroy()
		}
	}
	for k,_:=range this.mySData{
		key := this.prefix + fmt.Sprint(k)
		this.update.DelData(key)
	}
	this.update.DelData(this.prefix + "uid")
	this.update.DelData(this.prefix + "lang")
	this.update.DelData(this.prefix + "name2")
}

