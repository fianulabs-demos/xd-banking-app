# xd-banking-app

Demo Bancorp Application

--

[![Badges](http://badges.governance-system.34.132.74.168.sslip.io/badges?project=badgercorp&repository=xd-banking-app)](http://ui-badger.default.34.132.74.168.sslip.io/badgercorp/xd-banking-app)


### Attestor testing


```bash
badger attestor wrap --key k8s://default/cosign --rekor-url http://rekor.rekor-system.34.132.74.168.sslip.io --step ko.build --output-file ko.build.json -- ko build github.com/badgercorp/xd-banking-api/cmd/api --image-label version=$VERSION -t $VERSION --push=false --local
```

badger attestor wrap --key k8s://default/cosign --rekor-url http://rekor.rekor-system.34.132.74.168.sslip.io --output-file result.json --step sast.download -- curl -fsSL https://raw.githubusercontent.com/ZupIT/horusec/main/deployments/scripts/install.sh



          badger attestor wrap --key $BADGER_KEY --rekor-url $REKOR_HOST --output-file result.json --step deploy.policy \
          -- badger policy register resource -c repository -p $POLICY_FILE
          
          badger attestor hash --path result.json
          
          badger attestor wrap --step debug --key $BADGER_KEY --rekor-url $REKOR_HOST --output-file result.json -- \
            curl -fsSL https://raw.githubusercontent.com/ZupIT/horusec/main/deployments/scripts/install.sh | bash -s latest