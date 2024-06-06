package indexerbase

type Decoder interface {
	ExtractModuleDecoder(moduleName string, module any) ModuleStateDecoder
}

type ModuleStateDecoder interface {
	GetSchema() Schema
	DecodeSet(key, value []byte) (EntityUpdate, bool, error)
	DecodeDelete(key []byte) (EntityDelete, bool, error)
}