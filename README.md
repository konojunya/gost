# gost

gost = go + gist

gost is a command line tool written in Go (Golang), easily uploading code to GitHub's Gist.

## Installation

```sh
$ go get github.com/konojunya/gost
```

## Usage

1. login

```sh
$ gost login
```

After login, a token file is created in `~/.gost`.

2. create gist

```sh
$ gost create </file/to/path> [options]
```

### Options

input description

```sh
$ gost create </file/to/path> -m "description message"
```

private gist

```sh
$ gost create </file/to/path> --private
```

## Development

```
$ git clone https://github.com/konojunya/gost.git
$ cd gost
$ dep ensure
```

## Contribution

Please check the [issue](https://github.com/konojunya/gost/issues).

1. Fork it [https://github.com/konojunya/gost.git](https://github.com/konojunya/gost.git)
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create new Pull Request! :)

## Licence

MIT

## Auther

- twitter [@konojunya](https://twitter.com/konojunya)
