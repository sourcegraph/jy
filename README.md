# jy

The `jy` command reads JSON from stdin and writes YAML to stdout.

Convert YAML to JSON with [yj](https://github.com/sourcegraph/yj).

## Example

```sh
$ echo '{"hello": "world"}' | jy
hello: world
```

## Install

Install the latest with `go get -u github.com/sourcegraph/jy`.
