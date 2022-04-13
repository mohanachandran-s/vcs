// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/edge-service/cmd/vault-server

go 1.16

require (
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/go-kivik/couchdb/v3 v3.2.8 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/hyperledger/aries-framework-go v0.1.9-0.20220412155017-81442062e607
	github.com/hyperledger/aries-framework-go-ext/component/storage/couchdb v0.0.0-20220330151152-6bbd64bde42e
	github.com/hyperledger/aries-framework-go-ext/component/storage/mongodb v0.0.0-20220330151152-6bbd64bde42e
	github.com/hyperledger/aries-framework-go-ext/component/storage/mysql v0.0.0-20220330151152-6bbd64bde42e
	github.com/hyperledger/aries-framework-go-ext/component/vdr/orb v1.0.0-rc.1
	github.com/hyperledger/aries-framework-go/component/storageutil v0.0.0-20220330140627-07042d78580c
	github.com/hyperledger/aries-framework-go/spi v0.0.0-20220330140627-07042d78580c
	github.com/rs/cors v1.7.0
	github.com/spf13/cobra v1.3.0
	github.com/stretchr/testify v1.7.0
	github.com/trustbloc/edge-core v0.1.8
	github.com/trustbloc/edge-service v0.0.0-00010101000000-000000000000
)

replace github.com/trustbloc/edge-service => ../..
