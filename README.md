# dhash

dhash recursively hashes all of the files in a directory and produces
a cumulative hash string. It builds on the directory hashing functionality
used by the Go compiler.

## Features

- Hashes all files in a given directory (and child directories)
- Supports go.mod, SHA-256, and SHA-512 hash types
- Optionally write each individual file's hash to the output

## Examples

```sh
# Produce a single SHA-256 hash of the current directory:
dhash .
e21bb4960769b574e65dc440e5e70cabab363cd4d6b19290f389c77d8333efdc
# Include all of the file hashes:
dhash -e .
.git/HEAD - 28d25bf82af4c0e2b72f50959b2beb859e3e60b9630a5e8c603dad4ddb2b6e80
.git/config - cae33efdb02cf774435c1ff9cb16bcc1014606908530c6e1dc727615fe3e8cda
.git/description - 85ab6c163d43a17ea9cf7788308bca1466f1b0a8d1cc92e26e9bf63da4062aee
.git/hooks/applypatch-msg.sample - 0223497a0b8b033aa58a3a521b8629869386cf7ab0e2f101963d328aa62193f7
.git/hooks/commit-msg.sample - 1f74d5e9292979b573ebd59741d46cb93ff391acdd083d340b94370753d92437
.git/hooks/fsmonitor-watchman.sample - e0549964e93897b519bd8e333c037e51fff0f88ba13e086a331592bf801fa1d0
.git/hooks/post-update.sample - 81765af2daef323061dcbc5e61fc16481cb74b3bac9ad8a174b186523586f6c5
.git/hooks/pre-applypatch.sample - e15c5b469ea3e0a695bea6f2c82bcf8e62821074939ddd85b77e0007ff165475
.git/hooks/pre-commit.sample - f9af7d95eb1231ecf2eba9770fedfa8d4797a12b02d7240e98d568201251244a
.git/hooks/pre-merge-commit.sample - d3825a70337940ebbd0a5c072984e13245920cdf8898bd225c8d27a6dfc9cb53
.git/hooks/pre-push.sample - ecce9c7e04d3f5dd9d8ada81753dd1d549a9634b26770042b58dda00217d086a
.git/hooks/pre-rebase.sample - 4febce867790052338076f4e66cc47efb14879d18097d1d61c8261859eaaa7b3
.git/hooks/pre-receive.sample - a4c3d2b9c7bb3fd8d1441c31bd4ee71a595d66b44fcf49ddb310252320169989
.git/hooks/prepare-commit-msg.sample - e9ddcaa4189fddd25ed97fc8c789eca7b6ca16390b2392ae3276f0c8e1aa4619
.git/hooks/push-to-checkout.sample - a53d0741798b287c6dd7afa64aee473f305e65d3f49463bb9d7408ec3b12bf5f
.git/hooks/update.sample - 8d5f2fa83e103cf08b57eaa67521df9194f45cbdbcb37da52ad586097a14d106
.git/info/exclude - 6671fe83b7a07c8932ee89164d1f2793b2318058eb8b98dc5c06ee0a5a3b0ec1
LICENSE - d14655891b354138438a37f460052132fcb00555f62c5f29be732be74b94ff0e
LICENSE-THIRD-PARTY.md - 51cdf5646751388ce88c6f2fb80f1f71b241be6792068602c23f54e1ede7ad40
README.md - 003e90ea68c550cd8ce3383fdf0deef830475e044c76c8bb3f145898c583cc4a
go.mod - 5115ff8fa29d8637c6e7fd5fc99c5c76a909add2a09c0e09379bbaffdb9c147a
go.sum - 588457de154fca2d2dad93574eb1eedf96d1af36dadf4f46ca331be7a9ddf087
main.go - f8ae8089247e577d7e9b2d3435c1a3086b92bfa55320c3a59bf13687c9d15093
main.go~ - 3a1fafc856e3886962162028936ca02b9895e8d2690f646d0e99206d07b06477
e21bb4960769b574e65dc440e5e70cabab363cd4d6b19290f389c77d8333efdc
```

## Installation

The preferred method of installation is using `go install` (as this is
a Golang application). This automates downloading and building Go
applications from source in a secure manner. By default, applications
are copied into `~/go/bin/`.

You must first [install Go](https://golang.org/doc/install). After installing
Go, run the following command to install the application:

```sh
go install gitlab.com/stephen-fox/dhash@latest
```
