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
# Break and Continue exist
# Scan in fmt pkg take a pointer as an argument
