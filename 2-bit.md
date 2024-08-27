**Bit Manipulation**

* Computationally fast 
* Computer work on Bits 

**Binary Numbers:**

> number representation in => 1 , 0

```
0 -> 0
1 -> 1
2 -> 10   
3 -> 11
4 -> 100
```


**Addition/Substraction**

Noraml 

Substraction 

9 - 3 = ? 

9 binary + 3 ( 2's complement ) = ans 

```
2's complement ? 
01 => 10 ( 1s complement ) invert bits 
    + 01 ( 2s complement ) add 0001 bit
```

**Bitwise Operators**
> & , | ,  ^ , >> , <<

`&` bitwise and 
```
0 0 0
0 1 0
1 0 0
1 1 1
```
----------------------------

`|` bitwise or 
```
0 0 0
0 1 1
1 0 1
1 1 1

```
----------------------------

`^` bitwise xor
```
0 0 0
0 1 1
1 0 1
1 1 0
```

----------------------------

`>>\<<` shift | left | right 

> left  2 => 010 | 001

> right 2 => 010 | 100

----------------------------

`~` Invert operator 

01 => 10 (1's complement)


---------------------------


Questions: 
Odd/Even 

```if  x % 2 == 0 even else odd```

Solve with Masking 

```if x & 1 == 0 even else odd```

----------------------------

Swap Number using bits 
```
a = a ^ b 
a = a ^ b 
b = a ^ b
```



