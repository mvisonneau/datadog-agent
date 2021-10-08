// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build linux
// +build linux

package cgroups

// Cgroup interface abstracts actual Cgroup implementation (v1/v2)
type Cgroup interface {
	// Identifier returns the cgroup identifier that was generated by the selected `Filter` (default: folder name)
	Identifier() string
	// GetParent returns parent Cgroup (will fail if used on root cgroup)
	GetParent() (Cgroup, error)
	// GetStats returns all cgroup statistics at once. Each call triggers a read from filesystem (no cache)
	// The given Stats object is filled with new values. If re-using object, old values are not cleared on read failure.
	GetStats(*Stats) error
	// GetCPUStats returns all cgroup statistics at once. Each call triggers a read from filesystem (no cache)
	// The given CPUStats object is filled with new values. If re-using object, old values are not cleared on read failure.
	GetCPUStats(*CPUStats) error
	// GetMemoryStats returns all cgroup statistics at once. Each call triggers a read from filesystem (no cache)
	// The given MemoryStats object is filled with new values. If re-using object, old values are not cleared on read failure.
	GetMemoryStats(*MemoryStats) error
	// GetIOStats returns all cgroup statistics at once. Each call triggers a read from filesystem (no cache)
	// The given IOStats object is filled with new values. If re-using object, old values are not cleared on read failure.
	GetIOStats(*IOStats) error
	// GetPIDStats returns all cgroup statistics at once. Each call triggers a read from filesystem (no cache)
	// The given PIDStats object is filled with new values. If re-using object, old values are not cleared on read failure.
	GetPIDStats(*PIDStats) error
}