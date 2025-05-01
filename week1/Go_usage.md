# Week 1 - Go Basics Notes
## Pointer
- & (address-of operator) — gets the memory address of a variable.
- * (dereference operator) — gets the value at a memory address.

## Structs
- Composite data type that groups together zero or more fields (variables) under one name
- Template for an object 
- Use dot or pointer to access field

## Slices
- Array is fixed length, Slice is dynamic
- a[low, high], exclude last one 
- len VS capacity 
    - The length of a slice is the number of elements it contains.
    - The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
    - Capacity is recalculated as elements remaining from the new start to the end of the underlying array.
- Use 'make' to create dynamically sized array
- Use 'append' to add new elements
- Slice of slices

## Range
- return 1.index, 2.copy of element at that index

## Maps
- m := make(map[string]int)
- check if a key exist
    - elem, ok = m[key]
    - return boolean val

## Method and Interfaces
- A method is a function with a special receiver argument
    - func (r ReceiverType) MethodName(params) returnType [Method]
    - func Abs(v Vertex) float64 [Function]
- Pointer receiver
    - Methods with pointer receivers can modify the value to which the receiver points
- Interface

## Concurrency
- Goroutine
    - Lightweight thread
    - go f(x,y,z)
    - cheap, run independently
- Channels
    - a pipe through which goroutines can send and receive values of a specific type
    - Send: ch <- value
    - Receive: x := <-ch
    - buffered channels