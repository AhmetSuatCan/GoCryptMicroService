package main

import (
    "log"
    "runtime"
    "os"
    "fmt"
    "net/http"
    "goCaaS/internal/encryption"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/encrypt", encryption.HandleEncryptRequest)

    addr := ":8080"
    log.Printf("Server listening on %s", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
}


func logResourceUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	pid := os.Getpid()
	fmt.Printf("---- Resource Usage ----\n")
	fmt.Printf("PID: %d\n", pid)
	fmt.Printf("Alloc = %v KB\n", m.Alloc/1024)
	fmt.Printf("TotalAlloc = %v KB\n", m.TotalAlloc/1024)
	fmt.Printf("Sys = %v KB\n", m.Sys/1024)
	fmt.Printf("NumGC = %v\n", m.NumGC)
	fmt.Println("------------------------")
}
