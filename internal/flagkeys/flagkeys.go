// Package commonflags contains keys for all command-line flags related to openfeature CLI.
package flagkeys

const (
	// `generate` flags:
	// FlagManifestPath is the key for the flag that stores the flag manifest path.
	FlagManifestPath = "flag_manifest_path"
	// OutputPath is the key for the flag that stores the output path.
	OutputPath = "output_path"

	// `generate go` flags:
	// GoPackageName is the key for the flag that stores the Golang package name.
	GoPackageName = "package_name"
)