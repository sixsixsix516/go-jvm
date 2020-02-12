package main

import "flag"
import "fmt"
import "os"

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

/**
 * 解析命令行参数
 * 将参数组装起来返回
 */
func parseCmd() *Cmd {
	cmd := &Cmd{}

	// 将参数绑定到结构体中,并设缺省值
	flag.BoolVar(&cmd.helpFlag, "help", false, "打印帮助信息")
	flag.BoolVar(&cmd.helpFlag, "?", false, "打印帮助信息")

	flag.BoolVar(&cmd.versionFlag, "version", false, "打印版本信息")

	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")

	// 解析参数
	flag.Parse()

	// 获取解析后的参数列表
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("%s 命令就没 %s 参数，你瞎写什么呢 ?\n",os.Args[0],os.Args[1:])
}
