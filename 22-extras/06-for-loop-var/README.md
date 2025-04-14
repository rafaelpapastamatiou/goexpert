# For Loop Variable Capture Bug in Go Prior to 1.22

Prior to Go 1.22, there was a common programming mistake related to variable capture in for loops that could lead to subtle bugs. The issue was:

When using a loop variable in a closure or goroutine inside the loop, all iterations would capture the **same variable** (by reference) rather than capturing a unique value for each iteration.

## Example of the Problem

```go
func main() {
    var funcs []func()
    
    for i := 0; i < 5; i++ {
        funcs = append(funcs, func() {
            fmt.Println(i)  // Captures reference to loop variable 'i'
        })
    }
    
    // After loop completes, 'i' equals 5
    for _, f := range funcs {
        f()  // In Go < 1.22, this would print "5" five times
    }
}
```

## The Root Cause

The root cause was that in Go < 1.22:
- Only one variable `i` was created for the entire loop
- Each closure captured a reference to that same variable
- By the time the closures executed, the variable's value had changed to the final value

## The Fix in Go 1.22

Go 1.22 changed this behavior so that each loop iteration creates a new variable, ensuring that each closure captures the value it expects:

- Each iteration now has its own variable scope
- This makes the code behave as most developers intuitively expect
- The example above would now print 0, 1, 2, 3, 4 in Go 1.22 and later

For compatibility with existing code that relied on the old behavior, Go provides a way to explicitly request the old behavior using the `go:loopvar` directive.