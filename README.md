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
Chain has 4 blocks.
Block 3 is:
{
  "hash": "AADIiJwhbjrT0kJHMKZaVyffDJB44+HxiMEb6XjTl/o=",
  "data": "third block after genesis",
  "prevHash": "AACWgBRDSbaOrrPLg4Fk2XYh1YAWgdNntC8VZJ8Z4pQ=",
  "nonce": 21680
}
...
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
- https://jeiwan.net/posts/building-blockchain-in-go-part-1
- https://github.com/stretchr/testify
