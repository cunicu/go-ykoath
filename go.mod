// SPDX-FileCopyrightText: 2018 Joern Barthel <joern.barthel@kreuzwerker.de>
// SPDX-License-Identifier: Apache-2.0

module cunicu.li/go-ykoath/v2

go 1.23.0

toolchain go1.24.5

require (
	cunicu.li/go-iso7816 v0.8.6
	golang.org/x/crypto v0.40.0
)

require (
	github.com/ebfe/scard v0.0.0-20241214075232-7af069cabc25 // test-only
	github.com/stretchr/testify v1.10.0 // test-only
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
