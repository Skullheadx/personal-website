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
          curl
          fish
          (buildGoModule rec {
            pname = "pet";
            version = "0.3.4";

            src = fetchFromGitHub {
              owner = "knqyf263";
              repo = "pet";
              tag = "v${version}";
              hash = "sha256-Gjw1dRrgM8D3G7v6WIM2+50r4HmTXvx0Xxme2fH9TlQ=";
            };

            vendorHash = "sha256-6hCgv2/8UIRHw1kCe3nLkxF23zE/7t5RDwEjSzX3pBQ=";
          })
          # (pkgs.buildGoModule rec {
          #   pname = "gin-proxy";
          #   version = "unstable-2024-12-01";
          #   src = pkgs.fetchFromGitHub {
          #     owner = "codegangsta";
          #     repo = "gin";
          #     rev = "master";
          #     sha256 = "sha256-TZSpCcFBlXLPE50bYbXPU5ddoVsBG7YGa7oLmKDFBmE="; # ← you will update this
          #   };
          #   vendorHash = null;
          #   doCheck = false;
          #   modRoot = "/home/andrew/dev/personal-site/<D-;>";
          # })
        ];

        shellHook = ''
          echo "Personal Website Project"
          exec fish
        '';
      };

    };
}
