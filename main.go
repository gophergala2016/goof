package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"sync"

	"github.com/fatih/color"
)

type gitDirectory struct {
	path string
	root string
}

type gitUpdate struct {
	path   string
	result bool
}

var (
	list   = flag.Bool("list", true, "List .git directories found in your $HOME directory.")
	update = flag.Bool("update", false, "Update all .git directories found in your $HOME directory.")
	done   = make(chan struct{})

	// "colors" (colors, colors...)
	cyan  = color.New(color.FgCyan).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

func main() {
	flag.Parse()

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	dirs := getGitDirectories(usr.HomeDir)

	if *update && *list {
		*list = false
	}

	if *list {
		printGitDirectories(done, usr.HomeDir, dirs)
	} else {
		updateGitDirectories(done, dirs)
	}

	<-done
}

func getGitDirectories(dir string) <-chan gitDirectory {
	var wg sync.WaitGroup
	out := make(chan gitDirectory)

	go func() {
		wg.Wait()
		close(out)
	}()

	wg.Add(1)
	go readDir(dir, &wg, out)

	return out
}

func readDir(dir string, wg *sync.WaitGroup, out chan<- gitDirectory) {
	defer wg.Done()

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			if entry.Name() == ".git" {
				d := gitDirectory{
					root: dir,
					path: dir + entry.Name(),
				}
				out <- d
			} else {
				wg.Add(1)
				subdir := filepath.Join(dir, entry.Name())
				go readDir(subdir, wg, out)
			}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents
var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
		// acquire token
	}
	// release token
	defer func() { <-sema }()

	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "goof: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "goof: %v\n", err)
	}
	return entries
}

func listGitDirectories(dirs <-chan gitDirectory) <-chan []string {
	out := make(chan []string)
	go func() {
		for d := range dirs {
			s := []string{filepath.Base(d.root), "[" + d.path + "]"}
			out <- s
		}
		close(out)
	}()
	return out
}

func printGitDirectories(done chan struct{}, dir string, dirs <-chan gitDirectory) {

	defer close(done)

	fmt.Println("Git repositories")
	for s := range listGitDirectories(dirs) {
		fmt.Printf("%s - %s\n", cyan(s[0]), green(s[1]))
	}
}

func updateGitDirectories(done chan struct{}, dirs <-chan gitDirectory) {
	defer close(done)

	var status string
	var fn func(a ...interface{}) string

	for s := range updateGitDir(dirs) {
		if s.result {
			status = "updated"
			fn = green
		} else {
			status = "not updated"
			fn = red
		}
		fmt.Printf("[%s] - %s\n", cyan(s.path), fn(status))
	}
}

func updateGitDir(dirs <-chan gitDirectory) <-chan gitUpdate {

	var wg sync.WaitGroup
	out := make(chan gitUpdate)

	update := func(d gitDirectory) {
		wg.Add(1)
		go func(path string) {
			u := gitUpdate{path: path, result: true}

			os.Chdir(path)
			cmd := exec.Command("git", "pull")
			err := cmd.Start()
			if err != nil {
				u.result = false
				out <- u
			}
			err = cmd.Wait()
			if err != nil {
				u.result = false
				out <- u
			} else {
				out <- u
			}
			wg.Done()
		}(d.root)
	}

	for d := range dirs {
		go update(d)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
