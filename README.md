# URL to QR Code Converter

This simple utility is a command-line application written in Go that converts a given URL into a QR code image. This project showcases how to handle user input, generate QR codes, and work with external libraries in Go.

## Features

- **Simple and Interactive**: Just run the program and input your URL.
- **Fast QR Code Generation**: Converts any URL into a QR code saved as `qr.png` in the current directory.
- **Minimalist Design**: A clear focus on functionality with no unnecessary dependencies.

## Prerequisites

Before you run the application, make sure you have the following installed:
- [Go](https://golang.org/dl/) (Version 1.18 or later)

## Installation

To set up the URL to QR Code Converter on your local machine, follow these steps:

1. **Clone the Repository**

   If you have git installed, you can clone the repository directly. Otherwise, you can download the source code as a zip file and extract it.

   ```bash
   git clone https://github.com/yourusername/url-to-qrcode.git
   cd url-to-qrcode
   ```

2. **Build the Application**

   Compile the application using the Go compiler.

   ```bash
   go build -o url-to-qrcode
   ```

## Usage

To use the application, run the compiled binary and follow the interactive prompt.

```bash
./url-to-qrcode
```

You will be prompted to:

```plaintext
Enter the URL you want to convert to a QR code:
```

Enter the full URL (including `http://` or `https://`) and the application will generate a QR code, saving it as `qr.png` in the current directory.

### Example

```bash
./url-to-qrcode
Enter the URL you want to convert to a QR code:
https://www.example.com
QR code generated successfully!
```

Now, you will find `qr.png` in your current directory, which represents the QR code of the entered URL.

## Project Structure

```
.
├── go.mod
├── go.sum
└── main.go
```

- **`main.go`**: The main source file where the URL input and QR code generation logic is implemented.
- **`go.mod`** and **`go.sum`**: Go module files for managing dependencies.

## Dependencies

This project uses the following Go package:

- [`github.com/skip2/go-qrcode`](https://github.com/skip2/go-qrcode): A popular library for QR code generation in Go.

## Contributing

Feel free to fork the repository, make improvements or suggest changes. This is a simple project, but contributions that improve the usability or extend the functionality are welcome.

## Questions or Issues?

If you have any questions or encounter any issues, please feel free to open an issue in the project repository.
