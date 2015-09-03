// +build go1.2

// Copyright 2014-2015 The Gogs Authors. All rights reserved.
// Copyright 2015 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Gitea (Git service with a cup of tea) is a painless self-hosted Git Service written in Go.
package main

import (
	"os"
	"runtime"

	"github.com/codegangsta/cli"

	"github.com/go-gitea/gitea/cmd"
	"github.com/go-gitea/gitea/modules/setting"
)

var (
	version  = "0.7.0-beta0"
	revision = ""
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	setting.AppVer = version
}

func main() {
	app := cli.NewApp()
	app.Name = "Gitea"
	app.Usage = "Git service with a cup of tea"
	app.Version = version
	app.Commands = []cli.Command{
		cmd.CmdWeb,
		cmd.CmdServ,
		cmd.CmdUpdate,
		cmd.CmdDump,
		cmd.CmdGenerate,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
