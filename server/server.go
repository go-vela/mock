// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// FakeHandler returns an http.Handler that is capable of handling
// Vela API requests and returning mock responses.
func FakeHandler() http.Handler {
	gin.SetMode(gin.TestMode)

	e := gin.New()

	// mock endpoints for build calls
	e.GET("/api/v1/repos/:org/:repo/builds/:build", getBuild)
	e.POST("/api/v1/repos/:org/:repo/builds/:build", restartBuild)
	e.GET("/api/v1/repos/:org/:repo/builds/:build/logs", getLogs)
	e.GET("/api/v1/repos/:org/:repo/builds", getBuilds)
	e.POST("/api/v1/repos/:org/:repo/builds", addBuild)
	e.PUT("/api/v1/repos/:org/:repo/builds/:build", updateBuild)
	e.DELETE("/api/v1/repos/:org/:repo/builds/:build", removeBuild)

	// mock endpoints for log calls
	e.GET("/api/v1/repos/:org/:repo/builds/:build/services/:service/logs", getServiceLog)
	e.POST("/api/v1/repos/:org/:repo/builds/:build/services/:service/logs", addServiceLog)
	e.PUT("/api/v1/repos/:org/:repo/builds/:build/services/:service/logs", updateServiceLog)
	e.DELETE("/api/v1/repos/:org/:repo/builds/:build/services/:service/logs", removeServiceLog)
	e.GET("/api/v1/repos/:org/:repo/builds/:build/steps/:step/logs", getStepLog)
	e.POST("/api/v1/repos/:org/:repo/builds/:build/steps/:step/logs", addStepLog)
	e.PUT("/api/v1/repos/:org/:repo/builds/:build/steps/:step/logs", updateStepLog)
	e.DELETE("/api/v1/repos/:org/:repo/builds/:build/steps/:step/logs", removeStepLog)

	// mock endpoints for repo calls
	e.GET("/api/v1/repos/:org/:repo", getRepo)
	e.GET("/api/v1/repos", getRepos)
	e.POST("/api/v1/repos", addRepo)
	e.PUT("/api/v1/repos/:org/:repo", updateRepo)
	e.DELETE("/api/v1/repos/:org/:repo", removeRepo)
	e.PATCH("/api/v1/repos/:org/:repo/repair", repairRepo)
	e.PATCH("/api/v1/repos/:org/:repo/chown", chownRepo)

	// mock endpoints for secret calls
	e.GET("/api/v1/secrets/:engine/:type/:org/:name/:secret", getSecret)
	e.GET("/api/v1/secrets/:engine/:type/:org/:name", getSecrets)
	e.POST("/api/v1/secrets/:engine/:type/:org/:name", addSecret)
	e.PUT("/api/v1/secrets/:engine/:type/:org/:name/:secret", updateSecret)
	e.DELETE("/api/v1/secrets/:engine/:type/:org/:name/:secret", removeSecret)

	// mock endpoints for step calls
	e.GET("/api/v1/repos/:org/:repo/builds/:build/steps/:step", getStep)
	e.GET("/api/v1/repos/:org/:repo/builds/:build/steps", getSteps)
	e.POST("/api/v1/repos/:org/:repo/builds/:build/steps", addStep)
	e.PUT("/api/v1/repos/:org/:repo/builds/:build/steps/:step", updateStep)
	e.DELETE("/api/v1/repos/:org/:repo/builds/:build/steps/:step", removeStep)

	// mock endpoints for service calls
	e.GET("/api/v1/repos/:org/:repo/builds/:build/services/:service", getService)
	e.GET("/api/v1/repos/:org/:repo/builds/:build/services", getServices)
	e.POST("/api/v1/repos/:org/:repo/builds/:build/services", addService)
	e.PUT("/api/v1/repos/:org/:repo/builds/:build/services/:service", updateService)
	e.DELETE("/api/v1/repos/:org/:repo/builds/:build/services/:service", removeService)

	return e
}
