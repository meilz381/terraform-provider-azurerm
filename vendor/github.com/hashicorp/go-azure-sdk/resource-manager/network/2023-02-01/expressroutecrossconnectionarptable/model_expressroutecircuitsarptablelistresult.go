package expressroutecrossconnectionarptable

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressRouteCircuitsArpTableListResult struct {
	NextLink *string                        `json:"nextLink,omitempty"`
	Value    *[]ExpressRouteCircuitArpTable `json:"value,omitempty"`
}
