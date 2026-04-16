package model

type Config struct {
	Global Global  `yaml:"global"`
	Tasks  []Tasks `yaml:"tasks"`
}
type Global struct {
	Path string `yaml:"path"`
}

type Tasks struct {
	QQID    string                 `yaml:"qq_id"`
	Private bool                   `yaml:"private"`
	Info    map[string]MessageInfo `yaml:"info"`
}
type MessageInfo struct {
	Hour    int      `yaml:"hour"`
	Minute  int      `yaml:"minute"`
	Message []string `yaml:"message"`
}
