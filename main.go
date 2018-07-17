package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	o := flag.String("o", "", "Write output to this file instead of stdout")
	flag.Parse()

	out, err := toYAML(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	w := os.Stdout
	if *o != "" {
		w, err = os.OpenFile(*o, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	fmt.Fprint(w, string(out))
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
