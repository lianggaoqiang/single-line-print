package single_line_print

// NewWriter returns a new instance of writer
func NewWriter() *writer {
	w := writer(defaultIns("writer"))
	return &w
}

// NewWriterWithFlag returns a new instance of writer with flag
func NewWriterWithFlag(f uint8) *writer {
	w := writer(defaultInsWithFlag("writer", f))
	return &w
}

// Write has implemented the io.Writer interface
func (w *writer) Write(p []byte) (n int, err error) {
	pt := printerPointer(w)
	return pt.Print(string(p))
}

// Reload resets the state of an existing printer or writer instance
func (w *writer) Reload() {
	insPointer(w).Reload()
}

// Stop will set printer or writer to closed state, and eliminate the impact of flags on terminal
func (w *writer) Stop() {
	insPointer(w).Stop()
}
