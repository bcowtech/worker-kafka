package internal

var _ NameTransformProc = NopNameTransformer

func NopNameTransformer(name string) string { return name }
