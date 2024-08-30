# Mathematics and Number Theory Documentation

## Overview: Fundamental Concepts in Number Theory

This documentation covers essential concepts and problems in mathematics and number theory, including factorials, trailing zeros, palindromic numbers, the Sieve of Eratosthenes, and the Greatest Common Divisor (GCD).

---

### Q: Factorial of a Number

**Problem**: Calculate the factorial of a given number `n`. The factorial of `n` is the product of all positive integers less than or equal to `n`.

**How it works**:

- **Formula**: \( n! = n \times (n-1)! \)
- **Base Cases**: 
  - \( 0! = 1 \)
  - \( 1! = 1 \)

```go
func factorial(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    return n * factorial(n-1)
}
```

---

### Q: Trailing Zeros in a Factorial

**Problem**: Determine the number of trailing zeros in the factorial of a given number `n`. Trailing zeros are created by factors of 10, which is \( 2 \times 5 \). The number of trailing zeros is determined by the number of times `5` is a factor in the numbers from `1` to `n`.

**Hint**:

- Count how many times `5` is a factor in `n!`.
- Remember that numbers like `25`, `125`, etc., contribute more than one `5`.

**Formula**:
\[
\text{Number of trailing zeros} = \left\lfloor \frac{n}{5} \right\rfloor + \left\lfloor \frac{n}{25} \right\rfloor + \left\lfloor \frac{n}{125} \right\rfloor + \ldots
\]

> **Solution**:
```go
func trailingZeros(n int) int {
    res := 0
    for i := 5; i <= n; i *= 5 {
        res += n / i
    }
    return res
}
```

---

### Q: Palindrome Number

**Problem**: Check if a given number is a palindrome. A number is a palindrome if it reads the same forward and backward.

> **Solution**:
```go
func palindromeNumber(num int) bool {
    original := num
    pal := 0
    for num > 0 {
        last := num % 10
        pal = pal*10 + last
        num = num / 10
    }
    return pal == original
}
```

---

### Q: Sieve of Eratosthenes

**Problem**: Find all prime numbers up to a given number `n` using the Sieve of Eratosthenes. This algorithm efficiently identifies prime numbers by iteratively marking the multiples of each prime number starting from `2`.

> **Solution**:
```go
func sieveOfEratosthenes(n int) {
    primes := make([]bool, n+1)
    for i := 2; i <= n; i++ {
        primes[i] = true
    }
    for i := 2; i*i <= n; i++ {
        if primes[i] {
            for j := i * i; j <= n; j += i {
                primes[j] = false
            }
        }
    }
    for i := 2; i <= n; i++ {
        if primes[i] {
            fmt.Println(i)
        }
    }
}
```

---

### Q: Greatest Common Divisor (GCD)

**Problem**: Compute the GCD of two numbers `a` and `b`. The GCD is the largest number that divides both `a` and `b` without leaving a remainder.

**Euclidean Algorithm**:
\[
\text{gcd}(a, b) = \text{gcd}(b, a \% b)
\]
Continue until \( b = 0 \).

> **Solution**:
```go
func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a % b)
}
```

---

### Q: Modular Arithmetic (Introduction)

**Modular Arithmetic** involves performing arithmetic operations where numbers "wrap around" after reaching a certain value, known as the modulus. This concept is essential in number theory, particularly in cryptography and solving problems related to divisibility.

**Basic Operations**:
- **Addition**: \( (a + b) \mod m \)
- **Subtraction**: \( (a - b) \mod m \)
- **Multiplication**: \( (a \times b) \mod m \)
- **Exponentiation**: \( a^b \mod m \)

**Example Problem**: Compute the result of \( 7^4 \mod 5 \).

```go
func modExp(base, exp, mod int) int {
    res := 1
    for exp > 0 {
        if exp%2 == 1 {
            res = (res * base) % mod
        }
        base = (base * base) % mod
        exp /= 2
    }
    return res
}
```

---

This documentation provides an overview of essential topics in mathematics and number theory, with practical examples and solutions implemented in Go. These concepts form the foundation for solving various mathematical and algorithmic problems.