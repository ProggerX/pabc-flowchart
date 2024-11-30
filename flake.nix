{
    inputs = {
        nixpkgs.url = "github:nixos/nixpkgs";
        flake-parts.url = "github:hercules-ci/flake-parts";
    };
    outputs = { nixpkgs, flake-parts, ... }@inputs: flake-parts.lib.mkFlake { inherit inputs; } {
        systems = nixpkgs.lib.platforms.unix;
        perSystem = { pkgs, lib, ... }: {
			packages.default = pkgs.buildGoModule {
				name = "pabc-flowchart";
				src = ./.;
				vendorHash = null;
				meta = with lib; {
					description = "Simple PascalABC.NET -> Mermaid.js script";
					homepage = "https://github.com/ProggerX/pabc-flowchart";
					license = licenses.gpl3Only;
					mainProgram = "pabc-parse";
				};
			};
			devShells.default = pkgs.mkShell {
				packages = [ pkgs.go ];
			};
        };
    };
}
