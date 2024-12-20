// Copyright (C) 2024 Yota Hamada
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"testing"

	"github.com/dagu-org/dagu/internal/test"
)

func TestStartCommand(t *testing.T) {
	tests := []cmdTest{
		{
			args:        []string{"start", testDAGFile("success.yaml")},
			expectedOut: []string{"1 finished"},
		},
		{
			args:        []string{"start", testDAGFile("params.yaml")},
			expectedOut: []string{"params is p1 and p2"},
		},
		{
			args: []string{
				"start",
				`--params="p3 p4"`,
				testDAGFile("params.yaml"),
			},
			expectedOut: []string{"params is p3 and p4"},
		},
	}

	th := test.Setup(t)
	for _, tc := range tests {
		testRunCommand(t, th.Context, startCmd(), tc)
	}
}
