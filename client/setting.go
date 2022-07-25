package monitor

import (
	"path/filepath"
	"strings"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(confFile string) (*Setting, error) {
	confFileToken := strings.Split(filepath.Base(confFile), ".")
	confFileToken = confFileToken[:len(confFileToken)-1]

	vp := viper.New()
	vp.SetConfigName(strings.Join(confFileToken, "."))
	vp.AddConfigPath(filepath.Dir(confFile))
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

type ServerSetting struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  int
	WriteTimeout int
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
