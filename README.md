Client for Travis-CI
=========================================================

Decidedly unofficial client implementation for Travis CI / Travis Pro API.

Currently supports minimal CRUD for managing:

  * repository settings
  * environment variables

Need something that isn't here? Contributions are [very
welcome](CONTRIBUTING.md)!

## Usage

For travis-ci.org, create a new client with:

    token := "xyz" // your travis-ci.org token
    tc := travis.NewClient(token)

For Travis Pro (travis-ci.com), create a new client with:

    token := "xyz" // your travis-ci.org token
    tc := travis.NewProClient(token)

Once the client is configured, jump on the API!

    repo, err := tc.GetRepository("rjz", "go-travis")
    // etc

## License

MIT
