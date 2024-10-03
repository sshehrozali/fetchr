
# Fetchr

Fetchr is a simple, fast, and reliable CLI tool designed for downloading files from the internet. With a focus on speed and ease of use, Fetchr supports multiple platforms including macOS, Linux, and Windows. Whether you need to download a single file or automate batch downloads, Fetchr has you covered.

## Features

- **Multi-platform Support**: Runs on macOS, Linux, and Windows.
- **Fast Downloads**: Optimized for speed to ensure quick file retrieval.
- **Command Line Interface**: Easy-to-use CLI for seamless interaction.
- **Batch Downloading**: Automate downloads of multiple files with ease.

## Installation

### Option 1: Download Binaries (Recommended)

The easiest way to use Fetchr is to download the pre-built binaries for your operating system. You can find them under the [Releases tab](https://github.com/sshehrozali/fetchr/releases).

1. Go to the Releases tab.
2. Download the binary for your operating system (macOS, Linux, or Windows).
3. Add the binary to your system's PATH.

### Option 2: Build from Source

If you prefer to build Fetchr from source, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/sshehrozali/fetchr.git
   ```

2. Navigate into the directory:

   ```bash
   cd fetchr
   ```

3. Build the application:

   ```bash
   go build -o fetchr
   ```

4. Move the executable to your PATH:

   ```bash
   mv fetchr /usr/local/bin/  # For Linux or macOS
   ```

## Usage

You can use Fetchr to download files by running the following command in your terminal:

```bash
fetchr <URL>
```

To download multiple files at once, provide a list of URLs:

```bash
fetchr <URL1> <URL2> <URL3>
```

## Upcoming Features (WIP)

- **Automated Bulk Download**: Streamline downloading large batches of files from a predefined list or directory.
- **Improved Error Handling**: More detailed error messages for failed downloads.
- **Proxy Support**: Add the ability to download through a proxy server.

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

- **Nehad Shakoor** - [Nehadsys](https://github.com/Nehadsys)

## Acknowledgments

- Thanks to the Go community for their excellent resources and support.
