package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {
	repoPath := flag.String("repo", ".", "the git repo path")
	flag.Parse()
	if *repoPath == "" {
		fmt.Println("empty repo path")
		os.Exit(1)
	}

	repo, err := git.PlainOpen(*repoPath)
	if err != nil {
		fmt.Println("open repo err: ", err)
		os.Exit(1)
	}

	wt, err := repo.Worktree()
	if err != nil {
		fmt.Println("open worktree err: ", err)
		os.Exit(1)
	}

	fmt.Printf("Opened work tree: %+v\n", wt)

	filename := filepath.Join(*repoPath, "example-git-file")
	if err = os.WriteFile(filename, []byte("hello world!"), 0644); err != nil {
		fmt.Println("write file err: ", err)
		os.Exit(1)
	}

	hash, err := wt.Add(filename)
	if err != nil {
		fmt.Printf("Add file: %s err: %v\n", filename, err)
		os.Exit(1)
	}

	fmt.Printf("Add file: %s, hash: %s\n", filename, hash)

	commitHash, err := wt.Commit("first commit", &git.CommitOptions{
		All: true,
		Author: &object.Signature{
			Name:  "forld",
			Email: "forld@gmail.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		fmt.Printf("Commit file: %s err: %v\n", filename, err)
		os.Exit(1)
	}

	fmt.Printf("Commit file: %s, hash: %s\n", filename, commitHash)
}
