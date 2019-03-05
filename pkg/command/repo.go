package command

import "fmt"

func CreateRepo(repo string) (string, error) {
	fmt.Printf("Creating repo: %s...", repo)
	return repo, nil
}
