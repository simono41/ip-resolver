package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

type IPAddressResponse struct {
	IPv4Address string `json:"ipv4_address"`
	IPv6Address string `json:"ipv6_address"`
}

func getIPHandler(w http.ResponseWriter, r *http.Request) {
	ipHeader := r.Header.Get("X-Real-IP")
	if ipHeader == "" {
		http.Error(w, "X-Real-IP Header nicht vorhanden", http.StatusBadRequest)
		return
	}

	ip := strings.TrimSpace(strings.SplitN(ipHeader, ",", 2)[0]) // Um den Fall mit mehreren IP-Adressen zu behandeln

	ipResponse := IPAddressResponse{
		IPv4Address: getIPv4(ip),
		IPv6Address: getIPv6(ip),
	}

	jsonResponse, err := json.Marshal(ipResponse)
	if err != nil {
		http.Error(w, fmt.Sprintf("Fehler beim Erstellen der JSON-Antwort: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func getIPv4(ip string) string {
	ipAddr := net.ParseIP(ip)
	ipv4 := ipAddr.To4()
	if ipv4 != nil {
		return ipv4.String()
	}
	return ""
}

func getIPv6(ip string) string {
	ipAddr := net.ParseIP(ip)
	ipv6 := ipAddr.To16()
	if ipv6 != nil && strings.Contains(ip, ":") {
		return ipv6.String()
	}
	return ""
}

func main() {
	http.HandleFunc("/getip", getIPHandler)
	http.ListenAndServe(":8080", nil)
}
