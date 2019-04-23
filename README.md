<div align="center">
	<img width="500" src=".github/logo.svg" alt="pinpt-logo">
</div>

<p align="center" color="#6a737d">
	<strong>Golang SQL driver for Dremio database via a proxy server</strong>
</p>

## Overview

You probably don't want this. Check out [go-dremio](https://github.com/pinpt/go-dremio) instead.

## Install

```
go get -u github.com/pinpt/go-dremioproxy
```

## Usage

Use it like any normal SQL driver.

```
db, err := sql.Open("dremioproxy", "http://localhost:8047")
```

## License

All of this code is Copyright &copy; 2019 by Pinpoint Software, Inc. Licensed under the MIT License
