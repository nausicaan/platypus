# WordPress Update Check

Runs the standard `wp plugin list --update=available` command as well as some custom searches to grab those hard to find plugin updates.

```console
Below is the current list of plugins requiring updates for test.blog.gov.bc.ca.

wpackagist-plugin/gutenberg:14.8.2
wpackagist-plugin/stackable-ultimate-gutenberg-blocks:3.6.3
wpackagist-plugin/styles-and-layouts-for-gravity-forms:4.3.10
wpackagist-plugin/tablepress:2.0.1
bcgov-plugin/events-calendar-pro:6.0.5.1
bcgov-plugin/event-tickets-plus:5.6.4
bcgov-plugin/gravityforms:2.6.8.2
```

## Prerequisite

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

## Build

From the root folder containing *main.go*, use the command that matches your environment:

### Windows & Mac:

```bash
go build -o <build_location>/<program_name> main.go
```

### Linux:

```bash
GOOS=linux GOARCH=amd64 go build -o <build_location>/<program_name> main.go
```

## Run

```bash
./<program_name> <target server> <wordpress path> <wordpress url>
```

Example:

```bash
./upcheck coeurl.dmz /data/www-app/test_blog_gov_bc_ca/current/web/wp test.blog.gov.bc.ca
```

## License
Code is distributed under [The Unlicense](https://github.com/nausicaan/free/blob/main/LICENSE.md) and is part of the Public Domain.
