// aux_test.go.

////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019..2020 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Auxiliary Functions for Tests.

const SampleSourceFileName = "sample.html"

func _test_parseHtmlSourceIntoDom(
	htmlPageSource string,
) (domNode *html.Node, err error) {
	domNode, err = html.Parse(strings.NewReader(htmlPageSource))
	if err != nil {
		return
	}
	if domNode == nil {
		err = errors.New(ErrDomNodeIsNotFound)
		return
	}
	return
}

func _test_getSampleHtmlSource() (source string, err error) {
	var file *os.File
	file, err = os.Open(SampleSourceFileName)
	if err != nil {
		return
	}
	defer func() {
		var derr error
		derr = file.Close()
		if derr != nil {
			log.Println(derr)
		}
	}()
	var fileContents []byte
	fileContents, err = ioutil.ReadAll(file)
	if err != nil {
		return
	}
	source = string(fileContents)
	return
}

func _test_createSimpleDomForTests() (domNode *html.Node, err error) {
	var source string
	source, err = _test_getSampleHtmlSource()
	if err != nil {
		return
	}
	return _test_parseHtmlSourceIntoDom(source)
}

func _test_getContainerNodeFromSimpleDomForTests(
	domNode *html.Node,
) (containerNode *html.Node, err error) {

	// Fool Check.
	if domNode == nil {
		err = errors.New(ErrStartingPointIsNotSet)
		return
	}

	// Search.
	var node *html.Node = domNode.FirstChild // -> <html>
	if node == nil {
		err = errors.New(ErrDomNodeIsNotFound)
		return
	}
	node = node.FirstChild // -> <head>
	if node == nil {
		err = errors.New(ErrDomNodeIsNotFound)
		return
	}
	node = node.NextSibling.NextSibling // -> <body>
	if node == nil {
		err = errors.New(ErrDomNodeIsNotFound)
		return
	}
	node = node.FirstChild.NextSibling // -> <div class="class-container">
	if node == nil {
		err = errors.New(ErrDomNodeIsNotFound)
		return
	}
	return node, nil
}

func _test_prepareTestNode() (containerNode *html.Node, err error) {
	var domNode *html.Node
	domNode, err = _test_createSimpleDomForTests()
	if err != nil {
		return
	}
	return _test_getContainerNodeFromSimpleDomForTests(domNode)
}
