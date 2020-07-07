package f

import (
	"manager/func/module"
)

type Function struct {
	Uid int64
}

func (f *Function) Init() {

	module.Init()
}

func (f *Function) Close() {

}

var FF = new(Function)
