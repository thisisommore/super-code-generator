package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/go-git/go-git/v5"
)

func TestGenerateProject(t *testing.T) {

	tempDirStr := t.TempDir()
	remoteUrlWant := "https://github.com/thisisommore/thisreponotexist"
	GenerateProject(tempDirStr, remoteUrlWant)
	README_ME_FILE, e := os.ReadFile(tempDirStr + "/README.md")
	if e != nil {
		t.Fatal(e)
	}
	strRead := string(README_ME_FILE)

	if res := strings.Contains(strRead, "# Backend application template"); !res {
		t.Fatal("Clone validation failed")
	}

	repo, _ := git.PlainOpen(tempDirStr)

	remote, _ := repo.Remote("origin")

	if remoteUrl := remote.Config().URLs[0]; remoteUrl != remoteUrlWant {
		t.Fatal("Expected remote url not found")
	}
}
