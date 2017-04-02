# profile parrot 
Bring the party to your Slack photo!

[![Travis Ci Status](https://travis-ci.org/joshheinrichs/profile-parrot.svg?branch=master)](https://travis-ci.org/joshheinrichs/profile-parrot.svg?branch=master)
[![AppVeyor Status](https://ci.appveyor.com/api/projects/status/a2bhg1b1upsa4wxg/branch/master?svg=true)](https://ci.appveyor.com/project/joshheinrichs/profile-parrot/branch/master)

## Installing 

Either grab a build on the [releases page](github.com/joshheinrichs/profile-parrot/releases) or run:

```bash
$ go get -u github.com/joshheinrichs/profile-parrot
$ go generate github.com/joshheinrichs/profile-parrot
$ go install github.com/joshheinrichs/profile-parrot
```

Requires [bindata](github.com/jteeuwen/go-bindata) to build.

## Usage

```bash
$ profile-parrot -slack-token SLACK_TOKEN
```

Slack tokens can be generated at https://api.slack.com/custom-integrations/legacy-tokens
