package xf

/*
#cgo CFLAGS:-g -Wall -I./include
#cgo LDFLAGS:-L./libs -lmsc -lrt -ldl -lpthread
#include "convert.h"
*/
import "C"
import "fmt"

var ttsParams *C.char
var sleep C.int = C.int(0)

func SetTTSParams(params string) {
	ttsParams = C.CString(params)
}

func SetSleep(t int) {
	sleep = C.int(t)
}

func Login(loginParams string) error {
	ret := C.MSPLogin(nil, nil, C.CString(loginParams))
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("登陆失败，错误码： %d", int(ret))
	}
	return nil
}

func Logout() error {
	ret := C.MSPLogout()
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("注销失败，错误码： %d", int(ret))
	}
	return nil
}

func TextToSpeech(text, outPath string) error {
	ret := C.text_to_speech(C.CString(text), C.CString(outPath), ttsParams, sleep)
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("音频生成失败，错误码： %d", int(ret))
	}
	return nil
}
