package tree_sitter_metasp_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_metasp "github.com/potassco/tree-sitter-metasp/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_metasp.Language())
	if language == nil {
		t.Errorf("Error loading Metasp grammar")
	}
}
