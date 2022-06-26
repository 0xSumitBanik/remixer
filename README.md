## Remixer
![](https://img.shields.io/github/downloads/isumitbanik/remixer/total?color=blue&style=flat-square)

`remixer` allows you to generate shareable Remix IDE link for Solidity Smart contracts in your system.

The tool has only one command for utility:

- `generate`: Generate commands generates the shareable URL.

#### Installation

Download the appropriate binary from the [releases](https://github.com/iSumitBanik/remixer/releases) page and add it to the PATH.

#### Usage

```plain
remixer generate --help

Share Solidity Code Snippets as a shareable Remix IDE URL

Usage:
  remixer generate [flags]

Aliases:
  generate, gen

Flags:
      --auto-compile string   Toggle flag for turning on/off the auto compiler (default "true")
      --evm-version string    Version of the Ethereum Virtual Machine (default "default")
  -f, --file string           Smart Contract File for generating the shareable URL
  -h, --help                  help for generate
  -l, --language string       Supported Values: Solidity or Yul (default "Solidity")
      --optimize string       Toggle flag for turning on/off Solidity Optimizer (default "false")
  -p, --plugins string        List of Plugins to be activated for the Remix IDE (seperated by commas)
  -t, --theme string          Theme for the Remix IDE (default "Dark")
```

#### Contributing

This project has been developed as a part of implementing my knowledge on Go. If you feel there are certain areas that can be optimized, please feel free to contribute to this project.