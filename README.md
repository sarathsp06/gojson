## GoJSON [![Build Status](https://travis-ci.org/sarathsp06/gojson.svg?branch=master)](https://travis-ci.org/sarathsp06/gojson)  [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=sarathsp06_gojson&metric=alert_status)](https://sonarcloud.io/dashboard?id=sarathsp06_gojson) [![Go Report Card](https://goreportcard.com/badge/github.com/sarathsp06/gojson)](https://goreportcard.com/report/github.com/sarathsp06/gojson)

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

Projection from a slice
```sh
$ echo '{"people":[{"name":"saratha"},{"name":"syam"},{"name":"singh"},{"name":"ping"}]}' | gojson people.2:.name 
[
  "singh",
  "ping"
]

```

Slice of array

```sh
$ echo '{"people":[{"name":"saratha"},{"name":"syam"},{"name":"singh"},{"name":"ping"}]}' | gojson people.2:5     
[
  {
    "name": "singh"
  },
  {
    "name": "ping"
  }
]
```