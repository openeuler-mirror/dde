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
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var gitCoreEditor string

const (
	//editorProgram        = "" // choose ur own editor
	defaultEditorWindows = "notepad"
	defaultEditorLinux   = "vi"
)

func isTerminalDumb() bool {
	terminal := os.Getenv("TERM")
	return terminal == "" || strings.ToLower(terminal) == "dumb"
}

func gitEditor() string {
	if useDefaultEditor {
		if runtime.GOOS == "windows" {
			return defaultEditorWindows
		} else {
			return defaultEditorLinux
		}
	}

	editor := os.Getenv("GIT_EDITOR")

	if editor == "" && gitCoreEditor != "" {
		editor = gitCoreEditor
	}

	terminalIsDumb := isTerminalDumb()
	if editor == "" && !terminalIsDumb {
		editor = os.Getenv("VISUAL")
	}

	if editor == "" {
		editor = os.Getenv("EDITOR")
	}

	if editor == "" && runtime.GOOS == "windows" {
		editor = defaultEditorWindows
	} else if editor == "" && !terminalIsDumb {
		editor = defaultEditorLinux
	}

	return editor
}

func launchSpecifiedEditor(editor, path string) error {
	if editor == "" {
		return fmt.Errorf("terminal is dumb, but EDITOR unset")
	}

	if editor != ":" {
		realpath, err := exec.LookPath(editor)
		if err != nil {
			return fmt.Errorf("unable to start editor '%s', use -e to lunch defult editor", editor)
		}

		cmd := exec.Command(realpath, path)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("there was a problem with the editor '%s', %w", editor, err)
		}
	}

	return nil
}

func launchEditor(path string) error {
	return launchSpecifiedEditor(gitEditor(), path)
}
