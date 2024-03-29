# Contributing

> DISCLAIMER:
>
> The contents of this repository have been migrated into [go-vela/server](https://github.com/go-vela/server) or [go-vela/worker](https://github.com/go-vela/worker).
>
> This was done as a part of [go-vela/community#394](https://github.com/go-vela/community/issues/394) and [go-vela/community#395](https://github.com/go-vela/community/issues/395) to deliver [on a proposal](https://github.com/go-vela/community/blob/master/proposals/2021/08-25_repo-structure.md).

We'd love to accept your contributions to this project! There are just a few guidelines you need to follow.

## Bugs

Bug reports should be opened up as [issues](https://help.github.com/en/github/managing-your-work-on-github/about-issues) on the [go-vela/community](https://github.com/go-vela/community) repository!

## Feature Requests

Feature Requests should be opened up as [issues](https://help.github.com/en/github/managing-your-work-on-github/about-issues) on the [go-vela/community](https://github.com/go-vela/community) repository!

## Pull Requests

**NOTE: We recommend you start by opening a new issue describing the bug or feature you're intending to fix. Even if you think it's relatively minor, it's helpful to know what people are working on.**

We are always open to new PRs! You can follow the below guide for learning how you can contribute to the project!

## Getting Started

### Prerequisites

* [Review the commit guide we follow](https://chris.beams.io/posts/git-commit/#seven-rules) - ensure your commits follow our standards
* [Docker](https://docs.docker.com/install/) - building block for local development
* [Docker Compose](https://docs.docker.com/compose/install/) - start up local development
* [Golang](https://golang.org/dl/) - for source code and [dependency management](https://github.com/golang/go/wiki/Modules)

### Setup

* [Fork](/fork) this repository

* Clone this repository to your workstation:

```bash
# clone the project
git clone git@github.com:go-vela/mock.git $HOME/go-vela/mock
```

* Navigate to the repository code:

```bash
# change into the project directory
cd $HOME/go-vela/mock
```

* Point the original code at your fork:

```bash
# add a remote branch pointing to your fork
git remote add fork https://github.com/your_fork/mock
```

### Development

* Navigate to the repository code:

```bash
# change into the project directory
cd $HOME/go-vela/mock
```

* Write your code and tests to implement the changes you desire.
  * Please be sure to [follow our commit rules](https://chris.beams.io/posts/git-commit/#seven-rules)
  * Please address linter warnings appropriately. If you are intentionally violating a rule that triggers a linter, please annotate the respective code with `nolint` declarations [[docs](https://golangci-lint.run/usage/false-positives/)]. we are using the following format for `nolint` declarations:

    ```go
    // nolint:<linter(s)> // <short reason>
    ```
  
    Example:

    ```go
    // nolint:gocyclo // legacy function is complex, needs simplification
    func superComplexFunction() error {
      // ..
    }
    ```

    Check the [documentation for more examples](https://golangci-lint.run/usage/false-positives/).

* Test the repository code (ensures your changes don't break existing functionality):

```bash
# execute the `test` target with `make`
make test
```

* Clean the repository code (ensures your code meets the project standards):

```bash
# execute the `test` target with `make`
make clean
```

* Push to your fork:

```bash
# push your code up to your fork
git push fork master
```

* Open a pull request!
  * For the title of the pull request, please use the following format for the title:

    ```text
    feat(wobble): add hat wobble
    ^--^^------^  ^------------^
    |   |         |
    |   |         +---> Summary in present tense.
    |   +---> Scope: a noun describing a section of the codebase (optional)
    +---> Type: chore, docs, feat, fix, refactor, or test.
    ```

    * feat: adds a new feature (equivalent to a MINOR in Semantic Versioning)
    * fix: fixes a bug (equivalent to a PATCH in Semantic Versioning)
    * docs: changes to the documentation
    * refactor: refactors production code, eg. renaming a variable; doesn't change public API
    * test: adds missing tests, refactors tests; no production code change
    * chore: updates something without impacting the user (ex: bump a dependency in package.json or go.mod); no production code change

    If a code change introduces a breaking change, place ! suffix after type, ie. feat(change)!: adds breaking change. correlates with MAJOR in semantic versioning.

Thank you for your contribution!
