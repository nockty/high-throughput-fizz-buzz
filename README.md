# Fizz Buzz

## History

- naive: 11 MiB/s
- buffering: 170 MiB/s
- faster string conversion: 380 MiB/s
- unrolled loop: 450 MiB/s
- write integers without allocation: 840 MiB/s
- write integers in base 1000: 940 MiB/s
- write to stdout in a dedicated goroutine: 1.23 GiB/s
- compute data in parallel workers: 4.4 GiB/s
