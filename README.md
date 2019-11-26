# sandbox
scratch repo

https://github.com/golang/go

https://research.swtch.com/vgo-module



Multiple-Module Repositories

Developers may also find it useful to maintain a collection of modules in a single source code repository. We want vgo to support this possibility. In general, there is already wide variation in how different developers, teams, projects, and companies apply source control, and we do not believe it is productive to impose a single mapping like “one repository equals one module” onto all developers. Having some flexibility here should also help vgo adapt as best practices around souce control continue to change.

In the major subdirectory convention, v2/ contains the module "my/thing/v2". A natural extension is to allow subdirectories not named for major versions. For example, we could add a blue/ subdirectory that contains the module "my/thing/blue", confirmed by a blue/go.mod file with that module path. In this case, the source control commit tags addressing that module would take the form blue/v1.x.x. Similarly, the tag blue/v2.x.x would address the blue/v2/ subdirectory. The existence of the blue/go.mod file excludes the blue/ tree from the outer my/thing module.

In the Go project, we intend to explore using this convention to allow repositories like golang.org/x/text to define multiple, independent modules. This lets us retain the convenience of coarse-grained source control but still promote different subtrees to v1 at different times.	



https://golang.org/cmd/go/#hdr-Module_configuration_for_non_public_modules

The 'go list' command provides information about the main module and the build list. For example:

go list -m              # print path of main module
go list -m -f={{.Dir}}  # print root directory of main module
go list -m all          # print build list


The 'go mod' command provides other functionality for use in maintaining and understanding modules and go.mod files. See 'go help mod'.

The -mod build flag provides additional control over updating and use of go.mod.

If invoked with -mod=readonly, the go command is disallowed from the implicit automatic updating of go.mod described above. Instead, it fails when any changes to go.mod are needed. This setting is most useful to check that go.mod does not need updates, such as in a continuous integration and testing system. The "go get" command remains permitted to update go.mod even with -mod=readonly, and the "go mod" commands do not take the -mod flag (or any other build flags).

If invoked with -mod=vendor, the go command assumes that the vendor directory holds the correct copies of dependencies and ignores the dependency descriptions in go.mod.


go get github.com/gorilla/mux@latest    # same (@latest is default for 'go get')
go get github.com/gorilla/mux@v1.6.2    # records v1.6.2
go get github.com/gorilla/mux@e3702bed2 # records v1.6.2
go get github.com/gorilla/mux@c856192   # records v0.0.0-20180517173623-c85619274f5d
go get github.com/gorilla/mux@master    # records current meaning of master


