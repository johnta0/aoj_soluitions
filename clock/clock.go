package clock

const (
	oneHourAngle = 30
)

// CalcFromShortHand returns hour and minute in  integer
// calculated from given angle of the short hand
// of the clock
func CalcFromShortHand(ang int) (int, int) {
	h := ang / oneHourAngle

	rem := ang - oneHourAngle * h
	m := 2 * rem

	return h, m
}
