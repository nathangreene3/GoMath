# Mathematical Function Definitions and Dependencies

## Exp

Exp(x) = 1 + x + x^2/2 + x^3/6 + ... + x^n/n! for some sufficiently large n. If x = n*Ln2 + r for some integer n > 0 and r on range [0,1), then Exp(x) returns Pow2(n) * Exp(r). Ln2 is an exported, pre-computed value.

## Pow

Pow(x,y) = Exp(y*Ln(x)). If y is an integer, then PowInt(x,y) is returned. If y < 0, then 1/Pow(x,-y) is returned. If x < 0 and y does not decompose into an even root, then 