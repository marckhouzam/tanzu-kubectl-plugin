// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/vmware-tanzu/tanzu-framework/cli/runtime/component"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/buildinfo"
)

func main() {
	fmt.Println("version:", buildinfo.Version)
	component.AskForConfirmation("Are you happy?")
}
