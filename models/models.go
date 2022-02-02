package models

type ChildInput struct {
	Dummy *string
}

type ChildInputSliceAlias []ChildInput

// SampleInput is a struct that should be bind to an Input
type SampleInput struct {
	Ok    *bool                `json:"ok"`
	NotOk ChildInputSliceAlias `json:"notOk"`
}
