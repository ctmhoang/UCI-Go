# Go tools
* go doc
* go build
* go run
* go get
* go fmt
* go list
* go test
# Go notes
* Declaration
	* `var i typeName`
	* `var i = value`
	* `i := value` (declare in function not in global)
* Type alias `type aliasName typeName`;
* Has scope like JS (Closure)
* Conversion type has JS taste
* strconv pkg do String convertion to other primitive types and format numbers
* Has pointer
* Has range types of numeric (8 -> 64)
* Has complex number type
* Declare multiple const var
	* `const ( assign values with variables)`
* iota likes enum
	* syntax <br>
```
const (
	A type = iota // need to do on the top do not repeat
	B
	C
)
```
# Loop
* if statement mix of python vs curly braces
* for loop can have form of while loop in other program lang
* for loop can have no expression => inif loop
# Switch
* like old java switch with no break keyword
* Tagless switch like if elif else statement

_Break and Continue exist_

_Scan in fmt pkg take a pointer as an argument_

# Array
* Elements in array initialized to `zero value` like java
* Declare`var x [Quantity] Type`
* Array literal 
```
// len of literal have to be the same as the len of the array
var x [5]int = [5]{1,2,3,4,5}
// or we can use `...` key word to infer the size
x := [...]{1,2,3}
```
* Iterate like enumerate in python
```
for i, v range arrName {}
i -index
v -value
arrName - array Indentifier
```

# Slice
* Sub-Array in a underlying array
* Declare like sub list in python `x := arrName[strIdx:endIdx]`
* Can increase the size of the slice up to underlying array. If the size bigger => create new bigger array => time penalty
* Has three properties
	* Pointer - start of the array
	* Len - the number of elts in the slice `len()`
	* Cap - Maximun number of elts in the slice `cap()`
* 2 slice has the same ref to 1 underlying array having overlapped indexes can be collide eachother
* Slice literal `sli := []{1,2,3}`
* _Make slice_
```
// make(type, capacity)
sli := make([]int,10)
// make (type , len ,cap)
sli2 := make([]int , 1, 3)
```
* append() like append in python

