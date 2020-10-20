// node_attribute.go.

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

func NodeHasAttribute(
	node *html.Node,
	attributeName string,
) (attributeExists bool) {
	if node == nil {
		return false
	}
	for _, attribute := range node.Attr {
		if attribute.Key == attributeName {
			return true
		}
	}
	return false
}

func GetNodeAttributeValue(
	node *html.Node,
	attributeName string,
) (attributeValue string, attributeExists bool) {
	if node == nil {
		return
	}
	for _, attribute := range node.Attr {
		if attribute.Key == attributeName {
			return attribute.Val, true
		}
	}
	return
}
