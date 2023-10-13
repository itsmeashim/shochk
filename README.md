# shochk

Validate Shodan API keys quickly and efficiently. `shochk` checks the validity of Shodan API keys and outputs their status. 

## Usage Example

Imagine you have a file called `keys.txt` containing several Shodan API keys. You can use `shochk` to validate them in bulk:

```
▶ cat keys.txt
YOUR_FIRST_API_KEY
YOUR_SECOND_API_KEY
...

▶ shochk -file keys.txt
YOUR_FIRST_API_KEY - Valid (Member: true, Credits: 10)
YOUR_SECOND_API_KEY - Invalid
...
```

You can also pipe input from stdin:
```
▶ echo "YOUR_API_KEY" | shochk
YOUR_API_KEY - Valid (Member: true, Credits: 10)
```

Or check a single token:
```
▶ shochk -token YOUR_API_KEY
YOUR_API_KEY - Valid (Member: true, Credits: 10)
```


## Flags

- `-token`: Specify a single Shodan API key to check.
- `-file`: Specify a file containing a list of Shodan API keys to check (one per line).

## Install

You can either install using go:
```bash
go install -v github.com/itsmeashim/shochk@latest
```


Or clone the repository and build manually:
```
git clone https://github.com/itsmeashim/shochk.git
cd shochk
go build .
```
