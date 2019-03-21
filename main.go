package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"text2voice/server"

	"github.com/dyike/log"
)

var usageStr = `
Usage: text2voice [options]
单次合成模式选项:
    -t <text>                	待合成的文本
    -p <text_file_path>         待合成的文本路径
    -o <file>               	音频输出路径 
其他:
    -h                          查看帮助 
`

func configureLog(logFile, logLevel string) error {
	level := log.DEBUG
	switch strings.ToLower(logLevel) {
	case "debug":
		level = log.DEBUG
	case "info":
		level = log.INFO
	case "warn":
		level = log.WARN
	case "error":
		level = log.ERROR
	}

	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log.Set(level, file, log.Lshortfile|log.LstdFlags)
	return nil
}

type LogConf struct {
	Enable   bool   `json:"enable"`
	FileName string `json:"filename"`
	Level    string `json:"level"`
}

type TTSConf struct {
	VoiceName    string `json:"voice_name"`
	TextEncoding string `json:"text_encoding"`
	SampleRate   int    `json:"sample_rate"`
	Speed        int    `json:"speed"`
	Volume       int    `json:"volume"`
	Pitch        int    `json:"pitch"`
	Rdn          int    `json:"rdn"`
}

type Text2VoiceConfig struct {
	Log     LogConf `json:"log"`
	AppId   string  `json:"appid"`
	WorkDir string  `json:"work_dir"`
	Speed   int     `json:"speed"`
	TTS     TTSConf `json:"tts"`
}

func LoadConfig(fileName string) (*Text2VoiceConfig, error) {
	if len(fileName) == 0 {
		return nil, errors.New("配置文件为空")
	}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	text2voiceConfig := &Text2VoiceConfig{}
	err = decoder.Decode(text2voiceConfig)
	if err != nil {
		return nil, err
	}
	return text2voiceConfig, nil
}

func main() {
	var txt string
	var txtPath string
	var out string
	var help bool
	var confFile string

	flag.StringVar(&confFile, "c", "default.json", "配置文件")
	flag.StringVar(&txt, "t", "", "单次合成的文本")
	flag.StringVar(&txtPath, "p", "", "合成文本文件路径")
	flag.StringVar(&out, "o", "test.mp3", "单次合成的输出路径")
	flag.BoolVar(&help, "h", false, "Help")

	flag.Parse()

	if help {
		fmt.Printf("%s\n", usageStr)
		return
	}

	config, err := LoadConfig(confFile)
	if err != nil {
		fmt.Printf("加载配置文件失败:%v\n", err)
		return
	}

	logLevel := config.Log.Level
	logFile := config.Log.FileName

	if config.Log.Enable {
		err := configureLog(logFile, logLevel)
		if err != nil {
			log.Error("日志配置失败：%v", err)
		}
	}

	opts := &server.Options{}
	opts.TTSParams = fmt.Sprintf("voice_name = %s, text_encoding = %s, sample_rate= %d, speed = %d, volume = %d, pitch = %d, rdn = %d",
		config.TTS.VoiceName,
		config.TTS.TextEncoding,
		config.TTS.SampleRate,
		config.TTS.Speed,
		config.TTS.Volume,
		config.TTS.Pitch,
		config.TTS.Rdn)

	opts.LoginParams = fmt.Sprintf("appid = %s, work_dir = %s",
		config.AppId,
		config.WorkDir)

	opts.Speed = config.Speed

	s := server.New(opts)

	if txt != "" {
		if out == "" {
			out = txt + ".wav"
		}
		log.Debug("合成文本：%q，输出：%s", txt, out)

		if err := s.Once(txt, out); err != nil {
			log.Error("%v", err)
			return
		}
	}
}
