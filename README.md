# test3

Fibonacci number generator.

API: GET /fibonacci?n=3

Generator does HTTP response streaming, so it could be used to generate really huge
sequences. On my MBP it generates approximately 15MB/s of data (in serialized representation).

## Running

Go is required:

    make
    test3

Alternatively, it could be run from Docker container:

    docker run -d -p 3000:3000 smira/test3

## Tests

With service running (on default localhost:3000) run:

    make system-test

This performs simple integration testing.

## Notes

Error checking is really simple, it should have been returning HTTP 400 errors in most cases.
