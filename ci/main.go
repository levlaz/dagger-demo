package main

import (
	"context"
)

const (
	PROJECT = "demo"
	REPO    = "demo"
)

type Demo struct{}

// TODO - fix this part too
// Lint Backend Go Code
// func (d *Demo) Lint(ctx context.Context, dir *Directory) (string, error) {
// 	lintResult, err := dag.
// 		Golang().
// 		WithProject(dir).
// 		GolangciLint(ctx)
// 	if err != nil {
// 		return "", err
// 	}
// 	return lintResult, nil
// }

func (d *Demo) Ci(
	ctx context.Context,
	dir *Directory,
	token *Secret,
) (string, error) {
	// Lint
	// TODO - fix this part
	// out, err := d.Lint(ctx, dir)
	// if err != nil {
	// 	return "", err
	// }
	var out string

	// Scan for Code References
	ldOut, err := dag.Launchdarkly().Find(ctx, token, dir, PROJECT, REPO)
	if err != nil {
		return "", err
	}
	out = out + "\n" + ldOut

	return out, nil
}
