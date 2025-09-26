package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/hugolib/doctree"
)

func TestcreateMutableTrees_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	treePages := &doctree.NodeShiftTree[contentNodeI]{}
	treeResources := &doctree.NodeShiftTree[contentNodeI]{}
	treePagesResources := doctree.WalkableTrees[contentNodeI]{
		treePages,
		treeResources,
	}

	resourceTrees := doctree.MutableTrees{
		treeResources,
	}

	pageTrees := &pageTrees{
		treePages:                     treePages,
		treeResources:                 treeResources,
		treePagesResources:            treePagesResources,
		treeTaxonomyEntries:           nil,
		treePagesFromTemplateAdapters: nil,
		resourceTrees:                 resourceTrees,
	}

	pageTrees.createMutableTrees()

	if !reflect.DeepEqual(pageTrees.treePagesResources, treePagesResources) {
		t.Errorf("Expected treePagesResources to be %v, but got %v", treePagesResources, pageTrees.treePagesResources)
	}

	if !reflect.DeepEqual(pageTrees.resourceTrees, resourceTrees) {
		t.Errorf("Expected resourceTrees to be %v, but got %v", resourceTrees, pageTrees.resourceTrees)
	}
}
