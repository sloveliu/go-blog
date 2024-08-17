package setting

import "time"

type Server struct {
	Mode         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type App struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type Database struct {
	Type         string
	Username     string
	Password     string
	Host         string
	Port         string
	Name         string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

func (s *Setting) ReadSetting(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
