package types

const (
	ErrOk = iota
	ErrGeneral
	ErrEOF
	ErrInvalidEmailType
)

var (
	ErrorTable = map[int]string{
		ErrOk:               "ok",
		ErrGeneral:          "error",
		ErrEOF:              "eof error, no input data",
		ErrInvalidEmailType: "invalid email type",
	}
)
