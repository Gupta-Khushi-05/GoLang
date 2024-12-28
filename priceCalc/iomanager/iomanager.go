package iomanager

type IOManager interface {
	WriteJsonData(data any) error
	ReadData() ([]string, error)
}