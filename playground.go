package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, _ := exec.Command("docker", "run", "-it", "-d", "harbor.aibee.cn/store-config/aibee_beijing_showroom:20210209_205323_test_qa_87_224_16F_v510").Output()

	fmt.Printf("out is: %s", out)
}
