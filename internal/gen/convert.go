package gen

import (
	"log/slog"

	"github.com/handlename/protoc-gen-oas2connect/internal/oas"
	"github.com/handlename/protoc-gen-oas2connect/internal/proto"
)

func toOasDataType(fieldType proto.FieldType) oas.DataType {
	switch fieldType {
	case proto.FieldTypeDouble, proto.FieldTypeFloat:
		return oas.DataTypeNumber
	case proto.FieldTypeInt32, proto.FieldTypeInt64, proto.FieldTypeUint32, proto.FieldTypeUint64:
		return oas.DataTypeInteger
	case proto.FieldTypeBool:
		return oas.DataTypeBoolean
	case proto.FieldTypeString:
		return oas.DataTypeString
	case proto.FieldTypeBytes:
		return oas.DataTypeString
	default:
		slog.Warn("unknown FieldType, falled back to string", slog.String("type", string(fieldType)))
		return oas.DataTypeString
	}
}
