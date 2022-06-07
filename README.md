# pwdgen 
### Generate secure passwords in your terminal

![gif](./example.gif)

## CLI Usage

```sh
pwdgen [-s size]
```

## Install

- Brew:

  ```
  brew tap PauSabatesC/pwdgen https://github.com/PauSabatesC/pwdgen
  brew install pwdgen
  ```

- Docker:

  ```sh
  docker run pausc/pwdgen
  ```

  ```sh
  # docker run pausc/pwdgen -s {size}
  # docker run pausc/pwdgen version
  ```

- Downloading the binary:

  You can get the available binaries from here: https://github.com/PauSabatesC/pwdgen/releases
  ```sh
  wget pwdgen_{version}_{architecture}.tar.gz
  tar -xvf pwdgen_{version}_{architecture}.tar.gz
  cp pwdgen /usr/local/bin
  ```

## Install Module

```sh
go get github.com/PauSabatesC/pwdgen/generator
```


