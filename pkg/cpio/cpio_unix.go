// Copyright 2013-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows
// +build !windows

package cpio

import (
	"github.com/u-root/u-root/pkg/ls"
)

// String implements a fmt.Stringer for Record.
//
// String returns a string long-formatted like `ls` would format it.
func (r Record) String() string {
	s := ls.LongStringer{
		Human: true,
		Name:  ls.NameStringer{},
	}
	return s.FileString(LSInfoFromRecord(r))
}
