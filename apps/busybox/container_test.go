package main

import (
	"context"
	"testing"

	"github.com/s10e-hq/containers/testhelpers"
)

func Test(t *testing.T) {
	ctx := context.Background()
	image := testhelpers.GetTestImage("ghcr.io/s10e-hq/busybox:rolling")
	testhelpers.TestCommandSucceeds(t, ctx, image, nil, "/bin/busybox", "--list")
}
