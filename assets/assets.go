package assets

import (
	"bytes"
	_ "embed"
	"io"
)

var (
	//go:embed CHN-FRENTE.pdf
	pdfCnhFront []byte

	//go:embed CHN-VERSO.pdf
	pdfCnhBack []byte

	//go:embed RG-FRENTE.pdf
	pdfRgFront []byte

	//go:embed RG-VERSO.pdf
	pdfRgBack []byte
)

type Document struct {
	Name  string
	Ext   string
	Front []byte
	Back  []byte
}

var (
	RG  *Document
	CNH *Document
)

func init() {
	RG = &Document{
		Name:  "RG",
		Ext:   "pdf",
		Front: pdfRgFront,
		Back:  pdfRgBack,
	}
	CNH = &Document{
		Name:  "CNH",
		Ext:   "pdf",
		Front: pdfCnhFront,
		Back:  pdfCnhBack,
	}
}

func (d *Document) GetAsReadSeeker() (io.ReadSeeker, io.ReadSeeker) {
	return bytes.NewReader(d.Front), bytes.NewReader(d.Back)
}
