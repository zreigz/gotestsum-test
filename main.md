package main

import (
	"path/filepath"

	"gotest.tools/gotestsum/cmd"
	"k8s.io/klog/v2"
)

const (
	junitfile = "unit-tests.xml"
	jsonFile  = "unit-tests.json"
	outputDir = "test-output"
)

func main() {

	args := []string{
		"--format", "testname",
		"--junitfile", junitfile,
		"--jsonfile", filepath.Join(outputDir, jsonFile),
	}
	args = append(args, "--")
	goTestArgs := []string{
		"--test.count=1",
	}
	args = append(args, goTestArgs...)

	if err := cmd.Run("", args); err != nil {
		klog.Warningf("gotestsum returned an error: %v", err)
	}
}
