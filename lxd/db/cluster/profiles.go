//go:build linux && cgo && !agent

package cluster

// Code generation directives.
//
//go:generate -command mapper lxd-generate db mapper -t profiles.mapper.go
//go:generate mapper reset -i -b "//go:build linux && cgo && !agent"
//
//go:generate mapper stmt -p "github.com/lxc/lxd/lxd/db" -e profile id version=2
//
//go:generate mapper method -i -p "github.com/lxc/lxd/lxd/db" -e profile ID struct=Profile version=2
