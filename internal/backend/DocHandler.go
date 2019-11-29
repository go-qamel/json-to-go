package backend

import (
	"bytes"
	"fmt"
	"strings"

	"json-to-go/internal/converter"

	"github.com/alecthomas/chroma/formatters/html"

	"github.com/alecthomas/chroma/styles"

	"github.com/alecthomas/chroma/lexers"

	"github.com/go-qamel/qamel"
)

func init() {
	RegisterQmlDocHandler("BackEnd", 1, 0, "DocHandler")
}

// DocHandler is backend for our app
type DocHandler struct {
	qamel.QmlObject

	_ func(string, bool) `slot:"convert"`

	_ func(string) `signal:"error"`
	_ func(string) `signal:"converted"`
}

func (b *DocHandler) convert(jsonValue string, inlineStruct bool) {
	// Convert to Go
	cv := converter.Converter{
		InlineStruct: inlineStruct,
	}

	result, err := cv.Convert(jsonValue)
	if err != nil {
		b.error(err.Error())
		return
	}

	result, err = highlightGoCode(result)
	if err != nil {
		b.error(err.Error())
		return
	}

	b.converted(result)
}

func highlightGoCode(code string) (string, error) {
	lexer := lexers.Get("go")
	style := styles.Get("dracula")
	formatter := html.New(html.PreventSurroundingPre(true))

	iterator, err := lexer.Tokenise(nil, code)
	if err != nil {
		return "", fmt.Errorf("failed to tokenise code: %v", err)
	}

	var buffer bytes.Buffer
	err = formatter.Format(&buffer, style, iterator)
	if err != nil {
		return "", fmt.Errorf("failed to highlight code: %v", err)
	}

	htmlContent := `<meta name="qrichtext" content="1">`
	htmlContent += buffer.String()
	htmlContent = strings.ReplaceAll(htmlContent, "\n", "\n<br>")

	return htmlContent, nil
}
