{
  description = "Personal Website Project";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs =
    { self, nixpkgs }:
    let
      pkgs = import nixpkgs { system = "x86_64-linux"; };
    in
    {

      devShells.x86_64-linux.default = pkgs.mkShell {
        packages = with pkgs; [
          go
          fish
        ];

        shellHook = ''
          echo "Personal Website Project"
          exec fish
        '';
      };

    };
}
