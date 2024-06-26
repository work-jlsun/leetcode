package loader

import (
	"fmt"
	"os"
	"testing"

	"github.com/unidoc/unipdf/v3/common"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

// func init() {
// 	// Make sure to load your metered License API key prior to using the library.
// 	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
// 	err := license.SetMeteredKey("1234")
// 	if err != nil {
// 		panic(err)
// 	}
// }

func Test(t *testing.T) {
	f, err := os.Open("/Users/tomsun/Downloads/cv_ljh_updated.pdf")
	if err != nil {
		common.Log.Error("Failed to open PDF file: %v", err)
		return
	}
	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		panic(err)
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		panic(err)
	}

	fmt.Printf("--------------------\n")
	fmt.Printf("PDF to text extraction:\n")
	fmt.Printf("--------------------\n")
	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			panic(err)
		}

		ex, err := extractor.New(page)
		if err != nil {
			panic(err)
		}

		text, err := ex.ExtractText()
		if err != nil {
			panic(err)
		}

		fmt.Println("------------------------------")
		fmt.Printf("Page %d:\n", pageNum)
		fmt.Printf("\"%s\"\n", text)
		fmt.Println("------------------------------")
	}
}
