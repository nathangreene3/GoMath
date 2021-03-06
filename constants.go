package math

// Note: accuracy != precision. A precise, wrong answer is not accurate.
// An accurate, imprecise answer is a lucky guess that may not be
// duplicated or may fail under other circumstances. Here, all values
// are precise to the constant tol. Accuracy is only guarenteed by the
// algorithm of choice.

const (
	// E is Euler's number, the natural rate of growth of the exponential
	// function.
	E = float64(2.71828182845904523536028747135266249775724709369995957496696763)

	// EulerMascheroni is the Euler-Mascheroni constant, usually denoted as the Greek letter gamma.
	EulerMascheroni = float64(0.57721566490153286060651209008240243104215933593992)

	// Phi is the golden ratio (1+sqrt 5)/2.
	Phi = float64(1.61803398874989484820458683436563811772030917980576286213544862)

	// Pi is the ratio of a circle's circumfrence to its diameter.
	Pi = float64(3.14159265358979323846264338327950288419716939937510582097494459)

	// Ln2 is the natural logarithm of two.
	Ln2 = float64(0.693147180559945309417232121458176568075500134360255254120680009)

	// Ln10 is the natural logarithm of ten.
	Ln10 = float64(2.30258509299404568401799145468436420760110148862877297603332790)

	// tol (1e-16) is the tolerance required by all numerical functions.
	// All numerical methods must converge to a value determined by this
	// tolerance. This does not guarentee the value returned is correct
	// (accurate) to the number of digits represented here, but that the
	// value returned cannot be made more precise with the method used.
	tol = float64(1e-16)

	maxFloat32             = 3.40282346638528859811704183484516925440e+38  // 2**127 * (2**24 - 1) / 2**23
	smallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45 // 1 / 2**(127 - 1 + 23)

	maxFloat64             = 1.797693134862315708145274237317043567981e+308 // 2**1023 * (2**53 - 1) / 2**52
	smallestNonzeroFloat64 = 4.940656458412465441765687928682213723651e-324 // 1 / 2**(1023 - 1 + 52)

	maxInt8   = 1<<7 - 1
	minInt8   = -1 << 7
	maxInt16  = 1<<15 - 1
	minInt16  = -1 << 15
	maxInt32  = 1<<31 - 1
	minInt32  = -1 << 31
	maxInt64  = 1<<63 - 1
	minInt64  = -1 << 63
	maxUint8  = 1<<8 - 1
	maxUint16 = 1<<16 - 1
	maxUint32 = 1<<32 - 1
	maxUint64 = 1<<64 - 1
)
