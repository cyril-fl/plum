package cliconfig

type Cmdparams struct {
	Use     string   `yaml:"use"`
	Short   string   `yaml:"short"`
	Long    string   `yaml:"long"`
	Example string   `yaml:"example"`
	Args    []Cmdarg `yaml:"args"`
}

type Cmdarg struct {
	Name      string `yaml:"name"`
	Shorthand string `yaml:"shorthand"`
	Value     string `yaml:"value"`
	Usage     string `yaml:"usage"`
	Type      string `yaml:"type"`
}
