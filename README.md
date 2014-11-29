Turn file system events into JSON stream.

(Does not support recursive directory watching.)

# Install

```
go get github.com/hayeah/gowatch
```

# Usage

```
gowatch dir ...
```

It would ouput a JSON stream that you can pipe into another program for filtering & processing. Supposed we are watching a directory called `test`, you get output like:

```
> ./gowatch test
{"event":"chmod","path":"test/foobar"}
{"event":"remove","path":"test/foobar"}
{"event":"create","path":"test/foobar"}
{"event":"chmod","path":"test/foobar"}
```

Slice and dice as you please.

# Example

It is convenient to combine `gowatch` with a JSON processing tool like [jq](http://stedolan.github.io/jq/manual/).

Suppose we'd like to execute a command whenenver a new file is added, we can build a pipeline:

```
> ./gowatch test | unbuffer -p jq -M 'select(.event == "create")' | gogo 'echo process new file {{.path}}'
2014/11/29 11:35:04 watching: test
2014/11/29 11:35:06 run cmd: echo process new file test/foo
process new file test/foo
2014/11/29 11:35:06 run cmd: echo process new file test/bar
process new file test/bar
```

+ `unbuffer` - This [disables buffering](http://unix.stackexchange.com/questions/25372/turn-off-buffering-in-pipe) so jq sends output down the pipe as soon as possible.
  + `-p` is used in a pipeline to read input from upstream.
+ `jq` - It selects only the `create` events.
  + `-M` disables output color. This prevents JSON parser downstream from choking.
+ `gogo` - A utility to run processes for each object in a JSON stream. See: [hayeah/gogo](https://github.com/hayeah/gogo)

# File System Events Supported

See: [fsnotify.Op](https://godoc.org/gopkg.in/fsnotify.v1#Op)

```golang
const (
    Create Op  = 1 << iota
    Write
    Remove
    Rename
    Chmod
)
```