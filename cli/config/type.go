package cliconfig

type Cmdparams struct {
	Use     string     `yaml:"use"`
	Short   string     `yaml:"short"`
	Long    string     `yaml:"long"`
	Example string     `yaml:"example"`
	Flags   []Cmdflags `yaml:"args"`
}

type Cmdflags struct {
	Name      string `yaml:"name"`
	Shorthand string `yaml:"shorthand"`
	Value     string `yaml:"value"`
	Usage     string `yaml:"usage"`
	Type      string `yaml:"type"`
}
