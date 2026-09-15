package timesignature

type TimeSignature struct {
	BeatsPerMeasure int
}

func (ts *TimeSignature) Toggle() {
	if ts.BeatsPerMeasure == 4 {
		ts.BeatsPerMeasure = 6
	} else {
		ts.BeatsPerMeasure = 4
	}
}
