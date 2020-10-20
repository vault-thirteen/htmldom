// node_attribute_test.go.

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
	"testing"

	"github.com/vault-thirteen/tester"
	"golang.org/x/net/html"
)

func Test_NodeHasAttribute(t *testing.T) {

	var aTest *tester.Test = tester.New(t)
	var node *html.Node
	var result bool

	// Test #1. Null Node.
	node = nil
	result = NodeHasAttribute(node, "class")
	aTest.MustBeEqual(result, false)

	// Test #2. Negative.
	node = &html.Node{
		Attr: []html.Attribute{
			{
				Key: "John",
				Val: "Doe",
			},
		},
	}
	result = NodeHasAttribute(node, "class")
	aTest.MustBeEqual(result, false)

	// Test #3. Positive.
	node = &html.Node{
		Attr: []html.Attribute{
			{
				Key: "John",
				Val: "Doe",
			},
			{
				Key: "class",
				Val: "the-class",
			},
		},
	}
	result = NodeHasAttribute(node, "class")
	aTest.MustBeEqual(result, true)
}

func Test_GetNodeAttributeValue(t *testing.T) {

	var aTest *tester.Test = tester.New(t)
	var node *html.Node
	var attributeValue string
	var attributeExists bool

	// Test #1. Null Node.
	node = nil
	attributeValue, attributeExists = GetNodeAttributeValue(node, "class")
	aTest.MustBeEqual(attributeValue, "")
	aTest.MustBeEqual(attributeExists, false)

	// Test #2. Negative.
	node = &html.Node{
		Attr: []html.Attribute{
			{
				Key: "John",
				Val: "Doe",
			},
		},
	}
	attributeValue, attributeExists = GetNodeAttributeValue(node, "class")
	aTest.MustBeEqual(attributeValue, "")
	aTest.MustBeEqual(attributeExists, false)

	// Test #3. Positive.
	node = &html.Node{
		Attr: []html.Attribute{
			{
				Key: "John",
				Val: "Doe",
			},
			{
				Key: "class",
				Val: "the-class",
			},
		},
	}
	attributeValue, attributeExists = GetNodeAttributeValue(node, "class")
	aTest.MustBeEqual(attributeValue, "the-class")
	aTest.MustBeEqual(attributeExists, true)
}
