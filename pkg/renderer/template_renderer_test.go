package renderer_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/Ryooooooga/zouch/pkg/renderer"
	"github.com/Ryooooooga/zouch/pkg/repositories"
	"github.com/stretchr/testify/assert"
)

func newTestTextTemplateRenderer() *renderer.TextTemplateRenderer {
	r := renderer.NewTextTemplateRenderer()
	r.FuncMap["Now"] = func() time.Time {
		return time.Date(2021, 4, 30, 20, 45, 30, 0, time.Local)
	}

	return r
}

func TestRenderTemplate(t *testing.T) {
	scenarios := []struct {
		testname       string
		templateText   string
		data           interface{}
		expectedOutput string
	}{
		{
			testname:       "call-now",
			templateText:   `Today is {{ Now.Format "2006-01-02" }}`,
			data:           nil,
			expectedOutput: "Today is 2021-04-30",
		},
		{
			testname:       "variable",
			templateText:   `data.foo == {{ .foo }}`,
			data:           map[string]string{"foo": "BAR"},
			expectedOutput: "data.foo == BAR",
		},
		{
			testname:       "string",
			templateText:   `{{ . | UpperSnakeCase }}`,
			data:           "Hello world!",
			expectedOutput: "HELLO_WORLD!",
		},
	}

	for _, s := range scenarios {
		t.Run(s.testname, func(t *testing.T) {
			r := newTestTextTemplateRenderer()
			output := bytes.NewBufferString("")

			tpl := &repositories.TemplateFile{
				Path:    s.testname,
				Content: []byte(s.templateText),
			}

			err := r.RenderTemplate(output, tpl, s.data)
			assert.Nil(t, err)
			assert.Equal(t, s.expectedOutput, output.String())
		})
	}
}
