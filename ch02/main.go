package main
import "fmt"

func main() {
	var cmd = parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("startJVM! classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}