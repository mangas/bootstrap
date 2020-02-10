// +build mage

package main

import (

	// mage:import
	_ "github.com/getcouragenow/bootstrap/ci/gsheet"

	// mage:import
	_ "github.com/getcouragenow/bootstrap/ci/dep"

	// mage:import
	_ "github.com/getcouragenow/bootstrap/ci/mage_select"

	// mage:import
	_ "github.com/getcouragenow/bootstrap/ci/hugo"

	// mage:import
	_ "github.com/getcouragenow/bootstrap/ci/bootstrap"

	// mage:import
	_ "github.com/getcouragenow/bootstrap/ci/hover"
)
