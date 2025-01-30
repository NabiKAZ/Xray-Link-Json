/**
 * This program takes input from the command line and checks whether the input is a JSON or a link.
 * If the input is a JSON, it converts it to Share links, and if the input is a link, it converts it to JSON.
 * This is done using functions from the libXray library, and the result is encoded in Base64.
 * Programmer: NabiKAZ (x.com/NabiKAZ)
 * https://github.com/NabiKAZ/Xray-Link-Json
 */

package main

import (
        "encoding/base64"
        "fmt"
        "log"
        "os"
        "strings"

        "github.com/xtls/libxray"
)

func main() {
        // Check if command-line argument is provided
        if len(os.Args) < 2 {
                log.Fatal("Please provide a Share link or Xray JSON")
        }

        input := os.Args[1] // Get the input from the second command-line argument
        fmt.Println("Input:", input)

        // Determine if the input is a Share link or Xray JSON based on the first character
        if strings.HasPrefix(input, "{") {
                // Convert Xray JSON to Share links
                convertXrayJsonToShareLinks(input)
        } else {
                // Convert Share link to Xray JSON
                convertShareLinkToXrayJson(input)
        }
}

// Convert Share link to Xray JSON
func convertShareLinkToXrayJson(shareLink string) {
        // Print the input link
        fmt.Println("Processing Share link...")

        // Convert the Share link to Base64 (if it's not already Base64)
        shareLinkBase64 := base64.StdEncoding.EncodeToString([]byte(shareLink))

        // Call the ConvertShareLinksToXrayJson function from the libXray package
        result := libXray.ConvertShareLinksToXrayJson(shareLinkBase64)

        // Check and display the result
        if result == "" {
                log.Fatal("Error converting Share link to Xray JSON")
        } else {
                // Decode from Base64
                decodedBytes, err := base64.StdEncoding.DecodeString(result)
                if err != nil {
                        log.Fatalf("Error decoding: %v", err)
                }

                // Print the decoded data
                fmt.Println("Decoded Xray JSON:")
                fmt.Println(string(decodedBytes))
        }
}

// Convert Xray JSON to Share links
func convertXrayJsonToShareLinks(xrayJson string) {
        // Print the input Xray JSON
        fmt.Println("Processing Xray JSON...")

        // Base64 encode the Xray JSON string
        encodedJson := base64.StdEncoding.EncodeToString([]byte(xrayJson))

        // Call the ConvertXrayJsonToShareLinks function from the libXray package
        result := libXray.ConvertXrayJsonToShareLinks(encodedJson)

        // Check and display the result
        if result == "" {
                log.Fatal("Error converting Xray JSON to Share links")
        } else {
                // Decode from Base64
                decodedBytes, err := base64.StdEncoding.DecodeString(result)
                if err != nil {
                        log.Fatalf("Error decoding: %v", err)
                }

                // Print the decoded data
                fmt.Println("Decoded Share links:")
                fmt.Println(string(decodedBytes))
        }
}
