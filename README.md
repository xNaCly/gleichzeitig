# gleichzeitig

Run shell scripts, commands and executables in parallel (ger: gleichzeitig).

https://user-images.githubusercontent.com/47723417/228802151-056a549a-5bbf-43af-8de3-b1f84b2a5792.mp4

## Install

### From release

1. Download the [latest release](https://github.com/xNaCly/gleichzeitig/releases/latest) for your architecture.
2. Extract the archive
3. Move the binary to an executable path:
   - Linux:
     ```console
     $ sudo mv ./gleichzeitig /usr/bin/gleichzeitig
     ```
   - Windows:
     > copy `gleichzeitig` to `C:\Program Files`

### From source

> requires git and go

1. clone the repo

```console
$ git clone https://github.com/xnacly/gleichzeitig
$ cd gleichzeitig
```

2. build the project

```console
$ go build
```

3. move the build executable into the path

```console
$ sudo mv ./gleichzeitig /usr/bin/gleichzeitig
```

4. execute to check if everything worked

```console
$ gleichzeitig
```

## Concept

_gleichzeitig_ starts all commands specified in its configuration and keeps them running until they terminate.
It redirects all stdout and stderr to its own output and throws errors if a process terminates prematurely or with a non 0 exit code.

## Configuration

To keep the _gleichzeitig_ setup low friction, the `gleichzeitig init` command creates the configuration directory and the `config.json` file with the default configuration in it.

The configuration is located at `.gleichzeitig/config.json` in the current working directory and is composed of the following key value pairs:

```json
{
  "color_output": true,
  "commands": [],
  "log_file": "",
  "surpress_output": false
}
```

- `color_output`: whether the output of the individual commands should be coloured different to allow the user to differentiate between them
- `commands`: array of instructions _gleichzeitig_ executes, of the following format:
  ```json
  {
    "cwd": ".gleichzeitig",
    "cmd": "echo 'test'"
  }
  ```
  - `cwd`: specify the working directory for the command to execute in, supports relative paths
  - `cmd`: the command to execute, if not found error is thrown
- `log_file`: the file to write all logs to
- `surpress_output`: whether the commands stdout and stderr should be printed

## Usage

| command    | description                                                      | example                                          |
| ---------- | ---------------------------------------------------------------- | ------------------------------------------------ |
|            | executes the commands in the config                              | `gleichzeitig`                                   |
| run,r      | executes the arguments as commands                               | `gleichzeitig run "echo 'hello world'" "ls -la"` |
| init, i    | creates the config dir, file and writes the default config to it | `gleichzeitig init`                              |
| version, v | prints the version, exits                                        | `gleichzeitig version`                           |
| help, h    | prints the help, exits                                           | `gleichzeitig help`                              |
