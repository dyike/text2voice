package server

import (
	"fmt"
	"text2voice/xf"

	"github.com/dyike/log"
)

type Server struct {
	opts *Options
}

type Options struct {
	OutDir      string
	Level       int
	TTSParams   string
	LoginParams string
	RedisAddr   string
	RedisPass   string
	Speed       int
}

type Speech struct {
	Id  string `json:"id"`
	Txt string `json:"txt"`
}

func New(opts *Options) *Server {
	return &Server{
		opts: opts,
	}
}

// TODO
func (s *Server) Start() {

}

// Once set txt and desPath
func (s *Server) Once(txt string, desPath string) error {
	log.Debug("tts:%s, login:%s", s.opts.TTSParams, s.opts.LoginParams)
	xf.SetTTSParams(s.opts.TTSParams)
	err := xf.Login(s.opts.LoginParams)
	if err != nil {
		return err
	}

	log.Debug("txt:%s, output path:%s", txt, desPath)
	err = xf.TextToSpeech(txt, desPath)
	if err != nil {
		return err
	}
	return nil
}

func setXF(speedLevel int, ttsParmas, loginParams string) error {
	if speedLevel < 1 || speedLevel > 10 {
		return fmt.Errorf("wrong speed level: %d, it should between 1 and 10", speedLevel)
	}

	sleepTime := 15000 * (speedLevel - 1)
	xf.SetSleep(sleepTime)

	xf.SetTTSParams(ttsParmas)

	err := xf.Login(loginParams)
	if err != nil {
		return err
	}
	return nil
}
