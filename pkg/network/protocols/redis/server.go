// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package redis

import (
	"regexp"
	"testing"

	"github.com/DataDog/datadog-agent/pkg/network/protocols/http/testutil"
	protocolsUtils "github.com/DataDog/datadog-agent/pkg/network/protocols/testutil"
)

func RunRedisServer(t *testing.T, serverAddr, serverPort string) {
	env := []string{
		"REDIS_ADDR=" + serverAddr,
		"REDIS_PORT=" + serverPort,
	}

	t.Helper()
	dir, _ := testutil.CurDir()
	protocolsUtils.RunDockerServer(t, "redis", dir+"/testdata/docker-compose.yml", env, regexp.MustCompile(".*Ready to accept connections"))
}
