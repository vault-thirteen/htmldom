// child_node.go.

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

// Searches for the nearest Child Node having the specified Tag Name.
func GetChildNodeByTag(
	startingPoint *html.Node,
	tagName string,
) (childNode *html.Node) {
	if startingPoint == nil {
		return nil
	}
	childNode = startingPoint.FirstChild
	if childNode == nil {
		return nil
	}
	for {
		if childNode.Data == tagName {
			return
		} else {
			childNode = childNode.NextSibling
			if childNode == nil {
				return nil
			}
		}
	}
}

// Searches for the nearest Child Node having the specified Tag Name and Class.
func GetChildNodeByTagAndClass(
	startingPoint *html.Node,
	tagName string,
	className string,
) (childNode *html.Node) {
	if startingPoint == nil {
		return nil
	}
	childNode = startingPoint.FirstChild
	if childNode == nil {
		return nil
	}
	for {
		var classValue string
		var classExists bool
		classValue, classExists = GetNodeAttributeValue(childNode, AttributeClass)
		if (childNode.Data == tagName) &&
			(classExists) &&
			(classValue == className) {
			return
		} else {
			childNode = childNode.NextSibling
			if childNode == nil {
				return nil
			}
		}
	}
}

// Searches for the nearest Child Node having the specified Tag Name and Id.
func GetChildNodeByTagAndId(
	startingPoint *html.Node,
	tagName string,
	idName string,
) (childNode *html.Node) {
	if startingPoint == nil {
		return nil
	}
	childNode = startingPoint.FirstChild
	if childNode == nil {
		return nil
	}
	for {
		var idValue string
		var idExists bool
		idValue, idExists = GetNodeAttributeValue(childNode, AttributeId)
		if (childNode.Data == tagName) &&
			(idExists) &&
			(idValue == idName) {
			return
		} else {
			childNode = childNode.NextSibling
			if childNode == nil {
				return nil
			}
		}
	}
}
