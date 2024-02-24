package props

type Properties struct {
	Server ServerProps `yaml:"server"`
}

type ServerProps struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	ContextRoot string `yaml:"context-root"`
	NoOfResults string `yaml:"default-result-size"`
}
