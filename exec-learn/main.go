package main

func main() {
	// 只执行，不管返回结果
	//cmd := exec.Command("ls", "-l", "$HOME")
	//err := cmd.Run()
	//if err != nil {
	//	log.Fatalf("cmd.Run() failed with %s\n", err)
	//}

	// 设置当前执行环境的环境变量
	// 使用环境变量的值
	//os.Setenv("NAME", "DALAOHEI")
	//cmd := exec.Command("ls", "-l", os.ExpandEnv("$HOME"))
	//err := cmd.Run()
	//if err != nil {
	//	log.Fatalf("cmd.Run() failed with %s\n", err)
	//}

	// 那到执行的结果
	//cmd := exec.Command("ls", "-l", "/Users/xmly")
	//out, err := cmd.CombinedOutput()
	//if err != nil {
	//	fmt.Printf("combined out:\n%s\n", string(out))
	//	log.Fatalf("cmd.Run() failed with %s\n", err)
	//}
	//fmt.Printf("combined out:\n%s\n", string(out))

}
