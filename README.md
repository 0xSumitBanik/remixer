### Remixer

#### About
`remixer` allows you to generate shareable Remix IDE link for Solidity Smart contracts in your system.

The tool has only one command for utility:

- `generate`: Generate commands generates the shareable URL.

#### Installation



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