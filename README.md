# High throughput Fizz Buzz

This is my submission for the [High throughput Fizz Buzz challenge](https://codegolf.stackexchange.com/questions/215216/high-throughput-fizz-buzz).

On a MacBook Pro M1 Max 2021, the throughput is around 4.4 GiB/s.

## Implementation

The program spawns several workers responsible for filling buffers with the expected Fizz Buzz data. Those workers are CPU-bound and run in parallel (one worker per CPU). Meanwhile, another goroutine is responsible for fetching complete buffers in the right order and writing those to stdout. There are a few other optimizations: the main loop is unrolled to perform the same 15 computations on each step (no need to compare modulos) and integers are written in base 1000 without allocation.

## History

- [naive](https://github.com/nockty/high-throughput-fizz-buzz/blob/10ae663a72c98860af0b743891d26b8092d5fba4/main.go): 11 MiB/s
- [buffering](https://github.com/nockty/high-throughput-fizz-buzz/blob/bfa3e42fb0b94a3dbd823177f91cd80e8118ffe9/main.go): 170 MiB/s
- [faster string conversion](https://github.com/nockty/high-throughput-fizz-buzz/blob/f579de801e5ee2129f1b8837dbca2930f538e3bc/main.go): 380 MiB/s
- [unrolled loop](https://github.com/nockty/high-throughput-fizz-buzz/blob/fadba2e28ce8f27a455a818fa64c5830aff448dc/main.go): 450 MiB/s
- [write integers without allocation](https://github.com/nockty/high-throughput-fizz-buzz/blob/4032f56bba27aab36b08abb44da2ec56472af2c0/main.go): 840 MiB/s
- [write integers in base 1000](https://github.com/nockty/high-throughput-fizz-buzz/blob/3afa43747ef1d576370fa23833fcb21ebbab9136/main.go): 940 MiB/s
- [write to stdout in a dedicated goroutine](https://github.com/nockty/high-throughput-fizz-buzz/blob/68b2c6f2e1951e55b602e94d56c668a1161dd418/main.go): 1.23 GiB/s
- [compute data in parallel workers](https://github.com/nockty/high-throughput-fizz-buzz/blob/main/main.go): 4.4 GiB/s
