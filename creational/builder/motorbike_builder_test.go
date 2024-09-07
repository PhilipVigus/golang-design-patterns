package builder

import "testing"

func TestMotorbikeBuilder(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	motorbikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(motorbikeBuilder)
	manufacturingComplex.Build()

	bike := motorbikeBuilder.Build()

	if bike.Wheels != 2 {
		t.Errorf("expecting 2 wheels, got %d", bike.Wheels)
	}

	if bike.Structure != Bike {
		t.Errorf("expecting Bike structure, got %s", bike.Structure)
	}

	if bike.Seats != 1 {
		t.Errorf("expecting 1 seats, got %d", bike.Seats)
	}
}
