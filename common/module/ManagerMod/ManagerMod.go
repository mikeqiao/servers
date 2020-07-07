package ManagerMod

import (
	"common/proto"

	m "github.com/mikeqiao/newworld/manager"
	mod "github.com/mikeqiao/newworld/module"
)

var Mod *mod.Mod

func Init() {
	Mod := m.NewMod(0, "ManagerMod")
	Register()
	m.ModManager.Registe(Mod)
}

func Register() {
	Mod.Register("ServerStart", ServerStart, proto.Req{}, proto.Res{})
	Mod.Register("ServerStop", ServerStop, proto.Req{}, proto.Res{})
}
