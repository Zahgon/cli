package cli

type numberType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func getNumberSlice[T numberType](cmd *Command, name string) []T {
	_ = "STUB: not implemented"
	return nil
}
