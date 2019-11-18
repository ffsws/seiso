package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
)

func init() {
	rootCmd.AddCommand(gitCmd)
}

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Print the git HEAD",
	Long:  `tbd`,
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		worktree := osfs.New(args[0])
		gitdir, err := worktree.Chroot(".git")
		if err != nil {
			panic(err)
		}

		storer := filesystem.NewStorage(gitdir, nil)
		repository, err := git.Open(storer, worktree)
		if err != nil {
			panic(err)
		}

		fmt.Println(repository.Head())
	},
}