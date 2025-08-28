# FileInfo Utility

A simple and efficient command-line utility written in Go that displays file information including size, modification time, and permissions.

## Features

- 📁 Display file size in bytes with human-readable format (KB, MB, GB)
- 🕒 Show modification time with customizable format
- 🔐 View file permissions
- 📊 Process multiple files at once
- ✅ Proper error handling and exit codes
- 🚫 Skip directories and report errors appropriately

## Installation

### From Source

1. Clone or download this repository
2. Build the application:

```bash
go build -o fileinfo .
```

### Install Globally
```bash
go install
```

This will install the binary to your `$GOPATH/bin` directory.

### Download Pre-built Binaries
Check the [Releases page](https://github.com/lyubchenko-ivan/fileinfo/releases) for pre-built binaries for various platforms.

## Usage
### Basic Usage
```bash
# Single file
fileinfo document.txt

# Multiple files
fileinfo file1.txt file2.jpg file3.pdf
```

### Output Example
```text
File: document.txt
  Size: 2048 bytes (2.00 KB)
  Modified: 2024-01-15 14:30:25
  Permissions: -rw-r--r--
---
File: image.jpg
  Size: 1572864 bytes (1.50 MB)
  Modified: 2024-01-14 09:15:40
  Permissions: -rw-r--r--
---
```

### Command Line Options
```bash
# Show help
fileinfo -h
fileinfo --help

# Display version
fileinfo -v
fileinfo --version
```

## Building from Source
### Prerequisites
* Go 1.16 or higher

### Build Steps
```bash
# Clone the repository
git clone https://github.com/yourusername/fileinfo.git
cd fileinfo

# Build the binary
go build -o fileinfo .

# Alternatively, install directly
go install
```

### Cross-compilation
```bash
# Build for Linux
GOOS=linux GOARCH=amd64 go build -o fileinfo-linux .

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o fileinfo.exe .

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o fileinfo-macos .
```

## Project Structure
```text
fileinfo/
├── main.go          # Main application code
├── go.mod           # Go module definition
└── README.md        # This file
└── LICENSE
```

## Development
### Code Formatting
```bash
go fmt
```

### Adding Dependencies
```bash
go get <package-name>
```

## Exit Codes
* 0: Success - all files processed successfully

* 1: Error - one or more files could not be processed

## Error Handling
The utility provides clear error messages for:

* Non-existent files

* Permission denied errors

* Directory paths (directories are skipped)

* Other filesystem errors

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing
1. Fork the repository

2. Create a feature branch (``git checkout -b feature/amazing-feature``)

3. Commit your changes (``git commit -m 'Add amazing feature'``)

4. Push to the branch (``git push origin feature/amazing-feature``)

5. Open a Pull Request

## Support
If you encounter any issues or have questions:

1. Check the [Issues page](https://github.com/lyubchenko-ivan/fileinfo/issues)

2. Create a new issue with detailed description

## Version History
* v1.0.0 (2024-01-15)

  * Initial release

  * Basic file information display

  * Multiple file support

  * Human-readable size formatting