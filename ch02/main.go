package main

import (
	"fmt"
	"strings"
	"jvmgo/ch02/classpath"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
}
