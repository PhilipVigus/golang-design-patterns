package builder

import "testing"

func TestCarBuilder(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Build()

	car := carBuilder.Build()

	if car.Wheels != 4 {
		t.Errorf("expecting 4 wheels, got %d", car.Wheels)
	}

	if car.Structure != Car {
		t.Errorf("expecting Car structure, got %s", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("expecting 5 seats, got %d", car.Seats)
	}
}
