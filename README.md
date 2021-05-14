# Primality
## Computes highest prime less than a given number


Primality is a cloud-enabled, microservice, that takes in a number and outputs the highest prime number lower than the given input number.

- Play with some numbers
- See highest prime number lower than your numbers
- ✨Magic ✨


## Tech

Primality uses a number of open source projects to work properly:

- [Gorilla Mux](https://github.com/gorilla/mux) - A powerful HTTP router and URL matcher for building Go web servers with gorilla.

And of course Primality itself is open source with a [public repository][robeng1]
 on GitHub.

## Design
- Repeated Division: A brute force test for primality that uses repeated division. Start from
the smallest candidate divisor, and then try all possible divisors up from there. Since
2 is the only even prime, once we verify that x isn’t even we only need try the odd
numbers as candidate factors. 

- Optimized Repeated Division: This is just like the first approach the only difference is that we only have to test upto √x since any two non-trivial prime factors of the input value would have already appeared before √input and the product of any two factors each greater than √x is greater than x itself which creates a contradiction.
- [Fermat Method](https://en.wikipedia.org/wiki/Proofs_of_Fermat's_little_theorem): If n is a prime number, then for every a, 1 < a < n-1,
    a<sup>(n-1)</sup> ≡ 1 (mod n)
      OR
    a<sup>(n-1)</sup> % n = 1
- [Miller Rabin](https://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test): The Miller–Rabin primality test or Rabin–Miller primality test is a probabilistic primality test: an algorithm which determines whether a given number is likely to be prime, similar to the Fermat primality test and the Solovay–Strassen primality test. Click on the link to learn more
- [Go Implementation](https://golang.org/src/math/big/prime.go): The prime package in the go library also implements the primality test by applying the Miller-Rabin test with n pseudorandomly chosen bases as well as a Baillie-PSW test. You can read more on Baillie-PSW test 
  [here](https://en.wikipedia.org/wikiBaillie%E2%80%93PSW_primality_test)

Benchmarks
```
cpu: Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz
BenchmarkOptimizedRepeatedDivisonPrimality-4        6003            188096 ns/op
BenchmarkFermatPrimality-4                         53158             22222 ns/op
BenchmarkMillerRabinPrimality-4                    59107             20008 ns/op
```
You can run the benchmarks yourself by running
```sh
cd prime
go test -bench=.
```


## Installation

Primality requires [Go](https://golang.org/dl/) v1.11+ to run.

Install the dependencies and devDependencies and start the server.

```sh
go run .
```


## Development

Want to contribute? Great!



## Docker

Primality is very easy to install and deploy in a Docker container.

By default, the Docker will expose port 10000, so change this within the
Dockerfile if necessary. When ready, simply use the Dockerfile to
build the image. You can also pull the image from 
[here](https://hub.docker.com/repository/docker/robeng1/primality/)
by running 
```sh
docker pull robeng1/primality:latest
```

or build it as follows

```sh
cd primality
docker build -t <youruser>/primality:${version} .
```

This will create the primality image and pull in the necessary dependencies.
Be sure to swap out `${version}` with the actual
version of Primality.

Once done, run the Docker image and map the port to whatever you wish on
your host. In this example, we simply map port 10000 of the host to
port 10000 of the Docker (or whatever port was exposed in the Dockerfile):

```sh
docker run -d -p 10000:10000 --restart=always --name=primality <youruser>/primality:${version}
```

You can also quickly run 
```
docker pull robeng1/primality:latest && docker run -d -p 10000:10000 --restart=always --name=primality robeng1/primality
```
and starting issuing GET requests to 127.0.0.1:10000 with the input as parameter

Example 
```
curl http://127.0.0.1:10000/55 returns {result: 53}
```

Verify the deployment by navigating to your server address in
your preferred browser.

```sh
127.0.0.1:10000
```


## UI Driver
You can also clone [this](https://github.com/robeng1/primality-ui) simple UI and run it locally to test the service.

## License

MIT

**Free Software, Hell Yeah!**
