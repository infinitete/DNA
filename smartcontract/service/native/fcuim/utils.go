package fcuim

import (
	"bytes"
	"errors"

	"git.fe-cred.com/idfor/idfor/core/states"
	"git.fe-cred.com/idfor/idfor/smartcontract/service/native"
	"git.fe-cred.com/idfor/idfor/smartcontract/service/native/utils"
)

const (
	flag_exist byte = 0x01
)

var fcuim_methods = []byte("idfor_fcuim_methods")
var fcuim_spliter = []byte(":")

func checkSchemeExistence(srvc *native.NativeService, encScheme []byte) bool {
	val, err := srvc.CacheDB.Get(encScheme)
	if err == nil {
		val, err := states.GetValueFromRawStorageItem(val)
		if err == nil {
			if len(val) > 0 && val[0] == flag_exist {
				return true
			}
		}
	}
	return false
}

func encodeScheme(schema []byte) ([]byte, error) {
	length := len(schema)
	if length == 0 || length > 255 {
		return nil, errors.New("encode Fcuim Scheme error: invalid schema length")
	}
	enc := append(utils.IdforFcuimContractAddress[:], byte(length))
	enc = append(enc, schema...)

	return enc, nil
}

func decodeScheme(data []byte) ([]byte, error) {
	prefix := len(utils.IdforFcuimContractAddress)
	size := len(data)
	if size < prefix || size != int(data[prefix]+1)+prefix {
		return nil, errors.New("decode Fcuim Scheme error: invalid data length")
	}

	return data[prefix+1:], nil
}

func joinScheme(schemes []byte, scheme []byte) []byte {
	if schemes == nil || len(schemes) == 0 {
		return scheme
	}

	return bytes.Join([][]byte{schemes, scheme}, fcuim_spliter)
}
