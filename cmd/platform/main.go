package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	image = "micro/platform"
)

func main() {
	usage := fmt.Sprintf("%s {install|uninstall|update}", os.Args[0])

	switch os.Args[1] {
	case "install", "uninstall":
		usage = fmt.Sprintf("%s {install|uninstall} {dev|staging|platform}", os.Args[0])

		if len(os.Args) < 3 {
			fmt.Println(usage)
			return
		}

		// set the install/uninstall script
		action := os.Args[1] + ".sh"
		// create the args
		args := []string{action}
		args = append(args, os.Args[2:]...)
		// exec the command
		cmd := exec.Command("bash", args...)
		cmd.Dir = "./kubernetes"
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}
	case "update":
		usage = fmt.Sprintf("%s update [tag]", os.Args[0])

		if len(os.Args) < 3 {
			fmt.Println(usage)
			return
		}

		tag := os.Args[2]

		// set the tag for the micro deployment
		cmd := exec.Command("kubectl", "set", "image", "deployments", "micro="+image+":"+tag, "-l", "micro=runtime")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Println(usage)
		return
	}
}
