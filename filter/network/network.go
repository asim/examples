// Package network provides a way of filrering calls to routes on a given network
package network

import (
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/router"
	"github.com/micro/go-micro/v2/selector"
)

// Filter will filter the network
func Filter(n string) client.CallOption {
	filter := func(routes []router.Route) []router.Route {
		var filtered []router.Route

		for _, route := range routes {
			if route.Network == n {
				filtered = append(filtered, route)
			}
		}

		return filtered
	}

	return client.WithSelectOptions(selector.WithFilter(filter))
}
