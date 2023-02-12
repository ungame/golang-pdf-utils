package main

import (
	"bytes"
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/ungame/golang-pdf-utils/assets"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	merge(assets.RG)
	merge(assets.CNH)
}

func merge(doc *assets.Document) {
	start := time.Now()
	defer func() {
		log.Printf("Merge %s Finished. Elapsed: %s\n", doc.Name, time.Since(start).String())
	}()

	front, back := doc.GetAsReadSeeker()

	rs := make([]io.ReadSeeker, 2)
	rs[0] = front
	rs[1] = back

	buf := &bytes.Buffer{}
	if err := api.Merge(rs, buf, nil); err != nil {
		log.Panicf("error on merge %s front/back: %s\n", doc.Name, err)
	}

	filenameOutput := fmt.Sprintf("%s-MERGED.%s", doc.Name, doc.Ext)

	mergedFile, err := os.Create(filenameOutput)
	if err != nil {
		log.Panicf("error on create merged %s file: %s\n", doc.Name, err)
	}

	_, err = mergedFile.Write(buf.Bytes())
	if err != nil {
		log.Panicf("error on write merged %s file: %s\n", doc.Name, err)
	}
	err = mergedFile.Close()
	if err != nil {
		log.Panicf("error on close merged %s file: %s\n", doc.Name, err)
	}
}
