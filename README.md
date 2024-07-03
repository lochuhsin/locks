<a name="readme-top"></a>


<!-- PROJECT LOGO -->
<br />
<div align="center">
  <h3 align="center">Concurrency and Lock</h3>
  <p align="center">
    This is a repository for learning concurrency patterns and locks.
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project
This is a project that aims to learn and implement all the lock related algorithms and functions.

### Prerequisites
There are no prerequisites for this package.

### Installation
1. Go package manager
   ```sh
   go get -u https://github.com/lochuhsin/re-snowflake.git
   ```
2. Clone
   ```sh
    git clone https://github.com/lochuhsin/re-snowflake.git
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Performance
In general, it takes around 40ns for generating one id.
```sh
go test ./... -run=none -bench=. -benchmem

goos: darwin
goarch: arm64
pkg: cas
BenchmarkMutex_10-8               456631              2649 ns/op             512 B/op         23 allocs/op
BenchmarkMutex_100-8               52848             21476 ns/op            4832 B/op        203 allocs/op
BenchmarkMutex_200-8               21631             53116 ns/op            9637 B/op        403 allocs/op
BenchmarkMutex_1000-8               4017            333472 ns/op           48116 B/op       2003 allocs/op
BenchmarkCAS_10-8                 467710              2515 ns/op             424 B/op         22 allocs/op
BenchmarkCAS_100-8                 48086             26302 ns/op            4024 B/op        202 allocs/op
BenchmarkCAS_200-8                 23480             52219 ns/op            8024 B/op        402 allocs/op
BenchmarkCAS_1000-8                 3397            318700 ns/op           40024 B/op       2002 allocs/op
PASS
ok      cas     13.151s
```

<!-- USAGE EXAMPLES -->
## Usage
<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing
Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

ChuHsin(Albert) Lo - [@Linkedin](https://www.linkedin.com/in/lochuhsin/) - lochuhsin@gmail.com

Project Link: [https://github.com/lochuhsin/re-snowflake](https://github.com/lochuhsin/re-snowflake)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

