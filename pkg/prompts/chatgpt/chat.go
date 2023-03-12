package chatgpt

import (
	"bytes"
	"context"
	"log"
	tpl "text/template"

	"github.com/index-labs/pipelang/pkg/llms/chatgpt"
	"github.com/index-labs/pipelang/pkg/prompts"
)

// ChatPromptTemplate prompt implementation
type ChatPromptTemplate struct {
	prompts.Prompt
	// role
	role string
	// inputVariables A list of the names of the variables the prompt template expects.
	inputVariables []string
	// template The prompt template.
	template string
	// tpl template engine
	tpl *tpl.Template
}

// Name returns name of prompt implementation
func (receiver *ChatPromptTemplate) Name() string {
	return "ChatPromptTemplate"
}

// New returns a prompt template
func New(ctx context.Context, role string, inputVariables []string, template string) *ChatPromptTemplate {
	t, err := tpl.New("f-string").Parse(template)
	if err != nil {
		log.Fatal(err)
	}

	return &ChatPromptTemplate{
		role:           role,
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
func (receiver *ChatPromptTemplate) Format(ctx context.Context, inputVars map[string]string) chatgpt.ChatMessage {
	if !checkValidTemplate() {
		return chatgpt.ChatMessage{}
	}
	buf := &bytes.Buffer{}
	if err := receiver.tpl.Execute(buf, inputVars); err != nil {
		log.Fatal(err)
		return chatgpt.ChatMessage{}
	}
	switch receiver.role {
	case chatgpt.ChatMessageRoleSystem:
	case chatgpt.ChatMessageRoleUser:
		return chatgpt.HumanMessage(buf.String())
	case chatgpt.ChatMessageRoleAssistant:
		return chatgpt.AssistantMessage(buf.String())
	}
	return chatgpt.ChatMessage{}
}

// InputVars returns inputVariables
func (receiver *ChatPromptTemplate) InputVars(ctx context.Context) []string {
	return receiver.inputVariables
}

// TODO
func checkValidTemplate() bool {
	return true
}
