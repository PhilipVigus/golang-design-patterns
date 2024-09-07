package singleton

// Singleton interface defines a method AddOne that returns an integer
type Singleton interface {
	AddOne() int
}

// singleton struct implements the Singleton interface and has a count field.
type singleton struct {
	count int
}

// instance is a package-level variable that holds the single instance of the singleton struct.
var instance *singleton

// GetInstance returns the single instance of the singleton struct.
// If the instance is nil, it initializes it.
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

// AddOne increments the count field of the singleton struct and returns the new value.
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
