// Copyright 2014-2015 The Gogs Authors. All rights reserved.
// Copyright 2015 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"os"
	"os/user"
)

func CurrentUsername() string {
	curUserName := os.Getenv("USER")
	if len(curUserName) > 0 {
		return curUserName
	}

	curUserName = os.Getenv("USERNAME")
	if len(curUserName) > 0 {
		return curUserName
	}

	curUser, err := user.Current()
	if err == nil {
		return curUser.Username
	}

	return ""
}
