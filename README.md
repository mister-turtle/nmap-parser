# nmap-parser

This tool parses nmap XML files and produces file based output for easy use within other tools.

## Installation
If you have go installed already, you should be able to just run
```sh
go install github.com/mister-turtle/nmap-parser@latest
```

## Usage
```sh
Usage of nmap-parser:
  -o string
        Output directory to use (default "./nmap-parser")
  -x string
        NMap XML file to parse
```
This will provide a few useful outputs in IP:Port format:
```
output/
|-- service_all.txt             // all services discovered
|-- service_all_web.txt         // all web services
|-- service_http.txt            // each service then gets its own file
`-- service_https.txt
```