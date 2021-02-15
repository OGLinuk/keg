# Golang KEG API (with `keg` and `kn` tools)

See <https://github.com/afkworks/specs> for more information.

## Implementation Plan

Getting a good prototype created and available quickly for alpha testing
with a broad group is therefore the initial goal. This will allow rapid,
organic enhancements ensuring maximum usability before implementing the
main release.

Prototype development on `kn` is in Shell and Perl. Final version will
be in Go.

## Extension Support 

Extensions use the UNIX philosophy to allow them to be written in any
language that will execute on the target operating system.

## Environment Variables

The `kn` tool does not depend on any environment variables itself
(although extensions might have their own dependency). Instead, it uses
the [`conf-go`](https://github.com/rwxrob/conf-go) package with defaults
(which creates matching configuration files under `~/.local/config`).

*[Note that `KN` and `KNPATH` were once used, but no longer.]*
