package main

import (
	"context"
	"time"
	_ "unsafe" // for go:linkname

	_ "github.com/DataDog/dd-trace-go/v2/civisibility"
)

// ********************************************************************************************************************
// TEST EVENT HOOKS

//go:linkname SetGlobalEventFinishHook github.com/DataDog/dd-trace-go/v2/internal/civisibility/integrations.SetGlobalEventFinishHook
func SetGlobalEventFinishHook(func([]any))

type Test interface {
	Context() context.Context
	StartTime() time.Time
	TestID() uint64
	Name() string
	SetTag(key string, value interface{})
	GetTag(key string) (interface{}, bool)
}

// ********************************************************************************************************************
// CODEOWNERS

//go:linkname GetCodeOwners github.com/DataDog/dd-trace-go/v2/internal/civisibility/utils.GetCodeOwners
func GetCodeOwners() *CodeOwners

//go:linkname code_owners_match github.com/DataDog/dd-trace-go/v2/internal/civisibility/utils.(*CodeOwners).Match
func code_owners_match(*CodeOwners, string) (*Entry, bool)

type (
	CodeOwners struct {
		Sections []*Section
	}
	Section struct {
		Name    string
		Entries []Entry
	}
	Entry struct {
		Pattern string
		Owners  []string
		Section string
	}
)

func (co *CodeOwners) Match(value string) (*Entry, bool) {
	return code_owners_match(co, value)
}
