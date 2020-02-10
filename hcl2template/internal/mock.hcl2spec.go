// Code generated by "mapstructure-to-hcl2 -type MockConfig,NestedMockConfig"; DO NOT EDIT.
package hcl2template

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatMockConfig is an auto-generated flat version of MockConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatMockConfig struct {
	NotSquashed          *string                `mapstructure:"not_squashed" cty:"not_squashed"`
	String               *string                `mapstructure:"string" cty:"string"`
	Int                  *int                   `mapstructure:"int" cty:"int"`
	Int64                *int64                 `mapstructure:"int64" cty:"int64"`
	Bool                 *bool                  `mapstructure:"bool" cty:"bool"`
	Trilean              *bool                  `mapstructure:"trilean" cty:"trilean"`
	Duration             *string                `mapstructure:"duration" cty:"duration"`
	MapStringString      map[string]string      `mapstructure:"map_string_string" cty:"map_string_string"`
	SliceString          []string               `mapstructure:"slice_string" cty:"slice_string"`
	SliceSliceString     [][]string             `mapstructure:"slice_slice_string" cty:"slice_slice_string"`
	NamedMapStringString NamedMapStringString   `mapstructure:"named_map_string_string" cty:"named_map_string_string"`
	NamedString          *NamedString           `mapstructure:"named_string" cty:"named_string"`
	Nested               *FlatNestedMockConfig  `mapstructure:"nested" cty:"nested"`
	NestedSlice          []FlatNestedMockConfig `mapstructure:"nested_slice" cty:"nested_slice"`
}

// FlatMapstructure returns a new FlatMockConfig.
// FlatMockConfig is an auto-generated flat version of MockConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*MockConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatMockConfig)
}

// HCL2Spec returns the hcl spec of a MockConfig.
// This spec is used by HCL to read the fields of MockConfig.
// The decoded values from this spec will then be applied to a FlatMockConfig.
func (*FlatMockConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"not_squashed":            &hcldec.AttrSpec{Name: "not_squashed", Type: cty.String, Required: false},
		"string":                  &hcldec.AttrSpec{Name: "string", Type: cty.String, Required: false},
		"int":                     &hcldec.AttrSpec{Name: "int", Type: cty.Number, Required: false},
		"int64":                   &hcldec.AttrSpec{Name: "int64", Type: cty.Number, Required: false},
		"bool":                    &hcldec.AttrSpec{Name: "bool", Type: cty.Bool, Required: false},
		"trilean":                 &hcldec.AttrSpec{Name: "trilean", Type: cty.Bool, Required: false},
		"duration":                &hcldec.AttrSpec{Name: "duration", Type: cty.String, Required: false},
		"map_string_string":       &hcldec.BlockAttrsSpec{TypeName: "map_string_string", ElementType: cty.String, Required: false},
		"slice_string":            &hcldec.AttrSpec{Name: "slice_string", Type: cty.List(cty.String), Required: false},
		"slice_slice_string":      &hcldec.AttrSpec{Name: "slice_slice_string", Type: cty.List(cty.List(cty.String)), Required: false},
		"named_map_string_string": &hcldec.BlockAttrsSpec{TypeName: "named_map_string_string", ElementType: cty.String, Required: false},
		"named_string":            &hcldec.AttrSpec{Name: "named_string", Type: cty.String, Required: false},
		"nested":                  &hcldec.BlockSpec{TypeName: "nested", Nested: hcldec.ObjectSpec((*FlatNestedMockConfig)(nil).HCL2Spec())},
		"nested_slice":            &hcldec.BlockListSpec{TypeName: "nested_slice", Nested: hcldec.ObjectSpec((*FlatNestedMockConfig)(nil).HCL2Spec())},
	}
	return s
}

// FlatNestedMockConfig is an auto-generated flat version of NestedMockConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatNestedMockConfig struct {
	String               *string              `mapstructure:"string" cty:"string"`
	Int                  *int                 `mapstructure:"int" cty:"int"`
	Int64                *int64               `mapstructure:"int64" cty:"int64"`
	Bool                 *bool                `mapstructure:"bool" cty:"bool"`
	Trilean              *bool                `mapstructure:"trilean" cty:"trilean"`
	Duration             *string              `mapstructure:"duration" cty:"duration"`
	MapStringString      map[string]string    `mapstructure:"map_string_string" cty:"map_string_string"`
	SliceString          []string             `mapstructure:"slice_string" cty:"slice_string"`
	SliceSliceString     [][]string           `mapstructure:"slice_slice_string" cty:"slice_slice_string"`
	NamedMapStringString NamedMapStringString `mapstructure:"named_map_string_string" cty:"named_map_string_string"`
	NamedString          *NamedString         `mapstructure:"named_string" cty:"named_string"`
}

// FlatMapstructure returns a new FlatNestedMockConfig.
// FlatNestedMockConfig is an auto-generated flat version of NestedMockConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*NestedMockConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatNestedMockConfig)
}

// HCL2Spec returns the hcl spec of a NestedMockConfig.
// This spec is used by HCL to read the fields of NestedMockConfig.
// The decoded values from this spec will then be applied to a FlatNestedMockConfig.
func (*FlatNestedMockConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"string":                  &hcldec.AttrSpec{Name: "string", Type: cty.String, Required: false},
		"int":                     &hcldec.AttrSpec{Name: "int", Type: cty.Number, Required: false},
		"int64":                   &hcldec.AttrSpec{Name: "int64", Type: cty.Number, Required: false},
		"bool":                    &hcldec.AttrSpec{Name: "bool", Type: cty.Bool, Required: false},
		"trilean":                 &hcldec.AttrSpec{Name: "trilean", Type: cty.Bool, Required: false},
		"duration":                &hcldec.AttrSpec{Name: "duration", Type: cty.String, Required: false},
		"map_string_string":       &hcldec.BlockAttrsSpec{TypeName: "map_string_string", ElementType: cty.String, Required: false},
		"slice_string":            &hcldec.AttrSpec{Name: "slice_string", Type: cty.List(cty.String), Required: false},
		"slice_slice_string":      &hcldec.AttrSpec{Name: "slice_slice_string", Type: cty.List(cty.List(cty.String)), Required: false},
		"named_map_string_string": &hcldec.BlockAttrsSpec{TypeName: "named_map_string_string", ElementType: cty.String, Required: false},
		"named_string":            &hcldec.AttrSpec{Name: "named_string", Type: cty.String, Required: false},
	}
	return s
}
