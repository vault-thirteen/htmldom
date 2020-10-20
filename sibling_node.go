// sibling_node.go.

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
	"golang.org/x/net/html"
)

// Searches for the nearest Sibling Node having the specified Tag Name.
func GetSiblingNodeByTag(
	startingPoint *html.Node,
	tagName string,
) (siblingNode *html.Node) {
	if startingPoint == nil {
		return nil
	}
	siblingNode = startingPoint.NextSibling
	if siblingNode == nil {
		return nil
	}
	for {
		if siblingNode.Data == tagName {
			return
		} else {
			siblingNode = siblingNode.NextSibling
			if siblingNode == nil {
				return nil
			}
		}
	}
}

// Searches for the nearest Sibling Node having the specified Tag Name and Class.
func GetSiblingNodeByTagAndClass(
	startingPoint *html.Node,
	tagName string,
	className string,
) (siblingNode *html.Node) {
	if startingPoint == nil {
		return nil
	}
	siblingNode = startingPoint.NextSibling
	if siblingNode == nil {
		return nil
	}
	for {
		var classValue string
		var classExists bool
		classValue, classExists = GetNodeAttributeValue(siblingNode, AttributeClass)
		if (siblingNode.Data == tagName) &&
			(classExists) &&
			(classValue == className) {
			return
		} else {
			siblingNode = siblingNode.NextSibling
			if siblingNode == nil {
				return nil
			}
		}
	}
}

// Searches for the nearest Sibling Node having the specified Tag Name and Id.
func GetSiblingNodeByTagAndId(
	startingPoint *html.Node,
	tagName string,
	idName string,
) (siblingNode *html.Node) {
	if startingPoint == nil {
		return nil
	}
	siblingNode = startingPoint.NextSibling
	if siblingNode == nil {
		return nil
	}
	for {
		var idValue string
		var idExists bool
		idValue, idExists = GetNodeAttributeValue(siblingNode, AttributeId)
		if (siblingNode.Data == tagName) &&
			(idExists) &&
			(idValue == idName) {
			return
		} else {
			siblingNode = siblingNode.NextSibling
			if siblingNode == nil {
				return nil
			}
		}
	}
}
