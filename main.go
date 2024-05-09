package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/skip2/go-qrcode"
)

func generateQRCode(url string) error {
	err := qrcode.WriteFile(url, qrcode.Medium, 256, "qr.png")
	if err != nil {
		return err
	}
	return nil
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the url to generate QR code: ")

	// Read URL from the command line input

	scanner.Scan()
	url := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading url:", err)
	}

	// Generate QR code
	err := generateQRCode(url)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		os.Exit(1)
	}

	fmt.Println("QR code generated successfully")

}
