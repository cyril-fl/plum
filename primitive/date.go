package primitive

import "time"

// Datable stores creation and update timestamps for a value.
type Datable struct {
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp" bson:"createdAt" yaml:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp" bson:"updatedAt" yaml:"updatedAt" json:"updatedAt"`
	immutable bool      `bun:"-" bson:"-" yaml:"-" json:"-"`
}

// TemplateDatable configures the initial state of a Datable.
type TemplateDatable struct {
	Immutable bool
}

// Dated returns a Datable whose update timestamp changes when touched.
func Dated() Datable {
	return newdate(TemplateDatable{
		Immutable: false,
	})
}

// Steady returns a Datable whose update timestamp remains unchanged when touched.
func Steady() Datable {
	return newdate(TemplateDatable{
		Immutable: true,
	})
}

func newdate(params TemplateDatable) Datable {
	now := time.Now().UTC()

	return Datable{
		CreatedAt: now,
		UpdatedAt: now,
		immutable: params.Immutable,
	}
}

// GetcreatedAt returns the creation timestamp.
func (d *Datable) GetcreatedAt() time.Time {
	return d.CreatedAt
}

// GetupdatedAt returns the last update timestamp.
func (d *Datable) GetupdatedAt() time.Time {
	return d.UpdatedAt
}

// GetFormattedUpdated returns the last update timestamp formatted with format.
func (d *Datable) GetFormattedUpdated(format string) string {
	return d.UpdatedAt.Format(format)
}

// GetAge returns a copy of the Datable with the same timestamps and mutability.
func (d *Datable) GetAge() Datable {
	return Datable{
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		immutable: d.immutable,
	}
}

// Touch updates the last update timestamp unless the Datable is steady.
func (d *Datable) Touch() {
	if d.immutable {
		return
	}
	d.UpdatedAt = time.Now().UTC()
}
