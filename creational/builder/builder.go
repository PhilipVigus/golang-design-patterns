package builder

type VehicleType int

const (
	Car VehicleType = iota
	Bike
)

func (v VehicleType) String() string {
	return [...]string{"Car", "Bike"}[v]
}

// Vehicle is the product that will be built by the builders.
type Vehicle struct {
	Wheels    int         // Number of wheels on the vehicle
	Seats     int         // Number of seats in the vehicle
	Structure VehicleType // Type of structure (e.g., car, bike, bus)
}

// BuildProcess is the interface that defines the steps to build a vehicle.
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	Build() Vehicle
}

// ManufacturingDirector directs the building process using a BuildProcess.
type ManufacturingDirector struct {
	builder BuildProcess // The builder that will construct the vehicle
}

// Build constructs the vehicle by executing the build steps in order.
func (f *ManufacturingDirector) Build() {
	f.builder.SetWheels().SetSeats().SetStructure()
}

// SetBuilder sets the builder that the director will use to construct the vehicle.
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}
