// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1

import (
	sdk "github.com/go-gitea/go-sdk"

	"github.com/go-gitea/gitea/models"
	"github.com/go-gitea/gitea/modules/base"
	"github.com/go-gitea/gitea/modules/middleware"
)

// GET /users/:username/tokens
func ListAccessTokens(ctx *middleware.Context) {
	tokens, err := models.ListAccessTokens(ctx.User.Id)
	if err != nil {
		ctx.JSON(500, &base.ApiJsonErr{"ListAccessTokens: " + err.Error(), base.DOC_URL})
		return
	}

	apiTokens := make([]*sdk.AccessToken, len(tokens))
	for i := range tokens {
		apiTokens[i] = &sdk.AccessToken{tokens[i].Name, tokens[i].Sha1}
	}
	ctx.JSON(200, &apiTokens)
}

type CreateAccessTokenForm struct {
	Name string `json:"name" binding:"Required"`
}

// POST /users/:username/tokens
func CreateAccessToken(ctx *middleware.Context, form CreateAccessTokenForm) {
	t := &models.AccessToken{
		Uid:  ctx.User.Id,
		Name: form.Name,
	}
	if err := models.NewAccessToken(t); err != nil {
		ctx.JSON(500, &base.ApiJsonErr{"NewAccessToken: " + err.Error(), base.DOC_URL})
		return
	}
	ctx.JSON(201, &sdk.AccessToken{t.Name, t.Sha1})
}
