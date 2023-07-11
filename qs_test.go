package main

import (
	"strings"
	"testing"

	"github.com/atotto/clipboard"
	"github.com/stretchr/testify/assert"
	"github.com/untillpro/qs/git"
)

func TestDeleteDup(t *testing.T) {
	str := deleteDupMinus("13427-Show--must----go---on")
	assert.Equal(t, str, "13427-Show-must-go-on")
	str = deleteDupMinus("----Show--must----")
	assert.Equal(t, str, "-Show-must-")
}
func TestGeRepoNameFromURL(t *testing.T) {
	topicid := getTaskIDFromURL("https://dev.heeus.io/launchpad/#!13427")
	assert.Equal(t, topicid, "13427")
}

func TestGetBranchName(t *testing.T) {
	str, _ := getBranchName(false, "Show", "must", "go", "on", "https://dev.heeus.io/launchpad/#!13427")
	assert.Equal(t, str, "13427-Show-must-go-on")
	str, _ = getBranchName(false, "Show   ivv?", "must    ", "go", "on---", "https://dev.heeus.io/launchpad/#!13427")
	assert.Equal(t, str, "13427-Show-ivv-must-go-on")
	str, _ = getBranchName(false, "Show", "must", "go", "on")
	assert.Equal(t, str, "Show-must-go-on")
	str, _ = getBranchName(false, "Show")
	assert.Equal(t, str, "Show")
	str, _ = getBranchName(false, "Show   ivv? must $   go on---  https://dev.heeus.io/launchpad/#!13427")
	assert.Equal(t, str, "13427-Show-ivv-must-go-on")
	str, _ = getBranchName(false, "Show   ivv? must $   ", "go on---  https://dev.heeus.io/launchpad/#!13427")
	assert.Equal(t, str, "13427-Show-ivv-must-go-on")
	str, _ = getBranchName(false, "Show   ivv? must $   go  on---", "https://dev.heeus.io/launchpad/#!13427")
	assert.Equal(t, str, "13427-Show-ivv-must-go-on")
	str, _ = getBranchName(false, "Show", "ivv? must $   go  on--- https://dev.heeus.io/launchpad/#!13427")
	assert.Equal(t, str, "13427-Show-ivv-must-go-on")
	str, _ = getBranchName(false, "Show", "ivv? must $   go  on---", "https://dev.heeus.io/launchpad/#!13427")
	assert.Equal(t, str, "13427-Show-ivv-must-go-on")
	str, _ = getBranchName(false, "q", "dev", "https://dev.heeus.io/launchpad/#!13427")
	assert.Equal(t, str, "13427-q-dev")
	str, _ = getBranchName(false, "q", "dev", "https://dev.heeus.io/launchpad/#!13427")
	assert.Equal(t, str, "13427-q-dev")

	//Logn name
	str, _ = getBranchName(false, "Show", "me this  very long string more than fifty symbols in lenth with long task number 11111111111111", "https://dev.heeus.io/launchpad/#!13427")
	assert.Equal(t, str, "13427-Show-me-this-very-long-string-more-than-fift")

	//URL name
	str, _ = getBranchName(false, "https://www.projectkaiser.com/online/#!3206802")
	assert.Equal(t, str, "www-projectkaiser-com-online-#-3206802")

	str, _ = getBranchName(false, "https://github.com/voedger/voedger/issues/395")
	assert.Equal(t, str, "github-com-voedger-voedger-issues-395")

}

func TestClipBoard(t *testing.T) {
	err := clipboard.WriteAll("1,2,3,5")
	assert.Nil(t, err)

	arg, _ := clipboard.ReadAll()

	args := strings.Split(arg, "\n")
	var newarg string
	for _, str := range args {
		newarg += str
		newarg += " "
	}
	assert.NotEmpty(t, newarg)
}

func TestIssueRepoFromURL(t *testing.T) {
	repo := git.GetIssuerepoFromUrl("https://github.com/untillpro/qs/issues/24")
	assert.Equal(t, repo, "untillpro/qs")
}
