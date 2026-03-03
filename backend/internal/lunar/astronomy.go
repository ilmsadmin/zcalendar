package lunar

import "math"

const (
	PI = math.Pi
)

// JulianDayFromDate converts a calendar date to Julian Day Number
func JulianDayFromDate(day, month, year int) float64 {
	a := (14 - month) / 12
	y := float64(year) + 4800 - float64(a)
	m := float64(month) + 12*float64(a) - 3

	jd := float64(day) + math.Floor((153*m+2)/5) + 365*y +
		math.Floor(y/4) - math.Floor(y/100) + math.Floor(y/400) - 32045

	return jd
}

// DateFromJulianDay converts a Julian Day Number to calendar date
func DateFromJulianDay(jd float64) (day, month, year int) {
	z := int(jd + 0.5)
	a := z

	if z >= 2299161 {
		alpha := int((float64(z) - 1867216.25) / 36524.25)
		a = z + 1 + alpha - alpha/4
	}

	b := a + 1524
	c := int((float64(b) - 122.1) / 365.25)
	d := int(365.25 * float64(c))
	e := int(float64(b-d) / 30.6001)

	day = b - d - int(30.6001*float64(e))

	if e < 14 {
		month = e - 1
	} else {
		month = e - 13
	}

	if month > 2 {
		year = c - 4716
	} else {
		year = c - 4715
	}

	return day, month, year
}

// NewMoon calculates the Julian Day Number of the k-th new moon
// Based on Astronomical Algorithms by Jean Meeus
func NewMoon(k int) float64 {
	T := float64(k) / 1236.85
	T2 := T * T
	T3 := T2 * T

	// Mean new moon
	jd := 2415020.75933 + 29.53058868*float64(k) +
		0.0001178*T2 - 0.000000155*T3 +
		0.00033*math.Sin((166.56+132.87*T-0.009173*T2)*PI/180)

	// Sun's mean anomaly
	M := 359.2242 + 29.10535608*float64(k) - 0.0000333*T2 - 0.00000347*T3
	M = M * PI / 180

	// Moon's mean anomaly
	Mprime := 306.0253 + 385.81691806*float64(k) + 0.0107306*T2 + 0.00001236*T3
	Mprime = Mprime * PI / 180

	// Moon's argument of latitude
	F := 21.2964 + 390.67050284*float64(k) - 0.0016528*T2 - 0.00000239*T3
	F = F * PI / 180

	// Correction
	C1 := (0.1734-0.000393*T)*math.Sin(M) +
		0.0021*math.Sin(2*M) -
		0.4068*math.Sin(Mprime) +
		0.0161*math.Sin(2*Mprime) -
		0.0004*math.Sin(3*Mprime) +
		0.0104*math.Sin(2*F) -
		0.0051*math.Sin(M+Mprime) -
		0.0074*math.Sin(M-Mprime) +
		0.0004*math.Sin(2*F+M) -
		0.0004*math.Sin(2*F-M) -
		0.0006*math.Sin(2*F+Mprime) +
		0.0010*math.Sin(2*F-Mprime) +
		0.0005*math.Sin(M+2*Mprime)

	return jd + C1
}

// SunLongitude calculates the sun's ecliptic longitude at a given Julian Day
func SunLongitude(jd float64) float64 {
	T := (jd - 2451545.0) / 36525.0
	T2 := T * T

	// Geometric mean longitude of the Sun
	L0 := 280.46646 + 36000.76983*T + 0.0003032*T2
	L0 = math.Mod(L0, 360)
	if L0 < 0 {
		L0 += 360
	}

	// Mean anomaly of the Sun
	M := 357.52911 + 35999.05029*T - 0.0001537*T2
	Mrad := M * PI / 180

	// Equation of center
	C := (1.914602 - 0.004817*T - 0.000014*T2) * math.Sin(Mrad)
	C += (0.019993 - 0.000101*T) * math.Sin(2*Mrad)
	C += 0.000289 * math.Sin(3*Mrad)

	// Sun's true longitude
	theta := L0 + C

	// Apparent longitude (nutation correction)
	omega := 125.04 - 1934.136*T
	lambda := theta - 0.00569 - 0.00478*math.Sin(omega*PI/180)

	lambda = math.Mod(lambda, 360)
	if lambda < 0 {
		lambda += 360
	}

	return lambda
}

// getSunLongitudeAtJD returns the major solar term index at a given JD
// The ecliptic is divided into 12 equal parts of 30 degrees each
func getSolarTermIndex(sunLong float64) int {
	return int(math.Floor(sunLong / 30))
}
