# Make / Just integration

The pattern: **gokazi is the stop button, not the start button.** Your Makefile or justfile owns starting (because that's project-specific). gokazi owns identifying and stopping.

## Makefile

```makefile
.PHONY: dev
## Start dev servers
dev:
	@(cd web && npm run dev &)
	@(cd api && go run . &)

.PHONY: stop
## Stop everything gokazi can match
stop:
	@gokazi stop web || true
	@gokazi stop api || true

.PHONY: status
## Show task status
status:
	@gokazi list
```

`|| true` keeps `make stop` idempotent — stopping an already-stopped task returns a non-zero exit.

## justfile

```just
# Start dev servers
dev:
    (cd web && npm run dev &)
    (cd api && go run . &)

# Stop everything gokazi can match
stop:
    -gokazi stop web
    -gokazi stop api

# Show task status
status:
    gokazi list
```

The leading `-` in just makes the recipe ignore non-zero exits, same idea as `|| true` in make.

## Why not let gokazi start things?

`gokazi` deliberately does not spawn processes. The argument: starting is project-specific (env vars, working directory, build steps, hot reload), and your existing build tool already does it well. Bolting another start mechanism on top splits the source of truth.
