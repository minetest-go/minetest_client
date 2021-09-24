# Table of Contents

- [Introduction](#introduction)
- [Usage](#usage)
  * [Client Registration](#client-registration)
  * [Session Creation](#session-creation)
- [References](#references)

# Introduction

SRP is a key exchange protocol published by Stanford in 1998. It is described well on their [website](http://srp.stanford.edu/):

>SRP is a secure password-based authentication and key-exchange protocol. It solves the problem of authenticating clients to servers securely, in cases where the user of the client software must memorize a small secret (like a password) and carries no other secret information, and where the server carries a verifier for each user, which allows it to authenticate the client but which, if compromised, would not allow the attacker to impersonate the client. In addition, SRP exchanges a cryptographically-strong secret as a byproduct of successful authentication, which enables the two parties to communicate securely.

# Usage

There are two discrete processes defined by SRP:
1. **Client Registration**
  * *A client is created by generating it's credentials, and transmitting them to the server for storage.*
2. **Session Creation**
  * *A client and server negotiate a strong session key.*

## Client Registration

```go
// A client chooses an identifier and passphrase
identifier := []byte("user123@example.com")
passphrase := []byte("Password123!")

// SRP creates a salt and verifier based on the client's identifier and passphrase
s, v, err  := srp.NewClient(identifier, passphrase)

if err != nil {
  panic("Client creation failed!")
}
```

The client now sends it's `identifier`, `s`, and `v` values to the server for storage.

## Session Creation

#### Client

```go
// The client's identifier and passphrase, registered with the server.
identifier := []byte("user123@example.com")
passphrase := []byte("Password123!")

// The client creates a public key "A" and a private key "a"
A, a, err := srp.InitiateHandshake()

if err != nil {
  panic("Handshake initiation failed!")
}
```

The client now sends it's `identifier` and `A` to the server.

#### Server

```go
// The server receives a client's "identifier" and "A" value. Assume the
// following variables are populated accordingly.
var identifier []byte
var A []byte

// The server looks up a client's salt and verifier from the provided
// identifier. Assume the following variables are populated accordingly.
var s []byte
var v []byte

// Create a public key to share with the client, and compute the session key.
B, K, err := srp.Handshake(A, v)

if err != nil {
  panic("Handshake failed!")
}
```

The server should now persist the value of `K`, and send `s` and `B` to the client.

#### Client

```go
// The client receives its salt "s" along with a public key "B" from the server.
// Assume the following variables are populated accordingly.
var s []byte
var B []byte

// Recall that the client has "A", "a", and "passphrase" variables from the
// first step of session creation.

// Compute the session key!
K, err := srp.CompleteHandshake(A, a, passphrase, s, B)

if err != nil {
  panic("Failed to complete the handshake!")
}
```

At this point, the client and server will have a shared secret `K` if the authentication was successful. The shared secret can be verified as follows:

#### Client

```go
proof := srp.Hash(K)
```

 The client can then send `proof` to the server for easy verification, and demand proof of its own.

#### Server

```go
// The server received the client's proof, and assigns it to the variable below:
var clientProof []byte

// Check if the client's proof is acceptable
if subtle.ConstantTimeCompare(clientProof, srp.Hash(K)) != 1 {
  panic("Server does not accept client's proof!")
}

serverProof := srp.Hash(s, K)
```

Now the server sends `serverProof` to the client.

#### Client

```go
// The client received the server's proof, and assigns it to the variable below:
var serverProof []byte

if subtle.ConstantTimeCompare(serverProof, srp.Hash(s, K)) != 1 {
  panic("Client does not accept server's proof!")
}
```

# References

| Name                     | Link                                | Note                            |
|--------------------------|-------------------------------------|---------------------------------|
| Stanford's SRP Home Page | http://srp.stanford.edu/            |                                 |
| RFC 2945                 | https://tools.ietf.org/html/rfc2945 | Older SRP-3 implementation      |
| RFC 5054                 | https://tools.ietf.org/html/rfc5054 | Newer SRP-6(a) implementation   |
| node-srp                 | https://github.com/mozilla/node-srp | A compatible Javascript library |
| node-srp                 | https://github.com/voynic/node-srp  | A more modern fork of ^         |
