package main

import "fmt"

type Filter interface {
	execute(request string)
}

type AuthenticationFilter struct {
}

func (ae AuthenticationFilter) execute(request string) {
	fmt.Printf("Authnticating Request : %v\n", request)
}

type DebugFilter struct {
}

func (ae DebugFilter) execute(request string) {
	fmt.Printf("DebugFilter  Request : %v\n", request)
}

type Target struct {
}

func (ae Target) execute(request string) {
	fmt.Printf("Executing request: : %v\n", request)
}

type FilterChain struct {
	filters []Filter
	target  Target
}

func (filterChain *FilterChain) addFilter(filter Filter) {
	filterChain.filters = append(filterChain.filters, filter)
}

func (filterChain *FilterChain) execute(request string) {
	for _, filter := range filterChain.filters {
		filter.execute(request)
	}
	filterChain.target.execute(request)
}
func (filterChain *FilterChain) setTarget(target Target) {
	filterChain.target = target
}

type FilterManager struct {
	filterChain FilterChain
}

func NewFilterManager(target Target) *FilterManager {
	return &FilterManager{filterChain: FilterChain{target: target}}
}
func (fm *FilterManager) setFilter(filter Filter) {
	fm.filterChain.addFilter(filter)
}
func (fm *FilterManager) filterRequest(request string) {
	fm.filterChain.execute(request)
}

type Client struct {
	fm FilterManager
}

func (cl *Client) setFilterManager(filterManager FilterManager) {
	cl.fm = filterManager
}

func (cl *Client) sendRequest( request string){
	cl.fm.filterRequest(request)
}
