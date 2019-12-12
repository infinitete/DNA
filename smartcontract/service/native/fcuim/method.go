package fcuim

import (
	"bytes"
	"errors"

	"git.fe-cred.com/idfor/idfor/common/log"
	"git.fe-cred.com/idfor/idfor/common/serialization"
	"git.fe-cred.com/idfor/idfor/core/states"
	"git.fe-cred.com/idfor/idfor/smartcontract/service/native"
	"git.fe-cred.com/idfor/idfor/smartcontract/service/native/utils"
)

func registerFcuimScheme(srvc *native.NativeService) ([]byte, error) {
	log.Debug("registerFcuimScheme")
	log.Debug("srvc.Input", srvc.Input)

	args := bytes.NewBuffer(srvc.Input)
	arg0, err := serialization.ReadVarBytes(args)
	if err != nil {
		return utils.BYTE_FALSE, errors.New("register Fcuim Scheme error: parsing argument 0 failed")
	} else if len(arg0) == 0 {
		return utils.BYTE_FALSE, errors.New("register Fcuim Scheme error: invalid length of argument 0")
	}

	if bytes.Index(arg0, fcuim_spliter) > -1 {
		return utils.BYTE_FALSE, errors.New("register Fcuim Scheme error: scheme contains ':'")
	}

	// 获得已经注册的
	scheme, err := encodeScheme(arg0)
	if err != nil {
		return utils.BYTE_FALSE, errors.New("register Fcuim Schema error: " + err.Error())
	}

	if checkSchemeExistence(srvc, scheme) {
		return utils.BYTE_FALSE, errors.New("register Fcuim Schema error: Fcuim registered")
	}

	// Join pool
	// array [m1, m2, m3]
	// 返回的是解码过的
	schemes, err := getEncodedFcuimScheme(srvc)
	if err != nil {
		return nil, err
	}

	// flat bytes: m1:m2:m3
	flat_schemes := bytes.Join(append(schemes, scheme), fcuim_spliter)

	srvc.CacheDB.Put(fcuim_methods, flat_schemes)
	srvc.CacheDB.Put(scheme, states.GenRawStorageItem([]byte{flag_exist}))
	srvc.CacheDB.Commit()
	registerFcuimSchemeEvent(srvc, arg0)

	return utils.BYTE_TRUE, nil
}

func getHello(srvc *native.NativeService) ([]byte, error) {
	return []byte("hello"), nil
}

func getFcuimSchemes(srvc *native.NativeService) ([]byte, error) {
	log.Debug("getFcuimSchemes")

	bs, err := srvc.CacheDB.Get(fcuim_methods)
	if err != nil {
		getFcuimSchemesEvent(srvc, nil)
		return nil, err
	}
	if bs == nil {
		return nil, nil
	}

	raws := bytes.Split(bs, fcuim_spliter)

	schemes := [][]byte{}
	if len(raws) > 0 {
		for _, encodedScheme := range raws {
			scheme, err := decodeScheme(encodedScheme)
			if err != nil || len(scheme) == 0 {
				continue
			}
			schemes = append(schemes, scheme)
		}
	}

	scms := bytes.Join(schemes, fcuim_spliter)
	getFcuimSchemesEvent(srvc, scms)

	return scms, nil
}

func getEncodedFcuimScheme(srvc *native.NativeService) ([][]byte, error) {
	log.Debug("getEncodedFcuimSchemes")

	bs, err := srvc.CacheDB.Get(fcuim_methods)
	if err != nil {
		getFcuimSchemesEvent(srvc, nil)
		return nil, err
	}
	if bs == nil {
		return nil, nil
	}

	return bytes.Split(bs, fcuim_spliter), nil
}

func registerFcuimRecord(srvc *native.NativeService) {
	log.Debug("registerFcuimRecord")
}

func getFcuimRecord(srvc *native.NativeService) {
	log.Debug("getFcuimRecord")
}
