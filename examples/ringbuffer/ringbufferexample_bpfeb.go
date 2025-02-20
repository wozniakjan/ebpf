// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64be || armbe || mips || mips64 || mips64p32 || ppc64 || s390 || s390x || sparc || sparc64
// +build arm64be armbe mips mips64 mips64p32 ppc64 s390 s390x sparc sparc64

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// LoadRingBufferExample returns the embedded CollectionSpec for RingBufferExample.
func LoadRingBufferExample() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_RingBufferExampleBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load RingBufferExample: %w", err)
	}

	return spec, err
}

// LoadRingBufferExampleObjects loads RingBufferExample and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//     *RingBufferExampleObjects
//     *RingBufferExamplePrograms
//     *RingBufferExampleMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadRingBufferExampleObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadRingBufferExample()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// RingBufferExampleSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type RingBufferExampleSpecs struct {
	RingBufferExampleProgramSpecs
	RingBufferExampleMapSpecs
}

// RingBufferExampleSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type RingBufferExampleProgramSpecs struct {
	KprobeExecve *ebpf.ProgramSpec `ebpf:"kprobe_execve"`
}

// RingBufferExampleMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type RingBufferExampleMapSpecs struct {
	Events *ebpf.MapSpec `ebpf:"events"`
}

// RingBufferExampleObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadRingBufferExampleObjects or ebpf.CollectionSpec.LoadAndAssign.
type RingBufferExampleObjects struct {
	RingBufferExamplePrograms
	RingBufferExampleMaps
}

func (o *RingBufferExampleObjects) Close() error {
	return _RingBufferExampleClose(
		&o.RingBufferExamplePrograms,
		&o.RingBufferExampleMaps,
	)
}

// RingBufferExampleMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadRingBufferExampleObjects or ebpf.CollectionSpec.LoadAndAssign.
type RingBufferExampleMaps struct {
	Events *ebpf.Map `ebpf:"events"`
}

func (m *RingBufferExampleMaps) Close() error {
	return _RingBufferExampleClose(
		m.Events,
	)
}

// RingBufferExamplePrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadRingBufferExampleObjects or ebpf.CollectionSpec.LoadAndAssign.
type RingBufferExamplePrograms struct {
	KprobeExecve *ebpf.Program `ebpf:"kprobe_execve"`
}

func (p *RingBufferExamplePrograms) Close() error {
	return _RingBufferExampleClose(
		p.KprobeExecve,
	)
}

func _RingBufferExampleClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed ringbufferexample_bpfeb.o
var _RingBufferExampleBytes []byte
