package encryption

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"goCaaS/internal"
	"goCaaS/internal/pools/ciphers"
	"goCaaS/internal/pools/salts"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

type EncryptionResponse struct {
	EncryptedData string `json:"encryptedData"`
}

func HandleEncryptRequest(w http.ResponseWriter, r *http.Request) {
	var data internal.EncryptionData

	if r.Method != http.MethodPost {
		http.Error(w, "Method not Allowed!", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	ip := getIPAddress(r)

	encryptor := EncryptConfig{UUID: data.UserUUID, saltSize: len(salts.AllSalts), cipherSize: len(ciphers.AllCiphers)}
	cipherIndex, saltIndex := encryptor.GenerateIndexPair(encryptor.cipherSize, encryptor.saltSize)

	encryptedData := ciphers.AllCiphers[cipherIndex](data.SensitiveData, salts.AllSalts[saltIndex])

	log := InfoTable{
		usedAlgorithm:         ciphers.AllCiphersMetadata[cipherIndex],
		usedSaltValue:         salts.AllSalts[saltIndex],
		informedUUID:          data.UserUUID,
		informedSensitiveData: data.SensitiveData,
		informedAPIKey:        data.ApiKey,
		clientIpAddr:          ip,
	}
	printEncryptionDetails(log)

	response := EncryptionResponse{EncryptedData: encryptedData}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func getIPAddress(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		ips := strings.Split(ip, ",")
		return strings.TrimSpace(ips[0])
	}

	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // fallback (may include port)
	}

	return ip
}

type InfoTable struct {
	usedAlgorithm         string
	usedSaltValue         string
	informedUUID          string
	informedSensitiveData string
	informedAPIKey        string
	selectedStrategy      string
	clientIpAddr          string
}

func printEncryptionDetails(log InfoTable) {
	data := [][]string{
		{"Used Algorithm:", log.usedAlgorithm},
		{"Used Salt Value:", log.usedSaltValue},
		{"Informed UUID", log.informedUUID},
		{"Informed Sensitive Data", log.informedSensitiveData},
		{"Informed API Key", log.informedAPIKey},
		{"Selected Strategy", "Standard Security"},
		{"Client IP Address", log.clientIpAddr},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Encryption Details"})
	table.Bulk(data)
	table.Render()

}
