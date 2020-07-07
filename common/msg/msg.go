package msg

import (
	m"github.com/mikeqiao/newworld/manager"
	"reflect"
	"common/proto"
)

func init(){
	m.DefaultProcessor.RegisterMsg("Req", reflect.TypeOf(proto.Req{}))
	m.DefaultProcessor.RegisterMsg("Res",  reflect.TypeOf(proto.Res{}))
}

