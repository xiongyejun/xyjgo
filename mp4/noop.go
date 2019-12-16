package mp4

import (
	"io"
)

type noopFilter struct{}

// Noop returns a filter that does nothing
func Noop() Filter {
	return &noopFilter{}
}

func (f *noopFilter) FilterMoov(m *MoovBox) error {
	return nil
}

func (f *noopFilter) FilterMdat(w io.Writer, m *MdatBox) error {
	err := EncodeHeader(m, w)
	if err == nil {
		_, err = io.Copy(w, m.Reader())
	}
	return err
}
