Turn file system events into stream that you can process with script more easily.

Does not support recursive directory watching.

# Install

```
go install github.com/hayeah/gowatch
```

# Usage

```
gowatch dir ...
```

It would print a stream that you can pipe into another program for processing. The first word is the event type, and following it (separated by a space) is the relative path name.

```
chmod gowatch.go
create .sublca9.tmp
chmod .sublca9.tmp
chmod README.md
remove README.md
create README.md
create .sublaa1.tmp
chmod .sublaa1.tmp
```

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