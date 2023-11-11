> Script to get tlds from a list of subdomains

### Installation

```
git clone https://github.com/0xdln1/tld ; cd tld; go build .;sudo mv tld /usr/local/bin
```


### Usage

`cat subs.txt`

```text
meet.google.com
admin.google.co.uk
support.google.com.hk
```

`Command`

```
cat subs.txt | tld
```

`Output`

```
google.com
google.co.uk
google.com.hk
```
