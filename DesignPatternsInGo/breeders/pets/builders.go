package pets

import "errors"

// PetInterface defines the methods that we want our builder to have. These are
// used to set the fields in teh Pet type, and to build the final product.
// Everything except the Build() function returns the type *Pet because we are
// going to implement the fluent interface.
type PetInterface interface {
	SetSpecies(s string) *Pet
	SetBreed(s string) *Pet
	SetMinWeight(i int) *Pet
	SetMaxWeight(i int) *Pet
	SetWeight(i int) *Pet
	SetDescription(s string) *Pet
	SetLifeSpan(i int) *Pet
	SetGeographicOrigin(s string) *Pet
	SetColor(s string) *Pet
	SetAge(i int) *Pet
	SetAgeEstimated(b bool) *Pet
	Build() (*Pet, error)
}

func NewPetBuilder() PetInterface {
	return &Pet{}
}

// SetSpecies sets the species for our pet, and returns a *Pet
func (p *Pet) SetSpecies(s string) *Pet {
	p.Species = s
	return p
}

// SetBreed sets the breed for our pet, and returns a *Pet
func (p *Pet) SetBreed(s string) *Pet {
	p.Breed = s
	return p
}

// SetMinWeight sets the minimum weight for our pet, and returns a *Pet
func (p *Pet) SetMinWeight(i int) *Pet {
	p.MinWeight = i
	return p
}

// SetMaxWeight sets the maximum weight for our pet, and returns a *Pet
func (p *Pet) SetMaxWeight(i int) *Pet {
	p.MaxWeight = i
	return p
}

// SetWeight sets the weight for our pet, and returns a *Pet
func (p *Pet) SetWeight(i int) *Pet {
	p.Weight = i
	return p
}

// SetDescription sets the description for our pet, and returns a *Pet
func (p *Pet) SetDescription(s string) *Pet {
	p.Description = s
	return p
}

// SetLifeSpan sets the lifespan for our pet, and returns a *Pet
func (p *Pet) SetLifeSpan(i int) *Pet {
	p.LifeSpan = i
	return p
}

// SetGeographicOrigin sets the geographic origin for our pet, and returns a *Pet
func (p *Pet) SetGeographicOrigin(s string) *Pet {
	p.GeographicOrigin = s
	return p
}

// SetColor sets the color for our pet, and returns a *Pet
func (p *Pet) SetColor(s string) *Pet {
	p.Color = s
	return p
}

// SetAge sets the age for our pet, and returns a *Pet
func (p *Pet) SetAge(i int) *Pet {
	p.Age = i
	return p
}

// SetAgeEstimated sets the estimated age for our pet, and returns a *Pet
func (p *Pet) SetAgeEstimated(b bool) *Pet {
	p.AgeEstimated = b
	return p
}

// Build uses the various "Set" functions above to build a pet, using the fluent
// interface. The inclusion of this function makes this an example of the
// Builder pattern.
func (p *Pet) Build() (*Pet, error) {
	if p.MinWeight > p.MaxWeight {
		return nil, errors.New("minimum weight must be less than maximum weight")
	}

	p.AverageWeight = (p.MinWeight + p.MaxWeight) / 2

	return p, nil
}
