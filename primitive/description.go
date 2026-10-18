package primitive

// Descriptible stores a name and a description.
type Descriptible struct {
	Name        string `bun:"name" bson:"name" yaml:"name" json:"name"`
	Description string `bun:"description" bson:"description" yaml:"description" json:"description"`
}

// TemplateDescriptible contains the values used to create a Descriptible.
type TemplateDescriptible struct {
	Name        string //`bson:"name" json:"name" yaml:"name"`
	Description string //`bson:"description" json:"description" yaml:"description"`
}

// Describe returns a Descriptible initialized from params.
func Describe(params TemplateDescriptible) Descriptible {
	return Descriptible{
		Name:        params.Name,
		Description: params.Description,
	}
}

// GetName returns the descriptive name.
func (d Descriptible) GetName() string {
	return d.Name
}

// GetDescription returns the descriptive text.
func (d Descriptible) GetDescription() string {
	return d.Description
}

// SetName replaces the descriptive name.
func (d *Descriptible) SetName(name string) {
	d.Name = name
}

// SetDescription replaces the descriptive text.
func (d *Descriptible) SetDescription(description string) {
	d.Description = description
}
