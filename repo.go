package main

import (
	"fmt"
	"io"
	"os"

	memfs "gopkg.in/src-d/go-billy.v2/memfs"
	git "gopkg.in/src-d/go-git.v4"
	memory "gopkg.in/src-d/go-git.v4/storage/memory"
)

var fs = memfs.New()
var s = memory.NewStorage()

func CloneRepo() {

	//fs = memfs.New()

	repo, err := git.Clone(s, fs, &git.CloneOptions{
		URL:      "https://github.com/mickeyyawn/yawn.me",
		Progress: os.Stdout,
	})

	fmt.Println("just cloned...")
	HandleError(err)

}

func SyncRepo_bak() {

	//fs = memfs.New()

	repo, err := git.Clone(s, fs, &git.CloneOptions{
		URL:      "https://github.com/mickeyyawn/yawn.me",
		Progress: os.Stdout,
	})

	fmt.Println("just cloned...")
	HandleError(err)

	if err.Error() == "repository already exists" {
		//err = repo.Fetch(&git.FetchOptions{
		//RemoteName: "origin",

		//repo, err := git.PlainOpen(path)

		//HandleError(err)
		fmt.Println("let us go pull latest...............")

		//err = repo.Fetch(&git.FetchOptions{
		//	RemoteName: "",
		//	Progress:   os.Stdout,
		//})

		err = repo.Pull(&git.PullOptions{
			RemoteName: "",
		})

	} //)
	//HandleError(err)
	//fmt.Println("just fetched...")
	//}

	theFile, err := fs.Open("README.md")

	io.Copy(os.Stdout, theFile)

	// Prints the content of the CHANGELOG file from the cloned repository
	//changelog, _ := fs.Open("CHANGELOG")

	//io.Copy(os.Stdout, changelog)

}

func SyncRepoToDisk() {

	path := os.TempDir() + "/testing" // <- need uuid there

	fmt.Println(os.TempDir()) //. os.TempDir()+"/blah"

	repo, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      "https://github.com/mickeyyawn/yawn.me",
		Progress: os.Stdout,
	})

	if err != nil {
		if err.Error() == "repository already exists" {
			fmt.Println("oh noes, the repo was already on disk...")

			repo, err := git.PlainOpen(path)

			//HandleError(err)
			fmt.Println("let us go pull latest...............")

			//err = repo.Fetch(&git.FetchOptions{
			//	RemoteName: "",
			//	Progress:   os.Stdout,
			//})

			err = repo.Pull(&git.PullOptions{
				RemoteName: "",
			})

			fmt.Println(err.Error())
			//HandleError(err)
		}
	} else {
		err = repo.Pull(&git.PullOptions{
			RemoteName: "origin",
		})
		HandleError(err)
	}

	/*

		if err != nil {
			if err.Error() == "repository already exists" {
				fmt.Println("oh noes, the repo was already on disk...")

				repo, err := git.PlainOpen(os.TempDir() + "/blah")
				HandleError(err)
				fmt.Println("let us go pull latest...")

				err = repo.Fetch(&git.FetchOptions{
					RemoteName: "origin",
				})
				fmt.Println(err.Error())
				//HandleError(err)
			}
		} else {
			err = repo.Pull(&git.PullOptions{
				RemoteName: "origin",
			})
			HandleError(err)
		}

	*/

}
