package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

const MSG = "🦫"

func main() {
	commitAmount := flag.Int("amount", 100, "The number of commits")
	withPush := flag.Bool("with-push", false, "push to repo at the very end")
	flag.Parse()

	count := 0

	if _, err := os.Stat(".git/"); os.IsNotExist(err) {
		if err := exec.Command("git", "init").Run(); err != nil {
			panic("failed to init")
		}
	}

	for i := 0; i < *commitAmount; i++ {
		println(commitAmount)
		if err := exec.Command("git", "commit", "--allow-empty", "-m", MSG).Run(); err != nil {
			panic(err)
		} else {
			count++
			progressPercent := float32(count) / float32(*commitAmount) * 100
			fmt.Printf("\rprogress : %.2f%%", progressPercent)

		}
	}

	if count > 0 {
		fmt.Printf("\r %v         \n number of commits : %d", "Done Buddy 🦫", count)
	}

	if *withPush {
		fmt.Println("\n\npushing ...")

		if err := exec.Command("git", "push").Run(); err != nil {
			fmt.Printf("%s\n\n", "failed to push! please check origin.")
		}
	}

}
