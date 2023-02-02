package setting

type LogSettingS struct {
	Path       string
	FilePrefix string
	Encoder    string
	Output     string
	LumberJack LumberJack
}

type LumberJack struct {
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
