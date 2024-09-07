package builder

// CarBuilder is a concrete builder that implements the BuildProcess interface to build a car.
type CarBuilder struct {
	v Vehicle // The vehicle that is being built
}

// SetWheels sets the number of wheels for the car and returns the builder.
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

// SetSeats sets the number of seats for the car and returns the builder.
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

// SetStructure sets the structure type for the car and returns the builder.
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = Car
	return c
}

// Build returns the final car (vehicle) that has been built.
func (c *CarBuilder) Build() Vehicle {
	return c.v
}
