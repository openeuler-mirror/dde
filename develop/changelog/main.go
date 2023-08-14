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
	"flag"
	"fmt"
	"os"
)

var (
	changelogNum        int
	useShortEntry       bool
	useDefaultEditor    bool
	gitPath             string
	outputChangelogFile string
	version             string
)

func parseArgs() {
	flag.IntVar(&changelogNum, "c", 1, "num of changelog need to create")
	flag.BoolVar(&useShortEntry, "s", false, "use short entry")
	flag.BoolVar(&useDefaultEditor, "e", false, "useDefaultEditor. Windows: notepad; Linux: vi")
	flag.StringVar(&gitPath, "g", ".", "choose the git path")
	flag.StringVar(&outputChangelogFile, "o", "", "output file")
	flag.StringVar(&version, "v", "", "set changelog version")
	flag.Parse()
}

func main() {
	parseArgs()

	logs, err := getLog()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	tempFile, err := initChangelog(logs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	handler, err := changelogHandler(tempFile, logs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	changelog, err := run(handler, tempFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	fmt.Println(*changelog)
	if outputChangelogFile != "" {
		err := os.WriteFile(outputChangelogFile, []byte(*changelog), 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not open '%s' for writing: %s\n", tempFile, err.Error())
		}
	}
}
