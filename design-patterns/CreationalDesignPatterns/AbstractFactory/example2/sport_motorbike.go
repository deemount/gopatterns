package main

// SportMotorbike is a concrete product
type SportMotorbike struct{}

func (s *SportMotorbike) NumWheels() int {
	return 2
}
func (s *SportMotorbike) NumSeats() int {
	return 1
}
func (s *SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}
