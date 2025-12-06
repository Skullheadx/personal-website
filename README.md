use direnv:

```fish
direnv allow
# or
nix develop

```

update an input, run:

```fish
nix flake update
```

update the go.mod:

```fish
gomod2nix generate
```

run using nix:

```fish
nix run
```
