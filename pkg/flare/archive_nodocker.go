// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build !docker

package flare

import (
	workloadmeta "github.com/DataDog/datadog-agent/comp/core/workloadmeta/def"
	"github.com/DataDog/datadog-agent/pkg/util/optional"
)

func getDockerSelfInspect(_ optional.Option[workloadmeta.Component]) ([]byte, error) {
	return nil, nil
}

func getDockerPs() ([]byte, error) {
	return nil, nil
}
