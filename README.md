# orion

Orion is a toolchain for building, publishing and consuming Arma 3 scripted compositions. These are currently used primarily on BI's Public Zeus servers which have composition scripting enabled.

Orion has a few main target functions:

- [x] Bundling scripts into a composition format
- [ ] Linting and formatting SQF files
- [ ] Preprocessing SQF files
- [ ] Downloading, managing and using scripted compositions
- [ ] Automating the release of compositions to the steam workshop

Work has only just begun, and even "completed" functions are still considered alpha, and breaking changes _will_ happen before 1.0.

# Installing for testing

`go install github.com/endigma/orion@latest` should work on most platforms, releases will come soonish.

You'll need a `orion.toml` to get started. Populate it in your project folder with content like this:

```toml
name = "My Project Name"
author = "My Name"

[[compositions]]
name = "Composition 1"
category = "My Project Category"
initFile = "path/relative/to/orion.toml/to/script1"

[[compositions]]
name = "Composition 2"
category = "My Project Category"
initFile = "path/relative/to/orion.toml/to/script2"
```

And run `orion -p orion.toml -o [desired output folder]`. This functionality will eventually be moved under a subcommand (likely `orion bundle`)