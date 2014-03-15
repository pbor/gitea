// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"net/http"
	"strconv"

	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"github.com/gogits/gogs/models"
	"github.com/gogits/gogs/modules/auth"
	"github.com/gogits/gogs/modules/base"
	"github.com/gogits/gogs/modules/log"
)

func Setting(form auth.UpdateProfileForm, r render.Render, data base.TmplData, req *http.Request, session sessions.Session) {
	data["Title"] = "Setting"
	data["PageIsUserSetting"] = true

	user := auth.SignedInUser(session)
	data["Owner"] = user

	if req.Method == "GET" {
		r.HTML(200, "user/setting", data)
		return
	}

	if hasErr, ok := data["HasError"]; ok && hasErr.(bool) {
		r.HTML(200, "user/setting", data)
		return
	}

	user.Email = form.Email
	user.Website = form.Website
	user.Location = form.Location
	user.Avatar = base.EncodeMd5(form.Avatar)
	user.AvatarEmail = form.Avatar
	if err := models.UpdateUser(user); err != nil {
		log.Handle(200, "setting.Setting", data, r, err)
		return
	}

	data["IsSuccess"] = true
	r.HTML(200, "user/setting", data)
}

func SettingPassword(form auth.UpdatePasswdForm, r render.Render, data base.TmplData, session sessions.Session, req *http.Request) {
	data["Title"] = "Password"
	data["PageIsUserSetting"] = true

	if req.Method == "GET" {
		r.HTML(200, "user/password", data)
		return
	}

	user := auth.SignedInUser(session)
	newUser := &models.User{Passwd: form.NewPasswd}
	if err := newUser.EncodePasswd(); err != nil {
		log.Handle(200, "setting.SettingPassword", data, r, err)
		return
	}

	if user.Passwd != newUser.Passwd {
		data["HasError"] = true
		data["ErrorMsg"] = "Old password is not correct"
	} else if form.NewPasswd != form.RetypePasswd {
		data["HasError"] = true
		data["ErrorMsg"] = "New password and re-type password are not same"
	} else {
		user.Passwd = newUser.Passwd
		if err := models.UpdateUser(user); err != nil {
			log.Handle(200, "setting.SettingPassword", data, r, err)
			return
		}
		data["IsSuccess"] = true
	}

	data["Owner"] = user
	r.HTML(200, "user/password", data)
}

func SettingSSHKeys(form auth.AddSSHKeyForm, r render.Render, data base.TmplData, req *http.Request, session sessions.Session) {
	data["Title"] = "SSH Keys"

	// Delete SSH key.
	if req.Method == "DELETE" || req.FormValue("_method") == "DELETE" {
		println(1)
		id, err := strconv.ParseInt(req.FormValue("id"), 10, 64)
		if err != nil {
			data["ErrorMsg"] = err
			log.Error("ssh.DelPublicKey: %v", err)
			r.JSON(200, map[string]interface{}{
				"ok":  false,
				"err": err.Error(),
			})
			return
		}
		k := &models.PublicKey{
			Id:      id,
			OwnerId: auth.SignedInId(session),
		}

		if err = models.DeletePublicKey(k); err != nil {
			data["ErrorMsg"] = err
			log.Error("ssh.DelPublicKey: %v", err)
			r.JSON(200, map[string]interface{}{
				"ok":  false,
				"err": err.Error(),
			})
		} else {
			r.JSON(200, map[string]interface{}{
				"ok": true,
			})
		}
		return
	}

	// Add new SSH key.
	if req.Method == "POST" {
		if hasErr, ok := data["HasError"]; ok && hasErr.(bool) {
			r.HTML(200, "user/publickey", data)
			return
		}

		k := &models.PublicKey{OwnerId: auth.SignedInId(session),
			Name:    form.KeyName,
			Content: form.KeyContent,
		}

		if err := models.AddPublicKey(k); err != nil {
			data["ErrorMsg"] = err
			log.Error("ssh.AddPublicKey: %v", err)
			r.HTML(200, "base/error", data)
			return
		} else {
			data["AddSSHKeySuccess"] = true
		}
	}

	// List existed SSH keys.
	keys, err := models.ListPublicKey(auth.SignedInId(session))
	if err != nil {
		data["ErrorMsg"] = err
		log.Error("ssh.ListPublicKey: %v", err)
		r.HTML(200, "base/error", data)
		return
	}

	data["PageIsUserSetting"] = true
	data["Keys"] = keys
	r.HTML(200, "user/publickey", data)
}

func SettingNotification(r render.Render, data base.TmplData) {
	// todo user setting notification
	data["Title"] = "Notification"
	data["PageIsUserSetting"] = true
	r.HTML(200, "user/notification", data)
}

func SettingSecurity(r render.Render, data base.TmplData) {
	// todo user setting security
	data["Title"] = "Security"
	data["PageIsUserSetting"] = true
	r.HTML(200, "user/security", data)
}