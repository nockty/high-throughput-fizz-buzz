# High throughput Fizz Buzz

This is my submission for the [High throughput Fizz Buzz challenge](https://codegolf.stackexchange.com/questions/215216/high-throughput-fizz-buzz).

On a MacBook Pro M1 Max 2021, the throughput is around 4.4 GiB/s.

## Implementation

The program spawns several workers responsible for filling buffers with the expected Fizz Buzz data. Those workers are CPU-bound and run in parallel (one worker per CPU). Meanwhile, another goroutine is responsible for fetching complete buffers in the right order and writing those to stdout. There are a few other optimizations: the main loop is unrolled to perform the same 15 computations on each step (no need to compare modulos) and integers are written in base 1000 without allocation.

## History

- naive: 11 MiB/s
- buffering: 170 MiB/s
- faster string conversion: 380 MiB/s
- unrolled loop: 450 MiB/s
- write integers without allocation: 840 MiB/s
- write integers in base 1000: 940 MiB/s
- write to stdout in a dedicated goroutine: 1.23 GiB/s
- compute data in parallel workers: 4.4 GiB/s
