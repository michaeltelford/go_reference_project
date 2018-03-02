
# Go Lang Reference Project

A go-to example Go-lang project referencing some main points of the language, all in one place for convenience.

This reference project is a personal one, whilst I've tried to ensure that the practices within are standard language practices, I am not a Go expert.

Please view both the details below and the actual source code itself as needed. Enjoy!

## Prerequisites

The reference assumes:

- `go` is installed correctly
- `$GOPATH` is correctly setup and configured (with `src`, `bin` and `pkg` within)
- A basic working knowledge of `*nix` and `git` shell commands

## Technical

### Commands

| Example Command       | Effect                                                       |
| --------------------- | :----------------------------------------------------------- |
| `go build main.go`    | Builds a binary or lib.                                      |
| `go test ./...`       | Runs any tests that are found. Searches `pwd` recursively.   |
| `go install main.go`  | Builds and installs binary and libs locally to `$GOPATH`.    |
| `go run main.go`      | Runs the file or binary.                                     |
| `go get package_name` | Retrieves the software lib from the source and installs it <br />locally to `$GOPATH`. See 'Software Packages' below. |

### Code Standards

| Tool     | Effect                                                       |
| -------- | ------------------------------------------------------------ |
| `gofmt`  | Built in tool which formats your code. <br /><br />Built into many editors e.g. Atom's go-plus package runs it on each save. |
| `golint` | Lints your code. <br /><br />GitHub: https://github.com/golang/lint<br />Install: `go get -u golang.org/x/lint/golint`<br /> |

### Documentation

A lot of useful general information for Go can be found at:

https://golang.org/doc/

The Go standard lib docs can be found at:

https://golang.org/pkg/

### Project Workspace/Layout

Every Go project should have a basic directory and file structure in place, a common workspace, an example one might look like:

- `project_dir` (run `git init` here)
- - `main.go` (Start of execution, can be called anything but must have `func main()` defined; an example main file can be seen below)
  - `src` (optional dir can be called anything but 'src' makes sense)
    - `lib_dir` (this is the name of your lib)
      - `lib_file` (file can be called anything but `package` within must be `lib_dir` - the containing directory name - doesn't work without this!)
      - `another_lib_file` (again `package` must be `lib_dir` to be part of the same lib)
      - ...

### Software Packages

Go doesn't have a dedicated software repository like RubyGems.org for example. Instead, GitHub.com is typically used. Each lib or app is a repo on a developer's or company's GitHub account. 

#### Installation

The `go get` command is used to install a 3rd party lib/package locally e.g. :

```
	$ go get github.com/justinas/alice
```

This will install the above package in `$GOPATH/src/github.com/justinas/alice`. 

#### Importing

Go uses absolute import paths for projects, hence the need for a `$GOPATH`. The import lookup starts in the `$GOPATH/src` directory. The above installed package can be imported to your project using:

```go
import (
    // Std lib import
	"net/http"

    // 3rd party package import following a `go get ...`
	"github.com/justinas/alice"

    // Developer lib import, path is always absolute e.g. $GOPATH/src/...
	"github.com/michaeltelford/go_reference_project/src/middleware/logger"
)
```

You should note that because the import paths are absolute, the `main.go` file can be placed anywhere on your file system and then moved without the import paths needing updated. 

#### Dependancy Management

Go has a (soon to be) official dependancy management tool called Go Dep. There are a couple other 3rd party tools out there but this one is from the Go core team and should probably be used over anything else. 

##### Basic Usage

See the docs below for a full description on using `dep` but in a nutshell:

Install dep on a Mac with:

```
	$ brew install dep
```

To initialise a new repo/project, run:

```
	$ dep init
```

You should commit the generated files and directory to source.

Once initialised, keep your deps up to date with:

```
	$ dep ensure
```

Remember to run `dep ensure` before you commit new changes to keep your deps in sync. That's about it. 

##### GitHub

https://github.com/golang/dep

##### Docs

https://golang.github.io/dep

### Example Main

Every project must have an execution starting point which is a file with `package main` and `func main() {}` defined in it. 

The example main file is a working example of a web server using Go's `net/http` std lib. It has 2 endpoints and logs all requests and responses via middleware. The chaining of middleware to app is done using the 3rd party `alice` lib. The logging middleware source code is the 'Example Lib' snippet below. 

> src/main.go

```go
package main

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/michaeltelford/go_reference_project/src/middleware/logger"
)

func main() {
	chain := alice.New(logger.New).Then(newApp())
	http.ListenAndServe(`:8080`, chain)
}

func newApp() http.Handler {
	app := http.NewServeMux()

	app.HandleFunc(`/healthcheck`, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	app.HandleFunc(`/hello`, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`Hello world`))
	})
	app.Handle(`/`, http.NotFoundHandler())

	return app
}
```

### Example Lib

Most projects will create some form of lib or package for use elsewhere. Packages should be self contained and reusable where possible. 

Below is a working example of logging middleware (used in the `main.go` file above). Once installed to your `$GOPATH` it can be imported to any project. 

> src/middleware/logger/logger.go

```go
package logger

import (
	"log"
	"net/http"
	"os"
)

type logger struct {
	Next http.Handler
	Log  *log.Logger
}

func (l logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.Log.Printf("%s: %s - Request received", r.RemoteAddr, r.URL)
	l.Next.ServeHTTP(w, r)
	l.Log.Printf("%s: %s - Response sent", r.RemoteAddr, r.URL)
}

// New is a constructor function for logging middleware
func New(h http.Handler) http.Handler {
	return logger{
		Next: h,
		Log:  log.New(os.Stdout, ``, 1),
	}
}
```

Some things to note:

- **Constructor & Interface** - `New` is a constructor function which returns an instance of the `http.Handler` interface (which all middleware components adhere to). 
- **Struct** - The `logger` type is a struct containing the next piece of middleware to be called and a logger instance. 
- **Methods** - The `ServeHTTP` func is a method because it has a receiver of `logger` (identified by `l`). `New` is a function (not a method) because it has no receiver. 
- **Exports** - Go exports anything that is capitalised, otherwise its private, accessible only within its package. In this case, only `New` is exported which is typical of constructor functions. 

### Testing

Below are some commonly used testing libs for Go. All test files should be located in the same directory as the file they test, not a dedicated `test` directory like some languages. The test file name should also end in `*_test.go`. 

#### go test

Go provides a default `testing` package in the std lib for unit testing your code. Below is an excellent guide to getting started (with examples). 

https://smartystreets.com/blog/2015/02/go-testing-part-1-vanillla

##### Example

Below is a basic example of unit testing using Go's built in `testing` package. It's taken from the guide above. 

> vanilla_test.go

```go
package vanilla

import "testing"

func TestGutterBalls(t *testing.T) {
	t.Log("Rolling all gutter balls... (expected score: 0)")
	
    game := NewGame()
	game.rollMany(20, 0)

	if score := game.Score(); score != 0 {
		t.Errorf("Expected score of 0, but it was %d instead.", score)
	}
}
```

#### goconvey

GoConvey is a 3rd party BDD testing tool with a rich array of assertions built in. It also provides a useful web UI for displaying test results (and running them automatically on file change). The package is also backward compatible with `go test` meaning you can use its UI separate from its testing lib. 

##### Github

https://github.com/smartystreets/goconvey

##### Installation

```
	$ go get github.com/smartystreets/goconvey
```

##### Example

Below is an example which tests the logger middleware used above. `cd` into the repo root directory and run `goconvey` to see the UI with test results. 

> src/middleware/logger/logger_test.go

```go
package logger

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given I have a mock HTTP app", t, func() {
		app := http.NewServeMux()

		Convey("When I call logger.New", func() {
			middleware := New(app)

			Convey("Then a configured logger middleware is returned", func() {
				So(middleware, ShouldHaveSameTypeAs, logger{})
			})
		})
	})
}

func TestServeHTTP(t *testing.T) {
	Convey("Given I have a mock HTTP app", t, func() {
		app := http.NewServeMux()
		app.HandleFunc(`/hello`, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Write([]byte(`Hello world`))
		})

		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/hello", nil)

		var buf bytes.Buffer
		expectedLogOutput := `2018/03/01 : /hello - Request received
2018/03/01 : /hello - Response sent
`

		Convey("And I have a mock logger middleware", func() {
			middleware := logger{
				Next: app,
				Log:  log.New(&buf, ``, 1),
			}

			Convey("When I call ServeHTTP", func() {
				middleware.ServeHTTP(w, r)

				Convey("Then a response status of 200 is returned", func() {
					So(w.Code, ShouldEqual, 200)
				})
				Convey("And a response body of `Hello world` is returned", func() {
					So(w.Body.String(), ShouldEqual, `Hello world`)
				})
				Convey("And the request and response has been logged", func() {
					So(string(buf.Bytes()), ShouldEqual, expectedLogOutput)
				})
			})
		})
	})
}
```

### Debugging

There are several methods to help with debugging an application. Here are a couple of examples. 

#### watcher

The 3rd party watcher lib prevents you from having to restart your app with each new change by doing it for you on the saving of a watched file. To use it, install it and simply `cd` into the directory containing the `main.go` file of your application and run the `watcher` command. Watcher will automatically restart your app with any changes. This is very useful during development and debugging and requires no configuration. Just install and run. 

##### Github

https://github.com/canthefason/go-watcher

#### Print Statements

It might be old skool but sometimes print statements to `STDOUT` are enough to debug the logic in your application. You can do so by using Go's `fmt` or `log` std lib packages.

```go
import "fmt"

fmt.Println(someVar)
```

#### CLI Debugger

If you're not using an IDE but need a breakpoint or two then consider using Go Delve. It allows you to debug an application by setting breakpoints via the CLI. When stopped on a breakpoint you can type `help` for the list of commands. You can continue, step to the next line, step into a function, print and/or set a variable during runtime and lots of other cool debugging features. 

##### Github

https://github.com/derekparker/delve

##### Basic Usage

```sh
$ dlv debug
Type 'help' for list of commands.
(dlv) break main.main
Breakpoint 1 set at 0x12e8428 for main.main() ./main.go:10
(dlv) continue
> main.main() ./main.go:10 (hits goroutine(1):1 total:1) (PC: 0x12e8428)
     5:	
     6:		"github.com/justinas/alice"
     7:		"github.com/michaeltelford/go_reference_project/src/middleware/logger"
     8:	)
     9:	
=>  10:	func main() {
    11:		chain := alice.New(logger.New).Then(newApp())
    12:		http.ListenAndServe(`:8080`, chain)
    13:	}
    14:	
    15:	func newApp() http.Handler {
```

#### IDE Debuggers

There are probably a few Go specific IDE's out there so DuckDuckGo (search engine) is your friend here. However, one I've used is GoLand. It has the ability to set breakpoints and debug inside the IDE. 

##### Website

https://www.jetbrains.com/go

#### Text Editor Debuggers

There are also more than a few text editors out there that support the Go language via packages etc. So again, just check if a package exists for your current fav editor. I personally like to use Atom and its go-plus package which does some cool stuff straight out of the box. For example, it builds and tests your code on each file save. It also has the ability to debug using breakpoints but I haven't used this much. 

##### Website

https://atom.io/packages/go-plus

## Thanks

Thanks for looking, if you have any suggestions on improving the reference or new additions then feel free to raise an issue. Also, if you'd like to have your own reference project then feel free to use this as a starter and fork it. 
