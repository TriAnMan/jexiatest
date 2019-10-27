# Klingon transliteration helper
This commandline tool transliterates names from Latin to Klingon alphabets

## Installation
Install [Golang 1.12](https://golang.org/doc/install) or higher.

Run `go get github.com/TriAnMan/jexiatest/...` and `go build github.com/TriAnMan/jexiatest/cmd/klingon-translit`

## Run
`./klingon-translit Leonard McCoy`

## Design decisions
1. App exploits a fail-fast methodology.
2. Simple DDD pattern is used to prevent circular dependencies (https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/).
3. Non over-engineered architecture is implemented to reduce development time.
4. Main transliteration routine can be imported in other apps and libraries.
5. App logs to STDERR to facilitate log management (https://12factor.net/logs) and to separate logs from normal output.
6. Transliteration algorithm has O(n) space and time complexity for string interface;
    O(1) space complexity can be achieved with buffered input and output without a BC break;
    O(1) space for algorithm is non essential for this task because command line arguments have limited length (~128KiB for linux).
7. Algorithm uses better transliteration for strings containing "ngh". Output will be "n, gh" instead of "ng, H".

## Possible future improvements
1. Formalize specs to interpret stapi.co data.
2. Implement a microservice or a STDIN parsing application if demand for transliteration increases.
3. Implement a Swagger client for stapi.co (https://godoc.org/github.com/go-swagger/go-swagger)
    if more data from stapi.co be required.
4. Move transliteration and stapi.co data mining into different apps or services 
    because this tasks are very different in nature.
    