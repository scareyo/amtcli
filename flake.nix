{
  description = "A command-line Intel AMT tool";

  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  inputs.flake-parts.url = "github:hercules-ci/flake-parts";
  inputs.devenv.url = "github:cachix/devenv";

  outputs = inputs@{ self, devenv, flake-parts, nixpkgs }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [
        devenv.flakeModule
      ];

      systems = [
        "x86_64-linux"
        "aarch64-darwin"
      ];

      perSystem = { lib, pkgs, system, ... }: {
        devenv.shells.default = {
          languages = {
            go.enable = true;
          };
        };
        packages.default = pkgs.buildGoModule {
          pname = "amtcli";
          version = "0.1.0";

          src = pkgs.fetchFromGitHub {
            owner = "scareyo";
            repo = "amtcli";
            rev = "main";
            sha256 = "01hhy3prm8jbyam1n6qw1nl4ra5lghgafr9cd2wr3568ncdq8369";
          };

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
