# Github User Activity CLI
https://roadmap.sh/projects/github-user-activity
Github User Activity is a CLI tool for fetching user activity from Github by providing the username.

## Features

- Fetch recent activities of a Github user
- Display activities such as push events, issue events, watch events, fork events, and create events

## Installation

Firstly clone the repo

```bash
git clone https://github.com/omarisadev/github-activity
cd ./github-activity/
```

Then install it to PATH *optional*

```bash
go install .
github-activity <username>
```

## Building

To build the project, you need to have Go installed. You can install it from the [official website](https://golang.org/doc/install).

Once Go is installed, build the project using the following command:

```bash
go build -o github-activity
```

## Running

After building the project, you can run the CLI tool with:

```bash
./github-activity <username>
```

Replace `<username>` with the Github username you want to fetch activities for. For example:

```bash
./github-activity omarisadev
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
