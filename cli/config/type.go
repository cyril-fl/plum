package cliconfig

type Cmdparams struct {
	Use     string     `yaml:"use" json:"use"`
	Short   string     `yaml:"short" json:"short"`
	Long    string     `yaml:"long" json:"long"`
	Example string     `yaml:"example" json:"example"`
	Flags   []Cmdflags `yaml:"args" json:"args"`
}

type Cmdflags struct {
	Name      string `yaml:"name" json:"name"`
	Shorthand string `yaml:"shorthand" json:"shorthand"`
	Value     string `yaml:"value" json:"value"`
	Usage     string `yaml:"usage" json:"usage"`
	Type      string `yaml:"type" json:"type"`
}
