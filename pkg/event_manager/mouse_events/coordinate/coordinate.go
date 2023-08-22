package coordinate

type EventCoordinate struct {
	X int16
	Y int16
}

type DoubleStateCoordinates struct {
	Start *EventCoordinate
	End   *EventCoordinate
}
