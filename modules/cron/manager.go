// Copyright 2014-2015 The Gogs Authors. All rights reserved.
// Copyright 2015 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cron

import (
	"fmt"

	"github.com/go-gitea/gitea/models"
	"github.com/go-gitea/gitea/modules/setting"
)

var c = New()

func NewCronContext() {
	c.AddFunc("Update mirrors", "@every 1h", models.MirrorUpdate)
	c.AddFunc("Update wikis", "@every 10m", models.WikiUpdate)
	if setting.Git.Fsck.Enable {
		c.AddFunc("Health checks", fmt.Sprintf("@every %dh", setting.Git.Fsck.Interval), models.GitFsck)
	}
	c.Start()
}

func ListEntries() []*Entry {
	return c.Entries()
}
