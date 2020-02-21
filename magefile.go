// +build mage

package main

import (
	// mage:import
	_ "github.com/getcouragenow/bootstrap/ci/os"

	// mage:import
	_ "github.com/getcouragenow/bootstrap/ci/tools"

	// mage:import
	_ "github.com/getcouragenow/bootstrap/ci/server_tools"

	// mage:import
	_ "github.com/getcouragenow/bootstrap/ci/help"
)
