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
	"bufio"
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/object"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type ChangelogItem struct {
	log []*object.Commit
}

func initChangelog(logs []*object.Commit) (string, error) {
	tempFile := createTempFile()
	logBuffer := genLog(logs)
	err := os.WriteFile(tempFile, []byte(*logBuffer), 0666)
	if err != nil {
		return "", fmt.Errorf("could not open '%s' for writing: %w", tempFile, err)
	}
	return tempFile, nil
}

func createTempFile() string {
	tempFile, err := ioutil.TempFile("", "temp_file_*")
	if err != nil {
		fmt.Println("无法创建临时文件:", err)
		os.Exit(1)
	}
	defer tempFile.Close()

	return tempFile.Name()
}

func changelogHandler(filename string, logs []*object.Commit) (logItems []ChangelogItem, err error) {
	err = launchEditor(filename)
	if err != nil {
		return nil, fmt.Errorf("could not edit '%s': %w", filename, err)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#") || len(strings.TrimSpace(line)) == 0 {
			continue
		}

		parts := strings.SplitN(line, " ", 3)

		if len(parts) != 3 {
			return nil, fmt.Errorf("could not parse line: %d", count)
		}

		if !strings.HasPrefix(logs[count].Hash.String(), parts[1]) {
			return nil, fmt.Errorf("could not parse '%s'", parts[1])
		}

		if strings.HasPrefix("pick", parts[0]) {
			tmp := new(ChangelogItem)
			tmp.log = append(tmp.log, logs[count])
			logItems = append(logItems, *tmp)
		} else if strings.HasPrefix("squash", parts[0]) {
			if count == 0 {
				return nil, fmt.Errorf("cannot 'squash' without a previous commit")
			}
			logItems[len(logItems)-1].log = append(logItems[len(logItems)-1].log, logs[count])
		} else if strings.HasPrefix("drop", parts[0]) {
		} else {
			return nil, fmt.Errorf("could not parse '%s'", parts[0])
		}

		count++
	}
	if count == 0 {
		return nil, fmt.Errorf("error: nothing to do")
	}
	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return
}

func parseVersion(changelogText string, v *Version) error {
	firstNewlinePos := strings.Index(changelogText, "\n")
	if firstNewlinePos == -1 {
		return fmt.Errorf("fail to parse changelog")
	}

	firstLine := changelogText[:firstNewlinePos]

	versionRegex := regexp.MustCompile(`- (\d+\.\d+\.\d+-\d+)`)
	matches := versionRegex.FindStringSubmatch(firstLine)

	if len(matches) == 2 {
		versionStr := matches[1]
		var temp Version
		err := temp.InitVersion(versionStr)
		if err != nil {
			v.value = versionStr
			return fmt.Errorf("connot parse version in the Changelog")
		}
		// TODO: 版本号自增模式
		if temp.X != v.X || temp.Y != v.Y || temp.Z != v.Z {
			v.TAG = 0
		} else {
			v.TAG = temp.TAG
		}
		v.X = temp.X
		v.Y = temp.Y
		v.Z = temp.Z
	} else {
		return fmt.Errorf("version not found in the Changelog")
	}
	return nil
}

func run(logItems []ChangelogItem, tempFile string) (changelog *string, err error) {
	changelog = new(string)
	var v Version
	if version != "" {
		if v.InitVersion(version) != nil {
			v.value = version
		}
	}
	// TODO: 正序倒序
	//for _, item := range logItems {
	for i := len(logItems) - 1; i >= 0; i-- {
		item := logItems[i]
		buf := genChangelog(item.log, v)

		err := os.WriteFile(tempFile, []byte(*buf), 0666)
		if err != nil {
			return nil, fmt.Errorf("could not open '%s' for writing: %w", tempFile, err)
		}

		err = launchEditor(tempFile)
		if err != nil {
			return nil, fmt.Errorf("could not edit '%s': %w", tempFile, err)
		}

		data, err := os.ReadFile(tempFile)
		if err != nil {
			return nil, fmt.Errorf("could not read file '%s': %w", tempFile, err)
		}

		if version != "" {
			if parseVersion(strings.TrimSpace(string(data)), &v) == nil {
				v.ADDTag()
			}
		}

		if len(*changelog) > 0 && !strings.HasPrefix(*changelog, "\n\n") {
			*changelog = "\n\n" + *changelog + "\n"
		}
		*changelog = strings.TrimSpace(string(data)) + *changelog
		//lastIndex := len(*changelog) - 1
		//if lastIndex >= 0 {
		//	*changelog += "\n\n"
		//}
		//*changelog += strings.TrimSpace(string(data))
	}
	return
}
