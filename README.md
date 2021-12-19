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
Chain size is 4 blocks
Block 0 is:
{
  "hash": "AAAbYKPkOFcxWkh0z4iGQ20gkmRzC+9HuDRPynEPwhM=",
  "data": "Genesis",
  "prevHash": "",
  "nonce": 668
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
