package proto

import (
	"log/slog"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// https://protobuf.dev/programming-guides/proto3/#scalar
type FieldType string

const (
	FieldTypeDouble   FieldType = "double"
	FieldTypeFloat    FieldType = "float"
	FieldTypeInt32    FieldType = "int32"
	FieldTypeInt64    FieldType = "int64"
	FieldTypeUint32   FieldType = "uint32"
	FieldTypeUint64   FieldType = "uint64"
	FieldTypeSint32   FieldType = "sint32"
	FieldTypeSint64   FieldType = "sint64"
	FieldTypeFixed32  FieldType = "fixed32"
	FieldTypeFixed64  FieldType = "fixed64"
	FieldTypeSfixed32 FieldType = "sfixed32"
	FieldTypeSfixed64 FieldType = "sfixed64"
	FieldTypeBool     FieldType = "bool"
	FieldTypeString   FieldType = "string"
	FieldTypeBytes    FieldType = "bytes"
)

func NewFieldType(kind protoreflect.Kind) FieldType {
	switch kind {
	case protoreflect.DoubleKind:
		return FieldTypeDouble
	case protoreflect.FloatKind:
		return FieldTypeFloat
	case protoreflect.Int32Kind:
		return FieldTypeInt32
	case protoreflect.Int64Kind:
		return FieldTypeInt64
	case protoreflect.Uint32Kind:
		return FieldTypeUint32
	case protoreflect.Uint64Kind:
		return FieldTypeUint64
	case protoreflect.Sint32Kind:
		return FieldTypeSint32
	case protoreflect.Sint64Kind:
		return FieldTypeSint64
	case protoreflect.Fixed32Kind:
		return FieldTypeFixed32
	case protoreflect.Fixed64Kind:
		return FieldTypeFixed64
	case protoreflect.Sfixed32Kind:
		return FieldTypeSfixed32
	case protoreflect.Sfixed64Kind:
		return FieldTypeSfixed64
	case protoreflect.BoolKind:
		return FieldTypeBool
	case protoreflect.StringKind:
		return FieldTypeString
	case protoreflect.BytesKind:
		return FieldTypeBytes
	default:
		slog.Warn("unknown FieldType, falled back to string", slog.String("kind", kind.String()))
		return FieldTypeString
	}
}

// ToGoTypeName returns Go type name from FieldType,
// based on https://protobuf.dev/programming-guides/proto3/#scalar
func (t FieldType) ToGoTypeName() string {
	switch t {
	case "double":
		return "float64"
	case "float":
		return "float32"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "uint32", "sint32", "fixed32", "sfixed32":
		return "uint32"
	case "uint64", "sint64", "fixed64", "sfixed64":
		return "uint64"
	case "bool":
		return "bool"
	case "string":
		return "string"
	case "bytes":
		return "[]byte"
	default:
		slog.Warn("unknown FieldType, falled back to string", slog.String("type", string(t)))
		return "string"
	}
}
