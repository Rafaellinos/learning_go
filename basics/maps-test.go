package main

import "fmt"

/*
key-value pairs
no order

e.g.:

myMap := make(map[string]int)

OR

myMap := map[string]int{
	"item1": 1,
	"item2": 2,
	"item3": 3,
}

Dynamic size.

# insert

myMap["newKey"] = 5

# read

value := myMap["newKey"]
missing := myMap["NonExistingKey"] // default value

# Delete

delete(myMap, "newKey")

# Check existence

price, found := myMap["price"]
if !found {
	fmt.Println("price not found")
	return
}
*/

func mapsExample() {
	myMap := map[string]string{
		"name":   "Rafael",
		"age":    "23",
		"height": "1.71cm",
	}
	for key, value := range myMap {
		fmt.Println("DataKey:", key)
		fmt.Println("DataValue", value)
	}
	age, found := myMap["age"]
	if !found {
		fmt.Println("Age not found")
	} else {
		fmt.Println("Age found", age)
	}
	delete(myMap, "age")
}

func createServer(serverName, serverStatus string) map[string]string {
	return map[string]string{
		"name":   serverName,
		"status": serverStatus,
	}
}

func exerciseMaps() {
	possibleStatus := []string{
		"Online", "Offline", "Maintenance", "Retired",
	}
	servers := []map[string]string{
		createServer("Darkstar", possibleStatus[0]),
		createServer("Aiur", possibleStatus[0]),
		createServer("us-east-1", possibleStatus[0]),
	}
	fmt.Println("number of servers:", len(servers))
	var serversOnline int
	var serversOffline int
	var serversMaintenance int
	var serversRetired int
	for _, server := range servers {
		fmt.Println(server)
		switch status := server["status"]; status {
		case possibleStatus[0]:
			serversOnline += 1
		case possibleStatus[1]:
			serversOffline += 1
		case possibleStatus[2]:
			serversMaintenance += 1
		case possibleStatus[3]:
			serversRetired += 1
		default:
			panic("Status not found")
		}
	}
	fmt.Println("online servers:", serversOnline)
	fmt.Println("offline servers:", serversOffline)
	fmt.Println("maintenance servers:", serversMaintenance)
	fmt.Println("retired servers:", serversRetired)
}
