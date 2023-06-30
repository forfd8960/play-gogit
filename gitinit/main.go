package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	repoPath := flag.String("gitpath", ".", "the git repo path")
	flag.Parse()
	if *repoPath == "" {
		fmt.Println("empty repo path")
		os.Exit(1)
	}

	repo, err := git.PlainInit(*repoPath, false)
	if err != nil {
		fmt.Println("git init err: ", err)
		os.Exit(1)
	}

	fmt.Printf("Git init repo: %+v\n", *repo)
}
