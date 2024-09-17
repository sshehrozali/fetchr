package writer

type FileWriter interface {
	Write(data []byte) (int, error)
}