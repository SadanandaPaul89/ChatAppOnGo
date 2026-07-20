package main

func broadcastRoomSystem(Sender *Client, message string) {
	clientsMutex.RLock()
	defer clientsMutex.RUnlock()

}
