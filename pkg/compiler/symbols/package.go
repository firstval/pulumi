// Copyright 2016 Marapongo, Inc. All rights reserved.

package symbols

import (
	"github.com/marapongo/mu/pkg/diag"
	"github.com/marapongo/mu/pkg/pack"
	"github.com/marapongo/mu/pkg/tokens"
)

// Package is a fully bound package symbol.
type Package struct {
	Node         *pack.Package
	Dependencies PackageMap
	Modules      ModuleMap
}

var _ Symbol = (*Package)(nil)

func (node *Package) symbol()             {}
func (node *Package) Token() tokens.Token { return tokens.Token(node.Node.Name) }
func (node *Package) Tree() diag.Diagable { return node.Node }

// PackageMap is a map from package token to the associated symbols.
type PackageMap map[tokens.Package]*Package

// ModuleMap is a map from module token to the associated symbols.
type ModuleMap map[tokens.Module]*Module