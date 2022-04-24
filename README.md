# pwdgen 
### Generate secure passwords in your terminal

## CLI Usage

```sh
pwdgen [-s size]
```

## Install

- CLI:
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


