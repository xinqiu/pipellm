package prompts

import (
	"bytes"
	"context"
	"log"
	tpl "text/template"
)

// PromptTemplate prompt implementation
type PromptTemplate struct {
	Prompt
	// inputVariables A list of the names of the variables the prompt template expects.
	inputVariables []string
	// template The prompt template.
	template string
	// tpl template engine
	tpl *tpl.Template
}

// Name returns name of prompt implementation
func (receiver *PromptTemplate) Name() string {
	return "PromptTemplate"
}

// New returns a prompt template
func New(ctx context.Context, inputVariables []string, template string) *PromptTemplate {
	t, err := tpl.New("f-string").Parse(template)
	if err != nil {
		log.Fatal(err)
	}

	return &PromptTemplate{
		inputVariables: inputVariables,
		template:       template,
		tpl:            t,
	}
}

// Format returns formatted template
// 		Args:
//            inputVariables: Any arguments to be passed to the prompt template.
//		Returns:
//            A formatted string.
//      Example:
//      .. code-block:: golang
//          prompts.New(ctx, []string{"adjective", "content"}, "Tell me a {{.adjective}} joke about {{.content}}.")
//      """
func (receiver *PromptTemplate) Format(ctx context.Context, inputVars map[string]string) string {
	if !checkValidTemplate() {
		return ""
	}
	buf := &bytes.Buffer{}
	if err := receiver.tpl.Execute(buf, inputVars); err != nil {
		log.Fatal(err)
		return ""
	}
	return buf.String()
}

// InputVars returns inputVariables
func (receiver *PromptTemplate) InputVars(ctx context.Context) []string {
	return receiver.inputVariables
}

// TODO
func checkValidTemplate() bool {
	return true
}
