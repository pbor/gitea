// +build ignore

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	version = "0.7.0-beta0"
	sha     = rev()
)

var steps = map[string]step{
	"deps":    executeDeps,
	"lint":    executeLint,
	"fmt":     executeFmt,
	"vet":     executeVet,
	"test":    executeTest,
	"install": executeInstall,
	"build":   executeBuild,
	"bindata": executeBindata,
	"clean":   executeClean,
}

func init() {
	os.Setenv("GO15VENDOREXPERIMENT", "1")
}

func main() {
	for _, arg := range os.Args[1:] {
		step, ok := steps[arg]

		if !ok {
			fmt.Println("Error: Invalid step", arg)
			os.Exit(1)
		}

		err := step()

		if err != nil {
			fmt.Println("Error: Failed step", arg)
			os.Exit(1)
		}
	}
}

type step func() error

func executeDeps() error {
	deps := []string{
		"github.com/jteeuwen/go-bindata/...",
		"github.com/Masterminds/glide",
	}

	for _, dep := range deps {
		err := run(
			"go",
			"get",
			"-u",
			dep)

		if err != nil {
			return err
		}
	}

	return run(
		"glide",
		"install")
}

func executeLint() error {
	err := run(
		"go",
		"get",
		"github.com/golang/lint")

	if err != nil {
		return err
	}

	return run(
		"golint",
		"./...")
}

func executeFmt() error {
	return run(
		"go",
		"fmt",
		"./...")
}

func executeVet() error {
	return run(
		"go",
		"vet",
		"./...")
}

func executeTest() error {
	withTags := getTags()

	ldf := fmt.Sprintf(
		"-X main.revision=%s -X main.version=%s",
		sha,
		version)

	if len(withTags) > 0 {
		// Ned that seperate because of escaping
		tags := fmt.Sprintf(
			"%s",
			strings.Join(withTags, " "))

		return run(
			"go",
			"test",
			"-tags",
			tags,
			"-ldflags",
			ldf,
			"./models/...",
			"./modules/...",
			"./routers/...")
	} else {
		return run(
			"go",
			"test",
			"-cover",
			"-ldflags",
			ldf,
			"./models/...",
			"./modules/...",
			"./routers/...")
	}
}

func executeInstall() error {
	withTags := getTags()

	ldf := fmt.Sprintf(
		"-X main.revision=%s -X main.version=%s",
		sha,
		version)

	if len(withTags) > 0 {
		// Ned that seperate because of escaping
		tags := fmt.Sprintf(
			"%s",
			strings.Join(withTags, " "))

		return run(
			"go",
			"install",
			"-v",
			"-tags",
			tags,
			"-ldflags",
			ldf,
			"github.com/go-gitea/gitea")
	} else {
		return run(
			"go",
			"install",
			"-v",
			"-ldflags",
			ldf,
			"github.com/go-gitea/gitea")
	}
}

func executeBuild() error {
	withTags := getTags()

	ldf := fmt.Sprintf(
		"-X main.revision=%s -X main.version=%s",
		sha,
		version)

	if len(withTags) > 0 {
		// Ned that seperate because of escaping
		tags := fmt.Sprintf(
			"%s",
			strings.Join(withTags, " "))

		return run(
			"go",
			"build",
			"-v",
			"-tags",
			tags,
			"-ldflags",
			ldf,
			"github.com/go-gitea/gitea")
	} else {
		return run(
			"go",
			"build",
			"-v",
			"-ldflags",
			ldf,
			"github.com/go-gitea/gitea")
	}
}

func executeBindata() error {
	var paths = []struct {
		input  string
		output string
		pkg    string
	}{
		{"conf/...", "modules/bindata/bindata.go", "bindata"},
	}

	for _, path := range paths {
		err := run(
			"go-bindata",
			fmt.Sprintf("-o=%s", path.output),
			"-ignore=\"README\\\\.md\"",
			fmt.Sprintf("-pkg=%s", path.pkg),
			path.input)

		if err != nil {
			return err
		}
	}

	return nil
}

func executeClean() error {
	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		suffixes := []string{
			".out",
		}

		for _, suffix := range suffixes {
			if strings.HasSuffix(path, suffix) {
				if err := os.Remove(path); err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	files := []string{
		"gitea",
		"gitea.exe",
	}

	for _, file := range files {
		if _, err := os.Stat(file); err != nil {
			continue
		}

		if err := os.Remove(file); err != nil {
			return err
		}
	}

	return nil
}

func hasTag(require string) bool {
	for _, tag := range getTags() {
		if tag == require {
			return true
		}
	}

	return false
}

func getTags() []string {
	if len(os.Getenv("TAGS")) != 0 {
		return strings.Split(os.Getenv("TAGS"), ",")
	} else {
		return make([]string, 0)
	}
}

func run(command string, args ...string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	trace(cmd.Args)
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}

func rev() string {
	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	raw, err := cmd.CombinedOutput()

	if err != nil {
		return "HEAD"
	}

	return strings.Trim(string(raw), "\n")
}

func trace(args []string) {
	print("+ ")
	println(strings.Join(args, " "))
}
