# Golang KEG API (with `keg` and `kn` tools)

See <https://github.com/afkworks/specs> for more information.

## Implementation Plan

Getting a good prototype created and available quickly for alpha testing
with a broad group is therefore the initial goal. This will allow rapid,
organic enhancements ensuring maximum usability before implementing the
main release.

Prototype development on `kn` is in Shell and Perl and Go.

The core `kn` tool (and built-in Actions) are implemented in Go which
boasts the largest developer base and adoption of any modern, compiled
language, flexible modularity, rapid development, and is the fastest
modern language to compile, share, and learn.

Extensions use the UNIX philosophy to allow them to be written in any
language that will execute on the target operating system.

## Environment Variables

The `kn` tool does not depend on any environment variables itself
(although extensions might have their own dependency).

*[Note that `KN` and `KNPATH` were once used, but no longer.]*

## The `.kn/actions` Directory

The `$KN/.kn/actions` directory (within the currently active knowledge node)
may contain Actions that extend or override the defaults. This provides
a modular way for knowledge workers to extend and customize any Action
available to `kn` and ensures that anyone reviewing or using the node on
KEG will also have access to those customizations.

## Frequently Asked Questions

These are questions. They are frequent.

### How can I share my knowledge nodes? 

That would be what [`keg`](https://github.com/afkworks/keg) is for.

### Why only JSON for configuration?

The `kn` tool uses the `kn config` action for all changes to the
configuration file (like `git`) and therefore JSON is used instead of
alternatives like YAML or TOML or HCL.

YAML is meant for human maintenance and is therefore not good when it
needs to be re-written back out (frequently losing structure, comments,
whitespace, etc.). JSON is intended to be written and read by apps and
computers but reasonably readable by humans. It was never intended to be
be *maintained* by humans using text editors (which is why line returns
are not supported). The design priority of JSON is efficiency, not human
maintainability.

