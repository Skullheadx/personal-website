{
  pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
    import (fetchTree nixpkgs.locked) {
      overlays = [
        (import "${fetchTree gomod2nix.locked}/overlay.nix")
      ];
    }
  ),
  mkGoEnv ? pkgs.mkGoEnv,
  gomod2nix ? pkgs.gomod2nix,
  pre-commit-hooks,
  ...
}:

let
  goEnv = mkGoEnv { pwd = ./.; };
  pre-commit-check = pre-commit-hooks.lib.${pkgs.system}.run {
    src = ./.;
    hooks = {
      # gofmt.enable = true;  # still disabled — not needed

      goimports = {
        enable = true;
        entry = "${pkgs.gotools}/bin/goimports -w";
        files = "\\.go$";
      };

      golangci-lint = {
        enable = true;
        entry = "${pkgs.golangci-lint}/bin/golangci-lint run --fix";
        files = "\\.go$";
        pass_filenames = false;
      };
    };
  };
in
pkgs.mkShell {
  packages = [
    goEnv
    gomod2nix
    pkgs.golangci-lint
    pkgs.gotools
    pkgs.go-junit-report
    pkgs.go-task
    pkgs.delve
  ];
  inherit (pre-commit-check) shellHook;
}
