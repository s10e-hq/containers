package main

import (
	"context"
	"testing"

	"github.com/s10e-hq/containers/testhelpers"
)

func Test(t *testing.T) {
	ctx := context.Background()
	image := testhelpers.GetTestImage("ghcr.io/s10e-hq/decyparr:rolling")
	testhelpers.TestHTTPEndpoint(t, ctx, image, testhelpers.HTTPTestConfig{Port: "8282"}, nil)
}
