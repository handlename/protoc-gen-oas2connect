package oas

// https://swagger.io/docs/specification/data-models/data-types/
type DataType string

const (
	DataTypeString  DataType = "string"
	DataTypeNumber  DataType = "number"
	DataTypeInteger DataType = "integer"
	DataTypeBoolean DataType = "boolean"

	// currently not supported
	// EndpointOasFieldDataTypeArray   EndpointOasFieldDataType = "array"
)

// https://swagger.io/docs/specification/data-models/data-types/
type DataFormat string

const (
	// for number type
	DataFormatDouble DataFormat = "double"
	DataFormatInt32  DataFormat = "int32"
	DataFormatInt64  DataFormat = "int64"

	// for string type
)

// https://swagger.io/docs/specification/describing-parameters/
type ParamType string

const (
	ParamTypePath  ParamType = "path"
	ParamTypeQuery ParamType = "query"

	// currently not supported
	// ParamTypeHeader ParamType = "header"
	// ParamTypeCookie ParamType = "cookie"
)
