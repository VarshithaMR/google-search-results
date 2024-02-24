package props

type Properties struct {
	Server ServerProps `yaml:"server"`
}

type ServerProps struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	ContextRoot string `yaml:"context-root"`
	NoOfResults int    `yaml:"default-result-size"`
}
