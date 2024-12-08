{
  description = "A command-line Intel AMT tool";

  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  inputs.flake-parts.url = "github:hercules-ci/flake-parts";

  outputs = inputs@{ self, flake-parts, nixpkgs }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [
        "x86_64-linux"
        "aarch64-darwin"
      ];

      perSystem = { lib, pkgs, system, ... }: {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
          ];
        };
        packages.default = pkgs.buildGoModule {
          pname = "amtcli";
          version = "0.1.0";

          src = ./.;

          vendorHash = "sha256-s32vUZqPayI8CY/SvBzrhj600D/NkElu/d65hlP2+PM=";

          meta = {
            description = "A command-line Intel AMT tool";
            homepage = "https://github.com/scareyo/amtcli";
            license = lib.licenses.mit;
          };
        };
      };
    };
}
