package pkg

import (
	"errors"
	"fmt"
	"strings"

	"example.com/htmlParser/internal"
)

type node struct {
	name     string
	children []*node
}

func Parse() {
	var sampleHtml string = `<div> <s></s> This is a parent div </div>`
	var asRune []rune = []rune(sampleHtml)

	var isOpen = false

	var head *node
	var err error

	tagsOpened := new(internal.Stack[node])
	for i := 0; i < len(asRune); i++ {
		var current *node = tagsOpened.Peek()
		//look for opening tag
		if asRune[i] == '<' {

			if isOpen {
				err = errors.New("element tag opening found without closing previous tag")
				break
			}
			isOpen = true

			// find tag identifier
			var tagName strings.Builder

			for {
				i++
				if tagEnd(asRune[i]) {
					analyzeString(string(tagName.String()))
					var newNode, err = createTag(&tagName, tagsOpened)
					if err == nil && newNode != nil {
						if current == nil {
							head = newNode
							current = newNode

						} else {
							current.children = append(current.children, newNode)

						}
					}

					if asRune[i] == '>' {
						isOpen = false
					}
					break
				}
				tagName.WriteString(string(asRune[i]))

			}
		}

	}
	if err != nil {
		fmt.Println("Error found: ", err.Error())
	}
	fmt.Printf("The node is: %v ", head)

}

func analyzeString(element string) {
	fmt.Printf("found element %v", element)
}

func tagEnd(character rune) bool {
	return character == ' ' || character == '>'
}

func matchingCloseTag(tagName string, tagsOpen *internal.Stack[node]) bool {

	lastTag := tagsOpen.Peek().name
	return strings.HasSuffix(tagName, lastTag)

}

func createTag(builder *strings.Builder, tagsOpen *internal.Stack[node]) (*node, error) {

	var err error
	var tagName string = builder.String()
	tag := node{tagName, make([]*node, 0)}
	if tagName[0] == '/' {
		var ok = matchingCloseTag(tagName, tagsOpen)
		if !ok {
			err = errors.New("mismatch closing tag")

		}
		// close tag
		tagsOpen.Pop()

		return nil, err

	}

	tag.name = tagName
	tagsOpen.Push(&tag)
	return &tag, err

}
