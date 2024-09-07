package builder

// BikeBuilder is a concrete builder that implements the BuildProcess interface to build a bike.
type BikeBuilder struct {
	v Vehicle // The vehicle that is being built
}

// SetWheels sets the number of wheels for the bike and returns the builder.
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

// SetSeats sets the number of seats for the bike and returns the builder.
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 1
	return b
}

// SetStructure sets the structure type for the bike and returns the builder.
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = Bike
	return b
}

// Build returns the final bike (vehicle) that has been built.
func (b *BikeBuilder) Build() Vehicle {
	return b.v
}
