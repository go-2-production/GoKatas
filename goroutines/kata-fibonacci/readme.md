# Goroutine Kata: Fibonacci Sequence Calculation

### Problem Description

Implement a function in Go that calculates the Fibonacci sequence up to a given number of terms using goroutines. The Fibonacci sequence is defined as follows:

- The first two terms are 0 and 1.
- Each subsequent term is the sum of the two preceding terms.

Your task is to create a goroutine for each Fibonacci term calculation and ensure they run concurrently. The final result should be printed once all goroutines have completed.

### Signature

```go
func CalculateFibonacciSequence(n int)
```

### Example

```go
CalculateFibonacciSequence(10)
```

Output:
```
0 1 1 2 3 5 8 13 21 34
```

### Constraints

- The input `n` will be a positive integer.
- The function should handle large values of `n` efficiently.

### Notes

- You may choose to use a channel to collect the results from each goroutine and print the final sequence once all calculations are complete.
- Make sure to handle any necessary synchronization to ensure all goroutines finish before printing the result.
