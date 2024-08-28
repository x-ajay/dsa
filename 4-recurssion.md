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
