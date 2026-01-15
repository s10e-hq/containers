package main

import (
	"context"
	"testing"

	"github.com/s10e-hq/containers/testhelpers"
)

func Test(t *testing.T) {
	ctx := context.Background()
	image := testhelpers.GetTestImage("ghcr.io/home-operations/prowlarr:rolling")
	testhelpers.TestHTTPEndpoint(t, ctx, image, testhelpers.HTTPTestConfig{Port: "9696"}, nil)
}
