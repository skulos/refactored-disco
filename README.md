# MN8 DNS Setup for Internal Network


## The ISC BIND Nameserver

(**NOTE** `BIND9` is used - 25/11/2020)

**`BIND`** *(Berkeley Internet Name Domain)* is an implementation of the `DNS protocols` and provides an openly redistributable reference implementation of the major components of the Domain Name System, including:
* Domain Name System server
* Domain Name System resolver library
* Tools for managing and verifying the proper operation of the DNS server

The **`BIND`** DNS Server, `named`, is used on the vast majority of name serving machines on the Internet, providing a robust and stable architecture on top of which an organization's naming architecture can be built.

The resolver library included in the `BIND` distribution provides the standard APIs for translation between domain names and Internet addresses and is intended to be linked with applications requiring name service.

`BIND version 9` is a major rewrite of nearly all aspects of the underlying BIND architecture. Some of the important features of `BIND9` are DNS Security (DNSSEC, TSIG), IPv6, DNS Protocol Enhancements (IXFR, DDNS, DNS Notify, EDNS0), Views, Multiprocessor Support, and an Improved Portability Architecture.

**NOTE : BIND versions 4 and 8 are officially deprecated.**

## DNS Record types:
### 1.  A
`Address Mapping record:` Also known as a DNS host record, resolves a hostname to its corresponding IPv4 address.

### 2. AAAA
`IP Version 6 Address record:` Resolves a hostname tp its corresponding IPv6 address.

###

###

###

###