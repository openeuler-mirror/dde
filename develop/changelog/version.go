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
	"regexp"
	"strconv"
)

type Version struct {
	value        string
	X, Y, Z, TAG int
}

func (v *Version) InitVersion(versionStr string) (err error) {
	re := regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)-(\d+)$`)
	matches := re.FindStringSubmatch(versionStr)

	if len(matches) != 5 {
		return fmt.Errorf("invalid version string: %s", versionStr)
	}

	v.X, err = strconv.Atoi(matches[1])
	if err != nil {
		return err
	}

	v.Y, err = strconv.Atoi(matches[2])
	if err != nil {
		return err
	}

	v.Z, err = strconv.Atoi(matches[3])
	if err != nil {
		return err
	}

	v.TAG, err = strconv.Atoi(matches[4])
	if err != nil {
		return err
	}
	v.value = ""
	return nil
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d-%d", v.X, v.Y, v.Z, v.TAG)
}

func (v *Version) ADDTag() {
	v.TAG++
}
