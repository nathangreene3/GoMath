package math

import (
	"math"
	"testing"

	"github.com/cheynewallace/tabby"
)

// TODO: Add panic test scenarios

// func TestRoundingScenario(t *testing.T) {
// 	for i := 90; i < 121; i++ {
// 		fmt.Println(i, RoundToMag10(i, OrderMag10(i)-1))
// 	}
// }

// func TestFactorial(t *testing.T) {
// 	data := []struct {
// 		input int
// 		ans   int
// 	}{
// 		{0, 1},
// 		{1, 1},
// 		{2, 2},
// 		{3, 6},
// 		{4, 24},
// 		{5, 120},
// 		{6, 720},
// 		{7, 5040},
// 		{8, 40320},
// 		{9, 362880},
// 		{10, 3628800},
// 	}
// 	var actual int

// 	for i := range data {
// 		actual = Factorial(data[i].input)
// 		if actual != data[i].ans {
// 			table := tabby.New()
// 			table.AddHeader("Function", "Input", "Value")
// 			table.AddLine("exp", data[i].input, actual)
// 			table.AddLine("", "", data[i].ans)
// 			table.Print()
// 			t.Fatal("\n")
// 		}
// 	}
// 	fmt.Println("Factorial passed")
// }

// func TestPowInt(t *testing.T) {
// 	var x float64
// 	var n int
// 	var actual float64
// 	var expected float64
// 	var err float64

// 	for i := 0; i < 1000; i++ {
// 		x, n = rand.ExpFloat64(), rand.Int() // Cant test negative x using math.Pow
// 		actual = PowInt(x, n)
// 		expected = math.Pow(x, float64(n))
// 		err = Abs(expected - actual)
// 		if 10*tol < err {
// 			table := tabby.New()
// 			table.AddHeader("Function", "Input", "Value")
// 			table.AddLine("PowInt", x, actual)
// 			table.AddLine("math.Pow", x, expected)
// 			table.AddLine("Abs Error", "", err)
// 			table.Print()
// 			t.Fatal("\n")
// 		}
// 	}
// 	fmt.Println("PowInt passed")
// }

// func TestPow(t *testing.T) {
// 	var x, y float64
// 	var actual float64
// 	var expected float64
// 	var err float64

// 	for i := 0; i < 1000; i++ {
// 		x, y = rand.ExpFloat64(), rand.ExpFloat64()
// 		actual = Pow(x, y)
// 		expected = math.Pow(x, y)
// 		err = Abs(expected - actual)
// 		if 10*tol < err {
// 			table := tabby.New()
// 			table.AddHeader("Function", "Input x", "Input y", "Value")
// 			table.AddLine("Pow", x, y, actual)
// 			table.AddLine("math.Pow", x, y, expected)
// 			table.AddLine("Abs Error", "", "", err)
// 			table.Print()
// 			t.Fatal("\n")
// 		}
// 	}
// 	fmt.Println("Pow passed")
// }

// func TestExp(t *testing.T) {
// 	var x float64
// 	var actual float64
// 	var expected float64
// 	var err float64

// 	for i := 0; i < 1000; i++ {
// 		x = math.Pow(-1, float64(rand.Intn(2)+1)) * rand.ExpFloat64()
// 		actual = Exp(x)
// 		expected = math.Exp(x)
// 		err = Abs(expected - actual)
// 		if 10*tol < err {
// 			table := tabby.New()
// 			table.AddHeader("Function", "Input", "Value")
// 			table.AddLine("Exp", x, actual)
// 			table.AddLine("math.Exp", x, expected)
// 			table.AddLine("Abs Error", "", err)
// 			table.Print()
// 			t.Fatal("\n")
// 		}
// 	}
// 	fmt.Println("Exp passed")
// }

func TestLn(t *testing.T) {
	var x float64
	var actual float64
	var expected float64
	var err float64

	for x = 0.1; x < 2; x += 0.1 {
		actual = Ln(x)
		expected = math.Log(x)
		err = Abs(expected - actual)
		table := tabby.New()
		table.AddHeader("Function", "Input", "Value")
		table.AddLine("Ln", x, actual)
		table.AddLine("math.Log", x, expected)
		table.AddLine("Abs Error", "", err)
		table.Print()
	}

	// for i := 0; i < 1000; i++ {
	// 	x = rand.ExpFloat64()
	// 	actual = Ln(x)
	// 	expected = math.Log(x)
	// 	err = Abs(expected - actual)
	// 	if 10*tol < err {
	// 		table := tabby.New()
	// 		table.AddHeader("Function", "Input", "Value")
	// 		table.AddLine("Ln", x, actual)
	// 		table.AddLine("math.Log", x, expected)
	// 		table.AddLine("Abs Error", "", err)
	// 		table.Print()
	// 		t.Fatal("\n")
	// 	}
	// }
	// fmt.Println("Ln passed")
}

// func TestSqrt(t *testing.T) {
// 	var x float64
// 	var actual float64
// 	var expected float64
// 	var err float64

// 	for i := 0; i < 1000; i++ {
// 		x = rand.ExpFloat64()
// 		actual = Sqrt(x)
// 		expected = math.Sqrt(x)
// 		err = Abs(expected - actual)
// 		if 10*tol < err {
// 			table := tabby.New()
// 			table.AddHeader("Function", "Input", "Value")
// 			table.AddLine("Sqrt", x, actual)
// 			table.AddLine("math.Sqrt", x, expected)
// 			table.AddLine("Abs Error", "", err)
// 			table.Print()
// 			t.Fatal("\n")
// 		}
// 	}
// 	fmt.Println("Sqrt passed")
// }

// func TestNthRoot(t *testing.T) {
// 	var x float64
// 	var n int
// 	var actual float64
// 	var expected float64
// 	var err float64

// 	for i := 0; i < 1000; i++ {
// 		x = rand.ExpFloat64()
// 		n = rand.Int() + 1
// 		actual = NthRoot(x, n)
// 		expected = math.Pow(x, 1/float64(n))
// 		err = Abs(expected - actual)
// 		if 10*tol < err {
// 			table := tabby.New()
// 			table.AddHeader("Function", "Input x", "Input n", "Value", "Note")
// 			table.AddLine("NthRoot", x, n, actual, "")
// 			table.AddLine("math.Pow", x, n, expected, "x^(1/n)")
// 			table.AddLine("Abs Error", "", err, "")
// 			table.Print()
// 			t.Fatal("\n")
// 		}
// 	}
// 	fmt.Println("NthRoot passed")
// }

// func TestNumDigits(t *testing.T) {
// 	// var n int
// 	// var b int
// 	// var actual int
// 	// var expected int
// 	// var err float64

// 	// for i := 0; i < 1000; i++ {
// 	// 	x = rand.ExpFloat64()
// 	// 	n = rand.Int() + 1
// 	// 	actual = NthRoot(x, n)
// 	// 	expected = math.Pow(x, 1/float64(n))
// 	// 	err = Abs(expected - actual)
// 	// 	if 10*tol < err {
// 	// 		table := tabby.New()
// 	// 		table.AddHeader("Function", "Input", "Value", "Note")
// 	// 		table.AddLine("NthRoot", x, actual, "")
// 	// 		table.AddLine("math.Pow", x, expected, "x^(1/n)")
// 	// 		table.AddLine("Abs Error", "", err, "")
// 	// 		table.Print()
// 	// 		t.Fatal("\n")
// 	// 	}
// 	// }
// }

// func BenchmarkMathPow(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		math.Pow(100, 100)
// 	}
// }

// func BenchmarkPowInt(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		PowInt(100, 100)
// 	}
// }

// func BenchmarkMathPow10(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		math.Pow10(100)
// 	}
// }

// func BenchmarkPow(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Pow(100, 100)
// 	}
// }

// func BenchmarkPow2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Pow2(100)
// 	}
// }

// func BenchmarkPow10(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Pow10(100)
// 	}
// }

// func BenchmarkLn(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Ln(100)
// 	}
// }

// func BenchmarkMathLn(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		math.Log(100)
// 	}
// }

func TestPowSpec(t *testing.T) {
	// x := 0.1991534529041898 // Approx 0.2 and ln(0.2) = -ln(5) = -1.6...
	// y := 0.5625736705708702
	// n := int(y)
	// fmt.Println(Ln(x) - math.Log(x))
	// fmt.Println(Exp((y-float64(n))*Ln(x)) - math.Exp((y-float64(n))*math.Log(x)))
}
