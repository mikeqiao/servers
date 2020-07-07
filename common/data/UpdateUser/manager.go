package updateUser

import (
	"sync"
	"time"
)

var Manager *UpdateUserManager

type UpdateUserManager struct {
	closed	bool
	mutex sync.RWMutex
	data map[uint64]*UpdateUser
}

func init(){
	Manager := new(UpdateUserManager)
	Manager.data= make(map[uint64]*UpdateUser)
	go Manager.Update()
}

func (this *UpdateUserManager)AddData(d *UpdateUser){
	if  nil == d{
		return
	}
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.data[d.uid]= d
}

func (this *UpdateUserManager)DelData(uid uint64)bool{
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if  v,ok := this.data[uid];ok{
		if nil !=v{
			v.Close()
		}
		delete(this.data, uid)
		return true
	}
	return false
}

func (this *UpdateUserManager)GetData(uid uint64)*UpdateUser{
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	if  v,ok := this.data[uid];ok{
		return v
	}
	return nil
}

func AddData(data *UpdateUser){
	if nil== Manager || nil == data{
		return
	}
	Manager.AddData(data)
}

func DelData(uid uint64)bool{
	if nil== Manager {
		return false
	}
	return Manager.DelData(uid)
}

func GetData(uid uint64) *UpdateUser {
	if nil== Manager {
		return nil
	}
	return Manager.GetData(uid)
}

func (this *UpdateUserManager)Close(){
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.closed= true
	for k, v := range this.data{
		if nil !=v{
			v.Close()
		}
		delete(this.data, k)
	}
}

func (this *UpdateUserManager)Update(){
	t := time.Tick(500 * time.Millisecond)
	for _ = range t {
		this.mutex.RLock()
		if true == this.closed{
			this.mutex.RUnlock()
			break
		}
		for _, v := range this.data{
			if nil !=v{
				v.UpdateData()
			}
		}
		this.mutex.RUnlock()
	}
}

