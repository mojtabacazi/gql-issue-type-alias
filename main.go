package main

import (
	"app/graphql/generated"
	"app/graphql/resolvers"
)

//go:generate go  run github.com/99designs/gqlgen

func main() {
	rsv := resolvers.Resolver{}

	_ = generated.Config{Resolvers: &rsv}
	println("This is a sample project")
}
