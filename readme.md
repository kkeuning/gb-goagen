# gb-goagen
A gb plugin to run the goagen tool in a gb project.

The goagen generator is used to working with a $GOPATH.

## Installation
```
go get github.com/kkeuning/gb-goagen/...
```

## Pre-reqs
```
go get github.com/constabulary/gb/...
go get github.com/goadesign/goa/goagen
```

This is based on Doug Clark's gb-run plug-in, but tailored for running goagen specifically.

```
Usage:
  gb goagen [command] [arguments]

Run gb-goagen from the root of your gb project.
Specify the output directory with -o.

Example:
	gb goagen bootstrap -d goa-adder/design -o $PWD/src/goa-adder

Runs as if you had executed the following:
	GOPATH=$GOPATH:$PWD goagen bootstrap -d goa-adder/design -o $PWD/src/goa-adder
	
Gorma example:

	gb goagen gen --design=congo/design --pkg-path=github.com/goadesign/gorma -o $PWD/src/congo

Getting help:

	Same as gb -h, shows gb help, NOT goagen help:
		gb goagen -h

	To see this help together with goagen help:
		gb goagen

	To see just the goagen help:
		goagen
		goagen -h
```

Important!

Current behavior of the plugin requires the gogen command to be specified before its arguments:
```
gb goagen [command] [arguments]
```

This will NOT work currently:
```
gb goagen --design=congo/design gen --pkg-path=github.com/goadesign/gorma -o $PWD/src/congo`
```

But this equivalent works fine:
```
gb goagen gen --design=congo/design --pkg-path=github.com/goadesign/gorma -o $PWD/src/congo
```


