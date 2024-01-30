```markdown
# IP Resolver

Das ist eine einfache Go-Anwendung, die die IPv4- und IPv6-Adressen des Clients über eine HTTP-Schnittstelle zurückgibt.

## Installation und Verwendung mit Docker Compose

1. Stelle sicher, dass Docker und Docker Compose auf deinem System installiert sind.

2. Klone dieses Repository:

   ```bash
   git clone https://github.com/yourusername/ip-resolver.git
   ```

3. Navigiere in das Projektverzeichnis:

   ```bash
   cd ip-resolver
   ```

4. Passe bei Bedarf die Nginx-Konfiguration in der `nginx/nginx.conf`-Datei an.

5. Starte die Anwendung mit Docker Compose:

   ```bash
   docker-compose up -d
   ```

6. Die Anwendung ist jetzt unter [http://localhost:8080/getip](http://localhost:8080/getip) erreichbar.

7. Um die Anwendung zu stoppen, führe aus:

   ```bash
   docker-compose down
   ```

8. Curl Script zum extrahieren der Werte über die Bash

   ```bash
   #!/bin/bash

   # HTTP-Anfrage an die Go-Anwendung senden und Antwort in Variable speichern
   response=$(curl -s http://localhost:8080/getip)

   # IPv4-Adresse aus JSON extrahieren
   ipv4_address=$(echo "$response" | jq -r .ipv4_address)

   # IPv6-Adresse aus JSON extrahieren
   ipv6_address=$(echo "$response" | jq -r .ipv6_address)

   # Ausgabe der Adressen
   echo "IPv4-Adresse: $ipv4_address"
   echo "IPv6-Adresse: $ipv6_address"
   ```

## Anpassungen

- Du kannst die Nginx-Konfiguration in der `nginx/nginx.conf`-Datei anpassen, um den Reverse Proxy an deine Anforderungen anzupassen.

- Die Go-Anwendung kann bei Bedarf im `main.go`-Code weiter angepasst werden.

Hinweis: Stelle sicher, dass die Portnummern in der Docker Compose-Konfiguration und in deinen Anfragen übereinstimmen.
