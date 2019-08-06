package main

import "flag"
import "fmt"
import "os"

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	XjreOption  string // X表示extension
	class       string
	args        []string
}

func parseCmd() *Cmd {
	var cmd = &Cmd{} // 定义一个空的Cmd, 然后取其的地址, 赋值给cmd这个指针
	flag.Usage = printUsage  // 如果执行exe没输入什么东西的话, 就会冒出usage提示

	// flag指的是-options
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.versionFlag, "v", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	// 这是真正输入的参数 比如 ./ch03.exe java.lang.Object 中 args[0] = "java.lang.Object"
	var args = flag.Args() // read args from input
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
