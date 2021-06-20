# AccelerateDB

A heavily WIP distributed kv engine written in Go. 


### Feature implementation plan

- **HashIndex** - Implement simple concurrent hash index 
- **File I/O** - Implement the required file i/o that acts as the base for reading and writing the data on disk. 
- **Concurrency** - Enable multiple DB Session for multiple readers and writes to be able to read and write at the same time without failures.


### Setup

This is a Go project and currently only needs Go 1.15 installed. 
Set this project in your GOPATH. 