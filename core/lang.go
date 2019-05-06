package core

import "fmt"

/*
Lang shows programming languages
- Ext: extention(e.g., ".cpp")
- Compile: compile command (e.g., "g++ main.cpp")
	- you can use $SRC the target(main.cpp)
- Exec: execution command(e.g., "./a.out")
*/
type Lang struct {
	Ext     string `toml:"ext"`
	Compile string `toml:"compile"`
	Exec    string `toml:"exec"`
}

func (lang Lang) String() string {
	return fmt.Sprintf("[Ext:%v, Compile:'%v', Exec:'%v']", lang.Ext, lang.Compile, lang.Exec)
}
