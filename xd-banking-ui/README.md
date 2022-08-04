# Quasar App (xd-banking-ui)

A Quasar Project

## Install the dependencies
```bash
yarn
# or
npm install
```

### Start the app in development mode (hot-code reloading, error reporting, etc.)
```bash
quasar dev
```


### Lint the files
```bash
yarn lint
# or
npm run lint
```


### Format the files
```bash
yarn format
# or
npm run format
```



### Build the app for production
```bash
quasar build
```

### Customize the configuration
See [Configuring quasar.config.js](https://v2.quasar.dev/quasar-cli-webpack/quasar-config-js).


badger attestor wrap --step docker.build --key k8s://default/cosign --rekor-url http://rekor.rekor-system.34.132.74.168.sslip.io --output-file result.json -- \
docker build \
--tag "gcr.io/$PROJECT_ID/$ARTIFACT:$VERSION" \
--build-arg GITHUB_SHA="$GITHUB_SHA" \
--build-arg GITHUB_REF="$GITHUB_REF" \
.
