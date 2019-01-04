# GoMath

```golang
go get git.biscorp.local/serverdev/math
```

The *GoMath* package is a library of mathematical functions to replace the standard *GoLang math package.* Not all functions have been duplicated and not all functions share the same naming convention used in the GoLang math package. There are several custom functions available for specific scenarios in BIS applications.

*Use these functions at your own risk.* Testing is being performed to ensure the values returned are precise, accurate, and robust. However, this library is still *considered experimental* and is not intended to replace all GoLang math package functionality. Most values are accurate to +/-5e-14, but testing has not yet been completed. The goal is to meet +/-5e-15.

## Constants

The following exported constants are commonly used and are provided for quick access.

- **E** = 2.718... is Euler's number, the natural rate of growth of the exponential function.
- **Phi** = 1.618... is the golden ratio.
- **Pi** = 3.141... is the ratio of a circle's circumfrence to its diameter.
- **Ln2** = 0.693... is the natural logarithm of two.

## Functions

### Exponential, Power, and Logarithmic Functions

- **PowInt** returns x^n for any real x and any integer n. Panics if x and n are zero simultaneously.
  - (0,0) --> *panic*: 0^0 is undefined. **Note:** The GoLang math package and other languages define this case to be 1, because.
  - (0,3) --> 0 = 0^3
  - (2,0) --> 1 = 2^0
  - (2,3) --> 8 = 2^3
  - (-2,3) --> -8 = (-2)^3
  - (2,-3) --> 0.125 = 2^-3 = 1/2^3 = 1/8
- **Pow2** returns 2^n for any integer n.
- **Pow10** returns 10^n for any integer n.
- **Pow** returns x^y for real x and y.
- **Exp** returns e^x.
- **Sqrt** returns +x^0.5.
- **NthRoot** returns x^(1/n).
- **Ln** returns the natural logarithm of base e.
- **Log** returns the base b-logarithm of x.
- **Log10** returns the base-10 logarithm of x.
- **Log2** returns the base-2 logarithm of x.

### Number Property Functions

- **GCD** returns the greatest common divisor of two non-negative integers.
  - (a,b) = (2,4) --> 2 (2 and 4 share 2 as a divisor)
  - (a,b) = (3,4) --> 1 (3 and 4 are relatively prime)
- **LCM** returns the least common multiple of two integers.
  - (a,b) = (2,4) --> 4 (4 is a multiple of 2)
  - (a,b) = (3,4) --> 12 (12 is the smallest multiple of both 3 and 4)
- **Factorial** returns n! Panics if n < 0.
  - n = 0 --> 0! = 1
  - n = 1 --> 1! = 1
  - n = 2 --> 2! = 2
  - n = 3 --> 3! = 2 * 3 = 6
  - n = 4 --> 4! = 2 * 3 * 4 = 24
- **IsPrime** reports whether a number is prime or not. Panics if n is less than two.
  - n = 2 --> true
  - n = 4 --> false
- **Factor** returns a collection of each divisor of a positive integer mapped to the number of times each divisor divides said integer. The integer n passed is returned as a factor if and only if n is prime.
  - n = 12 = 2^2 * 3^1 --> [1:1, 2:2, 3:1] (12 is not prime)
  - n = 13 = 13^1 --> [1:1, 13:1] (13 is prime)

### Rounding Functions

- **Abs** returns |x|.
  - x = -1 --> 1
  - x = 1 --> 1
- **Floor** returns the largest integer value less than x as a float64. Currently, x < 0 is not supported.
  - x = 1.0 --> 1.0
  - x = 1.1 --> 1.0
  - x = 1.9 --> 1.0
- **Ceiling** returns the smallest integer greater than x as a float64. If x is an integer value, then x is returned, not x+1. Currently, x < 0 is not supported.
  - x = 1.0 --> 1.0
  - x = 1.1 --> 2.0
  - x = 1.9 --> 2.0
- **Round** returns x rounded to the nearest whole number as a float64.
- **RoundTo** rounds a number to the nth decimal place. For example, x = 0.123 becomes x.12 when n = 2. Special cases are n = 0 and n < 0. If n = 0, then 1 is returned. If n < 0, then 0 is returned.
- **RoundUpToBase** rounds a number up to the next multiple of b. Panics if base is negative.
  - x = -1, b = 5 --> *panic* (TODO: Define on negative ranges.)
  - x = 0, b = 5 --> 5
  - x = 1, b = 5 --> 5
  - x = 4, b = 5 --> 5
- **RoundToMag10** rounds an integer x to the the next 10^n. Panics if n is negative.
  - x = 1234, n = 0 -->  1235 (rounds to 10^0 =     1)
  - x = 1234, n = 1 -->  1240 (rounds to 10^1 =    10)
  - x = 1234, n = 3 -->  2000 (rounds to 10^3 =  1000)
  - x = 1234, n = 5 --> 10000 (rounds to 10^5 = 10000)
- **OrderMag10** returns the order of magnitude (largest power) n of 10 such that x >= 10^n. Panics if n is not positive.
  - n = 0 --> *panic* (10^n > 0 for all n.)
  - n = 1234 --> 3 (1234 >= 1000 = 10^3)
