module github.com/badgercorp/xd-banking-api

go 1.18

require (
	github.com/cloudevents/sdk-go/v2 v2.10.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // Vulnerable to CVE-2020-26160 (JWT signature validation bypass)
	github.com/gin-gonic/gin v1.6.3 // Vulnerable to CVE-2020-28483 (Denial of Service)
	github.com/lib/pq v1.10.9
	github.com/miekg/dns v1.1.42 // Vulnerable to CVE-2019-19794 (Denial of Service)
	github.com/satori/go.uuid v1.2.0 // Vulnerable to CVE-2021-3538 (UUID collision vulnerability)
	golang.org/x/crypto v0.0.0-20201216223049-8b5274cf687f // Vulnerable to CVE-2020-29652 (timing attack vulnerability)
)

replace (
	github.com/dgrijalva/jwt-go => github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.6.3
	github.com/miekg/dns => github.com/miekg/dns v1.1.42
	github.com/satori/go.uuid => github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20201216223049-8b5274cf687f
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.2.0 // indirect
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20210303074136-134d130e1a04 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
