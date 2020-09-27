package scan

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func RunScan() {
	cmd := exec.Command("handle", "java")
	output, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		cmd.Run()
	}()
	scanner := bufio.NewScanner(output)
	// 创建文件
	fileName := "result.txt"
	f, err3 := os.Create(fileName) //创建文件
	if err3 != nil {
		fmt.Println("创建文件失败")
	}
	w := bufio.NewWriter(f) //创建新的 Writer 对象

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if strings.Contains(line, "java") {
			result := fmt.Sprintf("Java进程Pid为%s<====>路径为%s\n", fields[2], fields[len(fields)-1])
			Write3(result, w)
		}
	}
	w.Flush()
	f.Close()
}

func Write3(result string, w *bufio.Writer) {
	_, err4 := w.WriteString(result)
	if err4 != nil {
		fmt.Println(err4)
	}
}
