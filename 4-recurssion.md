# Recursion Algorithms Documentation

## Overview: What is Recursion?
**Recursion** is a problem-solving technique where a function calls itself to solve smaller instances of the same problem. The key steps in recursion include:

1. **Identify the Base Case**: The condition under which the recursion stops.
2. **Establish the Relationship**: Find the connection between the problem and its subproblems.
3. **Generalize the Relationship**: Express the solution in a recursive formula.

### Example: Basic Recursion

```go
func hello(n int) {
    if n == 0 {  // Base Case
        return 
    }
    fmt.Println("Hello World") // Operation
    hello(n-1)   // Recursive Call
}
```

---

### Q: Find the Sum of N Numbers Using Recursion

**Problem**: Calculate the sum of the first `n` natural numbers using recursion.

**Solution**:

- **Base Case**: If `n` equals 1, return 1.
- **Relation**: The sum of `n` is the sum of `n-1` plus `n`.
- **Generalized**: `sum(n) = n + sum(n-1)`.

```go
func sum(n int) int {
    if n == 1 { // Base Case 
        return 1
    }
    // Recursive Case
    return n + sum(n-1) 
}
```

---

### Q: Number of Ways to Move from (0,0) to (m,n) in a Matrix

**Problem**: Given a matrix of size `m x n`, find the number of ways to move from the top-left corner to the bottom-right corner, moving only down or right.

**Approach**:

- **Base Case**: If either `m` or `n` is 1, there's only one way to move.
- **Divide the Problem**: The number of ways to reach `(m,n)` is the sum of ways to reach `(m-1,n)` and `(m,n-1)`.

```go
func numberOfWay(n, m int) int {
    if n == 1 || m == 1 {
        return 1
    }
    return numberOfWay(n-1, m) + numberOfWay(n, m-1)
}
```

---

### Q: Check if a String is a Palindrome Using Recursion

**Problem**: Determine if a given string is a palindrome using recursion.

**Solution**:

- **Base Case**: If the left index `l` is greater than or equal to the right index `r`, the string is a palindrome.
- **Check Characters**: Compare characters at `l` and `r`. If they match, move inward and recurse.

```go
func isPalindrome(str string, l, r int) bool {
    if l >= r {
        return true
    }
    if str[l] != str[r] {
        return false
    }
    return isPalindrome(str, l+1, r-1)
}
```

---

### Q: Print Superset of a Given String

**Problem**: Generate all possible subsets (the power set) of a given string.

**Solution**:

- **Recursive Exploration**: For each character, decide whether to include it in the current subset or not, and recurse.

```go
func findSuperSet(str string, i int, curr string) {
    if i == len(str) {
        println(curr)
        return
    }
    findSuperSet(str, i+1, curr)                       // Exclude current character
    findSuperSet(str, i+1, fmt.Sprintf("%s%c", curr, str[i]))  // Include current character
}
```

Example: For the string `"abc"`, the output would be `""`, `"a"`, `"b"`, `"c"`, `"ab"`, `"ac"`, `"bc"`, `"abc"`.

---

### Q: Find All Possible Permutations of a String

**Problem**: Generate all permutations of a given string.

**Solution**:

- **Recursive Swapping**: Recursively swap each character with every other character and generate permutations.

```go
func permute(str string, l int, r int) {
    if l == r {
        fmt.Println(str)
    } else {
        for i := l; i <= r; i++ {
            str = swap(str, l, i)
            permute(str, l+1, r)
            str = swap(str, l, i) // Backtrack
        }
    }
}

func swap(str string, i, j int) string {
    runes := []rune(str)
    runes[i], runes[j] = runes[j], runes[i]
    return string(runes)
}
```

Example: For the string `"abc"`, the permutations would be `"abc"`, `"acb"`, `"bac"`, `"bca"`, `"cab"`, `"cba"`.

---

This documentation provides a comprehensive guide to solving problems using recursion. The examples illustrate key recursion concepts and how to implement them effectively in Go.