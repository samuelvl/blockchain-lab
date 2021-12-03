# Blockchain Lab

Learn about blockchain technology.

## Usage

Build the program with the `Go` compiler (see Makefile):

```shell
$ make build
```

Then, generate your first block:

```shell
$ make run
Block:
{
  "hash": "gd3I0kiy3M3T/dXoTwytYrCPLRC1f5qDHBNFHlxcgKU=",
  "data": "Genesis",
  "prevHash": ""
}
```

## Testing

The whole project has been written using the `TDD` methodology with the help of
the `testify` framework.

Use the following command if you want to run the unit tests:

```shell
$ make test
```

## References

- https://github.com/nheingit/go-blockchain
- https://github.com/stretchr/testify
