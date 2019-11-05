// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/types"
	"github.com/go-vela/types/library"
)

const (
	// BuildResp represents a JSON return for a single build.
	BuildResp = `  
	{
		"id": 1,
		"repo_id": 1,
		"number": 1,
		"parent": 1,
		"event": "push",
		"status": "created",
		"error": "",
		"enqueued": 1563474077,
		"created": 1563474076,
		"started": 1563474077,
		"finished": 0,
		"deploy": "",
		"clone": "https://github.com/github/octocat.git",
		"source": "https://github.com/github/octocat/commit/48afb5bdc41ad69bf22588491333f7cf71135163",
		"title": "push received from https://github.com/github/octocat",
		"message": "First commit...",
		"commit": "48afb5bdc41ad69bf22588491333f7cf71135163",
		"sender": "OctoKitty",
		"author": "OctoKitty",
		"branch": "master",
		"ref": "refs/heads/master",
		"base_ref": "",
		"host": "example.company.com",
		"runtime": "docker",
		"distribution": "linux"
	}`

	// BuildsResp represents a JSON return for one to many builds.
	BuildsResp = `[
		{
			"id": 2,
			"repo_id": 1,
			"number": 2,
			"parent": 1,
			"event": "push",
			"status": "running",
			"error": "",
			"enqueued": 1563474204,
			"created": 1563474204,
			"started": 1563474204,
			"finished": 0,
			"deploy": "",
			"clone": "https://github.com/github/octocat.git",
			"source": "https://github.com/github/octocat/commit/48afb5bdc41ad69bf22588491333f7cf71135163",
			"title": "push received from https://github.com/github/octocat",
			"message": "Second commit...",
			"commit": "48afb5bdc41ad69bf22588491333f7cf71135163",
			"sender": "OctoKitty",
			"author": "OctoKitty",
			"branch": "master",
			"ref": "refs/heads/master",
			"base_ref": "",
			"host": "ed95dcc0687c",
			"runtime": "",
			"distribution": ""
		},
		{
			"id": 1,
			"repo_id": 1,
			"number": 1,
			"parent": 1,
			"event": "push",
			"status": "running",
			"error": "",
			"enqueued": 1563474077,
			"created": 1563474076,
			"started": 1563474077,
			"finished": 0,
			"deploy": "",
			"clone": "https://github.com/github/octocat.git",
			"source": "https://github.com/github/octocat/commit/48afb5bdc41ad69bf22588491333f7cf71135163",
			"title": "push received from https://github.com/github/octocat",
			"message": "First commit...",
			"commit": "48afb5bdc41ad69bf22588491333f7cf71135163",
			"sender": "OctoKitty",
			"author": "OctoKitty",
			"branch": "master",
			"ref": "refs/heads/master",
			"base_ref": "",
			"host": "82823eb770b0",
			"runtime": "",
			"distribution": ""
		}
	]`

	// BuildLogsResp represents a JSON return for build logs.
	BuildLogsResp = `[
		{
			"id": 1,
			"step_id": 1,
			"build_id": 1,
			"repo_id": 1,
			"data": "SGVsbG8sIFdvcmxkIQ=="
		},
		{
			"id": 2,
			"step_id": 2,
			"build_id": 1,
			"repo_id": 1,
			"data": "SGVsbG8sIFdvcmxkIQ=="
		}
	]`
)

// getBuilds returns mock JSON for a http GET.
func getBuilds(c *gin.Context) {

	data := []byte(BuildsResp)
	var body []library.Build
	_ = json.Unmarshal(data, &body)

	c.JSON(http.StatusOK, body)
}

// getBuild has a param :build returns mock JSON for a http GET.
func getBuild(c *gin.Context) {
	b := c.Param("build")

	if strings.EqualFold(b, "0") {
		msg := fmt.Sprintf("Build %s does not exist", b)
		c.AbortWithStatusJSON(404, types.Error{Message: &msg})
		return
	}

	data := []byte(BuildResp)
	var body library.Build
	_ = json.Unmarshal(data, &body)

	c.JSON(http.StatusOK, body)
}

// getLogs has a param :build returns mock JSON for a http GET.
//
// Pass "0" to :build to test receiving a http 404 response
func getLogs(c *gin.Context) {
	b := c.Param("build")

	if strings.EqualFold(b, "0") {
		msg := fmt.Sprintf("Build %s does not exist", b)
		c.AbortWithStatusJSON(404, types.Error{Message: &msg})
		return
	}

	data := []byte(BuildLogsResp)
	var body []library.Log
	_ = json.Unmarshal(data, &body)

	c.JSON(http.StatusOK, body)
}

// addBuild returns mock JSON for a http POST.
func addBuild(c *gin.Context) {
	data := []byte(BuildResp)
	var body library.Build
	_ = json.Unmarshal(data, &body)

	c.JSON(http.StatusCreated, body)
}

// updateBuild has a param :build returns mock JSON for a http PUT.
//
// Pass "0" to :build to test receiving a http 404 response
func updateBuild(c *gin.Context) {
	b := c.Param("build")

	if strings.EqualFold(b, "0") {
		msg := fmt.Sprintf("Build %s does not exist", b)
		c.AbortWithStatusJSON(404, types.Error{Message: &msg})
		return
	}

	data := []byte(BuildResp)
	var body library.Build
	_ = json.Unmarshal(data, &body)

	c.JSON(http.StatusOK, body)
}

// removeBuild has a param :build returns mock JSON for a http DELETE.
//
// Pass "0" to :build to test receiving a http 404 response
func removeBuild(c *gin.Context) {
	b := c.Param("build")

	if strings.EqualFold(b, "0") {
		msg := fmt.Sprintf("Build %s does not exist", b)
		c.AbortWithStatusJSON(404, types.Error{Message: &msg})
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("Build %s removed", b))
}

// restartBuild has a param :build returns mock JSON for a http POST.
//
// Pass "0" to :build to test receiving a http 404 response
func restartBuild(c *gin.Context) {
	b := c.Param("build")

	if strings.EqualFold(b, "0") {
		msg := fmt.Sprintf("Build %s does not exist", b)
		c.AbortWithStatusJSON(404, types.Error{Message: &msg})
		return
	}

	data := []byte(BuildResp)
	var body library.Build
	_ = json.Unmarshal(data, &body)

	c.JSON(http.StatusCreated, body)
}