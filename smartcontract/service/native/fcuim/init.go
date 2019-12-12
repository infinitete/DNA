package fcuim

import (
	"git.fe-cred.com/idfor/idfor/smartcontract/service/native"
	"git.fe-cred.com/idfor/idfor/smartcontract/service/native/utils"
)

func InitFcuim() {
	native.Contracts[utils.IdforFcuimContractAddress] = RegisterFcuimContract
}

func RegisterFcuimContract(srvc *native.NativeService) {
	srvc.Register("registerFcuimScheme", registerFcuimScheme)
	srvc.Register("getFcuimSchemes", getFcuimSchemes)
	srvc.Register("getHello", getHello)
}
