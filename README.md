# Wazuh Rules and Decoders Generator

Wazuh Rules and Decoders Generator, an application designed to facilitate interaction with Wazuh rules and decoders. The application is developed using Go and utilizes the fyne.io library for its GUI capabilities.

## Table of Contents

- [Wazuh Rules and Decoders Generator](#wazuh-rules-and-decoders-generator)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
    - [Compilation](#compilation)
  - [Usage](#usage)
  - [Contributing](#contributing)
  - [License](#license)

## Features

- Create decoders and rules from a syslog.
- Upload .xml files of rules and/or decoders to Wazuh, with several review methods to avoid errors.
- Search for existing decoders by name.
- Search for existing rules by their ID.

## Prerequisites

Before you can run or compile the application, you must have the following installed:

- Go programming language
- gcc compiler
- Fyne toolkit (Install using the command below)

  ```bash
  go install fyne.io/fyne/v2/cmd/fyne@latest
  ```

  Ensure that Go is properly added to your PATH:

  ```bash
  export PATH="$PATH:$(go env GOBIN):$(go env GOPATH)/bin"
  ```

## Installation

Clone the repository from GitHub:

```bash
git clone https://github.com/SebasJacome/wazuh-rulesndecoders-generator.git
cd wazuh-rulesndecoders-generator
```

## Configuration

Before compiling the application, ensure you have a configuration file named conf.toml in the /api directory with the following structure:

```toml
host = "WAZUH_API_URL"
port = "WAZUH_API_PORT"
username = "wazuh-ui"
password = "WAZUH_API_PASSWORD"
```

### Compilation

To compile the application, use the provided Makefile:

```bash
make
```

This will generate:

A `.exe` file on Windows platforms.
A `.app` file on macOS platforms.
Note: By default, it tries to open a `.app` file even on Windows, but the `.exe` file will be in the root of the repository.

## Usage

After compilation, you can run the application directly from the executable file generated in the root directory.

## Contributing

Contributions to this project are welcome. Please ensure to update tests as appropriate and follow the existing coding style.

## License

This project is licensed under the [GPL-3.0](LICENSE).
