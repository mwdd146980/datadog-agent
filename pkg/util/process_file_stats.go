// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package util

type ProcessFileStats struct {
	AgentOpenFiles float64 `json:"agent_open_files"`
	OsFileLimit    float64 `json:"os_file_limit"`
}
