# gql-issue-type-alias

Run `go run main.go` to see the error. 

```
./main.go:13:23: cannot use &rsv (type *resolvers.Resolver) as type generated.ResolverRoot in field value:
        *resolvers.Resolver does not implement generated.ResolverRoot (missing SampleInput method)
```

Switch to gqlgen 0.14 or older by modifying `go.mod` file and run `go generate ./...` then it works just fine.
