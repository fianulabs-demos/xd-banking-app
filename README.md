# xd-banking-app

Demo Bancorp Application

--

[![Badges](http://badges.governance-system.34.132.74.168.sslip.io/badges?project=badgercorp&repository=xd-banking-app)](http://ui-badger.default.34.132.74.168.sslip.io/badgercorp/xd-banking-app)


### Attestor testing


```bash
badger attestor wrap --key k8s://default/cosign --rekor-url http://rekor.rekor-system.34.132.74.168.sslip.io --step ko.build --output-file ko.build.json -- ko build github.com/badgercorp/xd-banking-api/cmd/api --image-label version=$VERSION -t $VERSION --push=false --local
```