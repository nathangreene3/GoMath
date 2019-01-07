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

	// Phi is the golden ratio (1+sqrt 5)/2.
	Phi = float64(1.61803398874989484820458683436563811772030917980576286213544862)

	// Pi is the ratio of a circle's circumfrence to its diameter.
	Pi = float64(3.14159265358979323846264338327950288419716939937510582097494459)

	// Ln2 is the natural logarithm of two (ln(2)).
	Ln2 = float64(0.693147180559945309417232121458176568075500134360255254120680009)

	// tol (1e-16) is the tolerance required by all numerical functions.
	// All numerical methods must converge to a value determined by this
	// tolerance. This does not guarentee the value returned is correct
	// (accurate) to the number of digits represented here, but that the
	// value returned cannot be made more precise with the method used.
	tol = float64(1e-16)
)
