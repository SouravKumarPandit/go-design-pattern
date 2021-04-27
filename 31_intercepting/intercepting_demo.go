package main

func main() {
	manager := NewFilterManager(Target{})
	manager.setFilter( AuthenticationFilter{})
	manager.setFilter( DebugFilter{})

	client := Client{}
	client.setFilterManager(*manager)
	client.sendRequest("Home")
}
