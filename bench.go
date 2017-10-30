package run

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	c := os.Args
	fmt.Printf("%+v", c)
	Run(c[1])
}

func Run(c string) string {
	out, err := exec.Command(c).Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
