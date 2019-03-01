GoJSON is a command line utility to handle json in command line. 

### What it does

- [x] Retrieve nested objects
- [x] Pretty print JSON
- [x] Validate JSON
- [ ] Aggregate finct


## Installing

With go

```sh
$ go get -u github.com/sarathsp06/gojson
```

Or you may download the binary here [download](https://github.com/sarathsp06/gojson/tree/master/release) and use the binary as such


**Tip:**
> In unix move the binary to PATH


#### Key Syntax
* Key is a set of `.` seperated nested keys
* Can use 0-n numbers to refer to index in arrays
 
### Usage Examples

##### Getting a value 

* Get a string:
```sh
$ echo '{"name":{"first":"Sarath","last":"Pillai"}}' | gojson name.last
"Pillai"
```

* Get a block of JSON:
```sh
$ echo '{"name":{"first":"Sarath","last":"Pillai"}}'  | gojson name

{
  "first": "Sarath",
  "last": "Pillai"
}
```

* Try to get a non-existent key:
```sh
$ echo '{"name":{"first":"Sarath","last":"Pillai"}}' | gojson names
nil
```

Get an array value by index:
```sh
$ echo '{"people":[{"name":"saratha"},{"name":"syam"}]}' | gojson people.1.name                                               
"syam"
```
