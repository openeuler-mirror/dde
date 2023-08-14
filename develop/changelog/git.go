/*
 * Copyright (c) 2023 lizengyi
 * changelog is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 */

package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"strings"
)

const (
	HASHWidth = 7
	TTYWidth  = 120
)

const (
	AppendTodoHelp = "\n" +
		"# Please pick the commit message for your changelogs. Lines starting\n" +
		"# with '#' will be ignored, and an empty message aborts the changelog.\n" +
		"# \n" +
		"# Commands:\n" +
		"# p, pick <commit> = use commit as changelog\n" +
		"# s, squash <commit> = use commit as changelog, but meld into previous commit\n" +
		"# d, drop <commit> = remove commit"
)

func getLog() (logs []*object.Commit, err error) {
	r, err := git.PlainOpen(gitPath)
	if err != nil {
		return
	}
	tmp, err := r.Config()
	gitCoreEditor = tmp.Raw.Section("core").Options.Get("editor")
	ref, err := r.Head()
	if err != nil {
		return
	}

	cc, err := r.Log(&git.LogOptions{From: ref.Hash(), Order: git.LogOrderDefault})
	if err != nil {
		return
	}

	for i := 0; i < changelogNum; i++ {
		commit, _ := cc.Next()
		if len(commit.ParentHashes) != 1 {
			i--
			continue
		}
		logs = append(logs, commit)
	}
	return logs, nil
}

func processString(input string) string {
	if index := strings.Index(input, "\n"); index != -1 {
		input = input[:index+1]
	}

	if len(input) > TTYWidth {
		input = input[:TTYWidth]
	}

	return input
}

func genLog(logs []*object.Commit) *string {
	var logBuffer string
	for _, log := range logs {
		logBuffer += "pick "
		logBuffer += log.Hash.String()[0:HASHWidth]
		logBuffer += " "
		logBuffer += processString(log.Message)
	}
	logBuffer += AppendTodoHelp
	return &logBuffer
}

func authorExist(authors *[]string, author string) bool {
	for _, a := range *authors {
		if author == a {
			return true
		}
	}
	return false
}

func genChangelog(logs []*object.Commit, v Version) *string {
	var changelogBuffer string
	dateFormat := "* Mon Jan 02 2006 "
	changelogDate := logs[0].Committer.When.Format(dateFormat)
	var author []string
	var changelogEntry string
	for _, log := range logs {
		if !authorExist(&author, log.Author.Name+" <"+log.Author.Email+">") {
			author = append(author, log.Author.Name+" <"+log.Author.Email+">")
		}

		if useShortEntry {
			changelogEntry += processString(log.Message)
		} else {
			changelogEntry += log.Message
		}

	}
	changelogEntry = strings.Replace(changelogEntry, "\n", "\n- ", -1)

	lastIndex := strings.LastIndex(changelogEntry, "\n- ")
	if lastIndex != -1 {
		changelogEntry = changelogEntry[:lastIndex+1]
	}

	authorNameAndEmail := strings.Join(author, ", ")

	versionStr := v.value
	if versionStr == "" {
		versionStr = v.String()
	}
	if version != "" {
		changelogBuffer = changelogDate + authorNameAndEmail + " - " + versionStr + "\n- " + changelogEntry
	} else {
		changelogBuffer = changelogDate + authorNameAndEmail + "\n- " + changelogEntry
	}

	return &changelogBuffer
}
