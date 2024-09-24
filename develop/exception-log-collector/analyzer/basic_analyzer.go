package analyzer

import (
    "regexp"
)

var (
    errorRegex = regexp.MustCompile(`\b(ERROR|WARN|FATAL)\b`)
)

type BasicAnalyzer struct {
    output chan string
}

func NewBasicAnalyzer(output chan string) *BasicAnalyzer {
    return &BasicAnalyzer{
        output: output,
    }
}

func (ba *BasicAnalyzer) Analyze(logLine string) {
    if errorRegex.MatchString(logLine) {
        ba.output <- logLine
    }
}
