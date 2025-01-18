# Go Array Index Out of Bounds Error

This repository demonstrates a common error in Go: an index out of bounds error when working with arrays.  The `bug.go` file contains code that attempts to access an array element beyond its defined length. The `bugSolution.go` file provides a corrected version.

## Bug Description

Go arrays have a fixed size.  Attempting to access an element using an index greater than or equal to the array's length will result in a runtime panic (error). This is because the program tries to access memory that is not allocated to the array.

## Solution

The solution involves carefully checking array indices before accessing elements.  Ensure that the loop iterates within the valid range of 0 to len(a)-1.