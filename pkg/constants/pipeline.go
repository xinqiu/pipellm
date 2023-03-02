package constants

const (
	// ChainTypeMapReduce = 0
	ChainTypeMapReduce ChainType = iota
)

// ChainType represents chain type
type ChainType int8

// SummarizeMapReducePromptInputVars input variables of map reduce summarize
var SummarizeMapReducePromptInputVars = []string{"text"}

// SummarizeMapReducePromptTemplate prompt template of map reduce summarize
const SummarizeMapReducePromptTemplate = "Write a concise summary of the following:\n\n\n\"{{.text}}\"\n\n\nCONCISE SUMMARY:"
