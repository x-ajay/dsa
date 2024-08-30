# Bit Manipulation Documentation

## Overview

Bit manipulation is a fundamental concept in computer science, providing computationally efficient methods for handling data at the bit level. Since computers operate on binary data, bit manipulation techniques can be utilized for a variety of tasks such as performing arithmetic operations, logical operations, and more.

## Binary Numbers

Binary numbers represent data using only two digits: `0` and `1`. Every binary number corresponds to a unique decimal value.

**Examples:**

```
Decimal  -> Binary
0        -> 0
1        -> 1
2        -> 10   
3        -> 11
4        -> 100
```

## Addition and Subtraction in Binary

### Normal Subtraction

For subtraction, binary arithmetic utilizes the concept of 2's complement to simplify the operation. 

**Example:**

Subtracting 3 from 9 (`9 - 3`):

1. Convert 3 to its 2's complement.
2. Add the 2's complement of 3 to 9 in binary.

### Calculating 2's Complement

To find the 2's complement of a binary number:

1. **1's Complement**: Invert all bits.
   - Example: `01` becomes `10`
2. **2's Complement**: Add `1` to the 1's complement.
   - Example: `10 + 01` results in `11`

## Bitwise Operators

Bitwise operators allow direct manipulation of bits within binary numbers. The most common operators are `&`, `|`, `^`, `<<`, `>>`, and `~`.

### `&` (Bitwise AND)

Performs a logical AND operation on each corresponding bit of two numbers.

```
0 & 0 -> 0
0 & 1 -> 0
1 & 0 -> 0
1 & 1 -> 1
```

### `|` (Bitwise OR)

Performs a logical OR operation on each corresponding bit of two numbers.

```
0 | 0 -> 0
0 | 1 -> 1
1 | 0 -> 1
1 | 1 -> 1
```

### `^` (Bitwise XOR)

Performs a logical XOR (exclusive OR) operation on each corresponding bit of two numbers.

```
0 ^ 0 -> 0
0 ^ 1 -> 1
1 ^ 0 -> 1
1 ^ 1 -> 0
```

### `<<` and `>>` (Bitwise Shift Left/Right)

Shifts the bits of a binary number to the left or right.

- **Left Shift (`<<`)**: Moves bits to the left by the specified number of positions.
  - Example: `010 << 2` results in `001000`
- **Right Shift (`>>`)**: Moves bits to the right by the specified number of positions.
  - Example: `010 >> 2` results in `0001`

### `~` (Bitwise NOT / Invert Operator)

Inverts all the bits of the binary number (also known as the 1's complement).

- Example: `01` becomes `10`

## Practical Applications

### Determining Odd or Even

To check if a number is odd or even using bit manipulation, use the bitwise AND operator (`&`).

**Traditional Method:**

```python
if x % 2 == 0:
    print("Even")
else:
    print("Odd")
```

**Bitwise Method:**

```python
if x & 1 == 0:
    print("Even")
else:
    print("Odd")
```

### Swapping Numbers Using Bit Manipulation

You can swap two numbers without using a temporary variable by using the XOR (`^`) operator.

**Steps:**

```python
a = a ^ b
b = a ^ b
a = a ^ b
```

After these operations, `a` and `b` will have swapped values.

## Conclusion

Bit manipulation is a powerful tool for efficient computation and optimization in software development. Understanding and utilizing bitwise operators can significantly improve the performance of low-level code, making them essential knowledge for any programmer.