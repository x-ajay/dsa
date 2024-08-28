**Recurssion**
steps
* find base case 
* find relation between problem and subproblem
* generalize the relation

```go
func hello(n) {
    if n == 0 {  // base case 
        return 
    }
    fmt.Println("Hello World") // operation
    hello(n-1)   // recurssion call 
}
```

---
Q: Find the sum of n number using recurssion

```go
func sum(n) int {
    if n == 1 {     // base case 
        return 1
    }
    // problem sum of n
    /*
    sum of 4 
        given number + sum of 3 ( dont think more on subproblem)
    */
    // generalize
    return n + sum(n-1) 
}

```

Q: Number of way to go from 0,0 to m,n in matrix

Approach: 
* if we have [0,0,0] matrix we have only one way to read end 
* if we have 
    ```
    [0]
    [0]
    [0]
    ```
    matrix then we have only one way to go 

* base case ```if m or n == 1 return 1```
* devide the problem into sub
    ```f( n,m-1) + f(n-1,m)```

> solution
```go
func numberOfWay(n, m int) int {
	if n == 1 || m == 1 {
		return 1
	}
	return numberOfWay(n-1, m) + numberOfWay(n, m-1)
}
```
---
Q: string palindrome by recurrsion

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
Q: Print superset of given string 

"abc" ->"","a","b","c","ab","ac","bc","abc"

```go

func findSuperSet(str string, i int, curr string) {
	if i == len(str) {
		println(curr)
		return
	}
	findSuperSet(str, i+1, curr)
	findSuperSet(str, i+1, fmt.Sprintf("%s%c", curr, str[i]))
}
```
---
Q: Find all possible permutataions:

> a,b => {a,b},{b,a}

> a,b,c => {a,b,c}{a,c,b}{b,a,c},{b,c,a},{c,a,b},{c,b,a}