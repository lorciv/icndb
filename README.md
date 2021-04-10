# icndb 🦸‍♂️

icndb is a Go package that provides access to the Internet Chuck Norris Database (ICNDb) API. For information about the API, refer to [api.md](api.md).

The package exposes 4 functions: `Rand`, to get a random joke; `Randn` to get multiple random jokes; `Fetch` to get a specific joke, given its ID; and `Count`, to count all available jokes.

For example, to get a random joke:

```go
joke, err := icndb.Rand()
if err != nil {
    log.Fatal(err)
}
fmt.Println(joke.Text)
```

Most of the exposed functions return objects of type `Joke`. Joke represents the basic unit of information of the ICNDb API, and it exposes the following fields:

```go
type Joke struct {
	ID         int
	Text       string   // the actual joke
	Categories []string
}
```
Have fun! And remember: code runs faster when Chuck Norris watches it.
