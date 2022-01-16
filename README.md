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
  "data": "dGhpcmQgYmxvY2sgYWZ0ZXIgZ2VuZXNpcw==",
  "hash": "0000b9de761e9c4bb7a62b878811f897f96f9254579ebf2e818130a5a9633fd2",
  "prevHash": "0000ca3dea3f51de88bcc8c4d42169450ce219655562a2b8a2a8444e83351eaa",
  "nonce": 87333
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
