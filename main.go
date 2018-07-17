package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	out, err := toYAML(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprint(os.Stdout, string(out))
}

func toYAML(r io.Reader) ([]byte, error) {
	ybuf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var y interface{}
	if err := json.Unmarshal(ybuf, &y); err != nil {
		return nil, err
	}
	j, err := yaml.Marshal(y)
	if err != nil {
		return nil, err
	}
	return j, nil
}
