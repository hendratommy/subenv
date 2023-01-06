# subenv

A command line tools to help you substitute variable placeholder in a file with environment variables or files (`env` files)

```bash
subenv hello.txt
```

## Usage

```bash
$ subenv --help
Usage of subenv:
  -c string
        Encoder name to use, available encoder: [ base64 ]
  -d string
        Decoder name to use, available decoder: [ base64 ]
  -e value
        File to use as env source
  -noos
        This flag can only be use when using -e. If set to true, will
        not use OS environment variables
  -v    Print version
  -version
        Print version
```

## Build from source

- Pull this repository

  ```bash
  git pull https://github.com/hendratommy/subenv.git
  ```

- To build without version:

  ```bash
  # MacOS
  make build-darwin
  
  # Linux
  make build-linux

  # windows 
  make build-windows
  ```

- To build with version

  ```bash
  # MacOS
  make VERSION=CURRENT_VERSION build-darwin
  
  # Linux
  make VERSION=CURRENT_VERSION build-linux

  # windows 
  make VERSION=CURRENT_VERSION build-windows
  ```

- To build distribution:

  ```bash
  make VERSION=CURRENT_VERSION all
  ```

### Using .env files

```bash
subenv -e test.env test.txt
```

We can supply more than 1 env files by repeating `-e` arguments:

```bash
subenv -e test1.env -e test2.env test.txt
```

### Using `encoder` & `decoder`

We can use `encoder` (`-c`) to encode substituted values, or `decoder` (`-d`) to decode the variable value (if the environment variable is
already decoded)

```bash
$ NAME=World subenv -c base64 hello2.txt
Hello V29ybGQ=
```

```bash
$ NAME=V29ybGQ= subenv -d base64 hello2.txt
Hello World
```

Supported `encoder`:

| Name | Description | Encoder | Decoder |
| --- | --- | --- | --- |
| base64 | Encode or decode using base64 | Y | Y |
