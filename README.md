# profile parrot

Bring the party to your Slack photo!

## Installing 

Either grab a build on the [releases page](github.com/joshheinrichs/profile-parrot/releases) or run:

```bash
$ go get -u github.com/joshheinrichs/profile-parrot
$ go generate github.com/joshheinrichs/profile-parrot
$ go install github.com/joshheinrichs/profile-parrot
```

Requires [bindata](github.com/jteeuwen/go-bindata)

## Usage

```bash
$ profile-parrot -slack-token SLACK_TOKEN
```

Slack tokens can be generated at https://api.slack.com/custom-integrations/legacy-tokens