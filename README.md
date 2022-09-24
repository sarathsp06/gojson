## GoJSON [![Build Status](https://travis-ci.org/sarathsp06/gojson.svg?branch=master)](https://travis-ci.org/sarathsp06/gojson)  [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=sarathsp06_gojson&metric=alert_status)](https://sonarcloud.io/dashboard?id=sarathsp06_gojson) [![Go Report Card](https://goreportcard.com/badge/github.com/sarathsp06/gojson)](https://goreportcard.com/report/github.com/sarathsp06/gojson) ![Last Commit](https://img.shields.io/github/last-commit/sarathsp06/gojson.svg)
GoJSON is a command line utility to handle json in command line. 

### What it does

- [x] Retrieve nested objects
- [x] Pretty print JSON
- [x] Validate JSON
- [x] Aggregate functions


## Installing

**Go Dev version**

```sh
go install  github.com/sarathsp06/gojson@latest
```

**Binray Release**

[download](https://github.com/sarathsp06/gojson/releases) and use the binary as such for your platform


**Tip:**
> In unix move the binary to PATH

#### Key Syntax

- Key is a set of `.` seperated nested values
- Can use 0-n numbers to refer to index in arrays
- Can use `lower:upper` syntax to refer to a range of an array. Eg: players.1:3 
- Can use keys of inner objects directly on arrays or range of them. Eg:  players.name where players is an array

### Usage Examples

##### Getting a value 

- Get a string:

```sh
$ echo '{"name":{"first":"Sarath","last":"Pillai"}}' | gojson name.last
"Pillai"
```

- Get a block of JSON:

```sh
$ echo '{"name":{"first":"Sarath","last":"Pillai"}}'  | gojson name

{
  "first": "Sarath",
  "last": "Pillai"
}
```

- Try to get a non-existent key:

```sh
$ echo '{"name":{"first":"Sarath","last":"Pillai"}}' | gojson names
nil

```

- Get an array value by index:

```sh
$ echo '{"people":[{"name":"saratha"},{"name":"syam"}]}' | gojson people.1.name
"syam"
```

- Projection from a slice

```sh
$ echo '{"people":[{"name":"saratha"},{"name":"syam"},{"name":"singh"},{"name":"ping"}]}' | gojson people.2:.name 
[
  "singh",
  "ping"
]
```

- Slice of array

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

- Handling JSON key names with a `.`:

```sh
$ echo '{"first.name":"Sarath","last.name":"Pillai"}' | gojson \"first.name\"
"Sarath"
```
