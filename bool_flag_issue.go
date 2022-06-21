package main

import "flag"

type NotBool struct{}

func (f *NotBool) IsBoolFlag() bool {
	return false
}

func (f *NotBool) Set(v string) error {
	return nil
}
func (f *NotBool) String() string {
	return ""
}

func main() {
	value := NotBool{}
	flag.Var(&value, "test", "test usage, line above should be -test value")
	flag.PrintDefaults()
}
