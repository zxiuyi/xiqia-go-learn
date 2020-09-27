package tomcat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

func KillTomcat(port string) {
	cmd := exec.Command(`netstat`, `-aon`)
	output, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		cmd.Run()
	}()
	processOutPut(output, port)
}

func processOutPut(output io.ReadCloser, port string) {
	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		line := scanner.Text()
		if hasTomcatPort(line, port) {
			parts := strings.Fields(line)
			tomcatPid := parts[len(parts)-1]
			log.Println(tomcatPid)
			killCmd := exec.Command(`taskkill`, `/F`, `/PID`, tomcatPid)
			err := killCmd.Run()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("关闭端口[%s]成功\n", port)
			}
		}
	}

}

func hasTomcatPort(line string, port string) bool {
	return strings.Contains(line, `LISTENING`) && (strings.Contains(line, port) || strings.Contains(line, port))
}
