package fcuim

import (
	"strings"

	"git.fe-cred.com/idfor/idfor/smartcontract/event"
	"git.fe-cred.com/idfor/idfor/smartcontract/service/native"
)

func newEvent(srvc *native.NativeService, st interface{}) {
	e := event.NotifyEventInfo{}
	e.ContractAddress = srvc.ContextRef.CurrentContext().ContractAddress
	e.States = st
	srvc.Notifications = append(srvc.Notifications, &e)
	return
}

func registerFcuimSchemeEvent(srvc *native.NativeService, scheme []byte) {
	newEvent(srvc, []string{"RegisterFcuimScheme", string(scheme)})
}

func getFcuimSchemesEvent(srvc *native.NativeService, schemes []byte) {
	scms := strings.Split(string(schemes), ":")
	newEvent(srvc, []interface{}{"GetFcuimSchemes", true, scms})
}
