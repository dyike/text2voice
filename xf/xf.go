package xf

/*
#cgo CFLAGS:-g -Wall -I./include
#cgo LDFLAGS:-L./libs -lmsc -lrt -ldl -lpthread -lstdc++
#include "tts.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

/*
* rdn:           合成音频数字发音方式
* volume:        合成音频的音量
* pitch:         合成音频的音调
* speed:         合成音频对应的语速
* voice_name:    合成发音人
* sample_rate:   合成音频采样率
* text_encoding: 合成文本编码格式
*
 */
var ttsParams *C.char
var sleep C.int = C.int(0)

func SetTTSParams(params string) {
	ttsParams = C.CString(params)
}

func SetSleep(t int) {
	sleep = C.int(t)
}

func Login(loginParams string) error {
	l := C.CString(loginParams)
	defer C.free(unsafe.Pointer(l))

	ret := C.MSPLogin(nil, nil, l)
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("登录失败,错误码: %d", int(ret))
	}
	return nil
}

func Logout() error {
	ret := C.MSPLogout()
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("注销失败,错误码: %d", int(ret))
	}
	return nil
}

func TextToSpeech(text, outPath string) error {
	t := C.CString(text)
	o := C.CString(outPath)
	defer C.free(unsafe.Pointer(t))
	defer C.free(unsafe.Pointer(o))

	ret := C.text_to_speech(t, o, ttsParams, sleep)
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("音频生成失败，错误码： %d", int(ret))
	}
	return nil
}
