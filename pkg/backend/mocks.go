package backend

import (
	"context"

	opentracing "github.com/opentracing/opentracing-go"

	"sourcegraph.com/sourcegraph/sourcegraph/pkg/accesscontrol"
	authpkg "sourcegraph.com/sourcegraph/sourcegraph/pkg/auth"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/localstore"
)

var Mocks MockServices

type MockServices struct {
	Defs     MockDefs
	Pkgs     MockPkgs
	RepoTree MockRepoTree
	Repos    MockRepos
	Orgs     MockOrgs
}

// testContext creates a new context.Context for use by tests
func testContext() context.Context {
	localstore.Mocks = localstore.MockStores{}
	Mocks = MockServices{}

	ctx := context.Background()
	ctx = authpkg.WithActor(ctx, &authpkg.Actor{UID: "1", Login: "test"})
	ctx = accesscontrol.WithInsecureSkip(ctx, true)
	_, ctx = opentracing.StartSpanFromContext(ctx, "dummy")

	return ctx
}
