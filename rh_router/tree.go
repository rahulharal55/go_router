package rhrouter

import (
	"net/http"
	"strings"
)

const (
	pathRoot      string = "/"
	pathDelimiter string = "/"
)

type Tree struct {
	Node *Node
}

type Node struct {
	label    string
	actions  map[string]*action
	children map[string]*Node
}

type action struct {
	handler http.Handler
}

type result struct {
	actions *action
}

func NewResult() *result {
	return &result{}
}

func NewTree() *Tree {
	return &Tree{
		Node: &Node{
			label:    pathRoot,
			actions:  make(map[string]*action),
			children: make(map[string]*Node),
		},
	}
}

func (t *Tree) Insert(methods []string, path string, handler http.Handler) error {

	curNode := t.Node
	if path == pathRoot {
		curNode.label = path
		addActionHandlers(methods, curNode, handler)
		return nil
	}
	//get all path elements
	elements := splitPath(path)

	// iterate all elements
	for i, r := range elements {
		// if
		//		path is alreay in children then replace currNode with that child and continue
		//		path is missing then create new node and add that to childrens map and replace currNode with new node
		if nextNode, ok := curNode.children[r]; ok {
			curNode = nextNode
		} else {
			newNode := &Node{
				label:    r,
				actions:  map[string]*action{},
				children: map[string]*Node{},
			}
			curNode.children[r] = newNode
			curNode = newNode
		}
		if i == len(elements)-1 {
			curNode.label = r
			addActionHandlers(methods, curNode, handler)
			break
		}
	}
	return nil
}

func addActionHandlers(methods []string, curNode *Node, handler http.Handler) {
	for _, method := range methods {
		curNode.actions[method] = &action{
			handler: handler,
		}
	}
}

// Search is search a word from a tree.
func (t *Tree) Search(method string, path string) (*result, error) {
	result := NewResult()
	curNode := t.Node

	if path != pathRoot {
		for _, r := range splitPath(path) {
			if nextNode, ok := curNode.children[r]; ok {
				curNode = nextNode
			} else {
				if r == curNode.label {
					break
				} else {
					return nil, ErrNotFound
				}
			}
		}
	}

	result.actions = curNode.actions[method]
	if result.actions == nil {
		// no matching handler was found.
		return nil, ErrMethodNotAllowed
	}
	return result, nil
}

func splitPath(path string) []string {
	s := strings.Split(path, pathRoot)
	var r []string
	for _, v := range s {
		if v != "" {
			r = append(r, v)
		}
	}
	return r
}
