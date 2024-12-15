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

And run `orion bundle` in the directory containing `orion.toml`. 

# Copyright

Orion is licenced under the MIT License.

Orion embeds code from [Enhanced Zeus Modules](https://github.com/expung3d/Enhanced-Zeus-Modules) for compliance testing. The license of these files, including the minified version, is that of the upstream repository as of [c451dd88](https://github.com/expung3d/Enhanced-Zeus-Modules/tree/c451dd88270940911b03573c52579d11f029f686). It is unlikely this version of EZM will ever need to be updated, as it's only used as a large sample of SQF source code.