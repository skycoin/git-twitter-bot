# git-twitter-bot

Skycoin Twitter bot

## Requirements (for building)

- `Go` version `1.16.x`

## Usage

### Obtain credentials

1. Create twitter developer
   account [https://developer.twitter.com/en/apply/user.html](https://developer.twitter.com/en/apply/user.html)
2. Create a new developer project for your twitter
   account [https://developer.twitter.com/en/docs/projects/overview](https://developer.twitter.com/en/docs/projects/overview)
3. You'll obtain consumer key (API key) and consumer secret (API secret) once your project has been approved. Store them
   somewhere safe.
4. Generate Access Token for your application, make sure that it has both `read` and `write` permissions.

### Config

- Set Environment Variables:

```
   export TW_CONSUMER_KEY="YOUR CONSUMER KEY"
   export TW_CONSUMER_SECRET="YOUR CONSUMER SECRET"
   export TW_ACCESS_TOKEN="YOUR BOT ACCESS TOKEN"
   export TW_ACCESS_TOKEN_SECRET="YOUR BOT ACCESS TOKEN SECRET"
   export TW_TARGET_ORG_URL="YOUR ORG URL example: https://api.github.com/orgs/skycoin/events"
```

### Build

- Clone this repository:

```
git clone https://github.com/SkycoinPro/twitter-bot
```

- Run `make build`

- Binary will be named `twitter-bot` at the root directory of this repository.

*or build Docker image*

- Run `make docker`

### Running

- **Baremetal**

Make sure you've already set the environment variables above.

```
twitter-bot
```

- Docker

```
docker run --rm \
   -e TW_CONSUMER_KEY=<YOUR CONSUMER KEY> \
   -e TW_CONSUMER_SECRET=<YOUR CONSUMER KEY> \
   -e TW_ACCESS_TOKEN=<YOUR ACCESS TOKEN> \
   -e TW_ACCESS_TOKEN_SECRET=<YOUR ACCESS TOKEN SECRET> \
   -e TW_TARGET_ORG_URL=https://api.github.com/orgs/skycoin/events \
   -it git-twitter-bot:latest
```