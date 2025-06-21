# Usage

## 🚀 Commands

1. create
   Generate a new full stack project with optional flags for customization.

```bash
sersi create
```

this generates a wizard to create a full stack application. If you would like to create either a frontend or just a backend, use the subcommands below.

### subcommands:

-   backend
    Scaffold a backend project with optional flags for customization.

```bash
sersi create backend [flags]
```

Available Flags for `backend`:

| Flag        | Type   | Description                              | Optional |
| ----------- | ------ | ---------------------------------------- | -------- |
| --name      | string | Name of the project                      | Yes      |
| --lang      | string | Language to use (ts, nodejs, go, python) | Yes      |
| --framework | string | Framework to use (fastapi, express, gin) | Yes      |
| --database  | string | Database to use (postgres, mongodb)      | Yes      |

Example:

```bash
sersi create backend --name my-backend --framework fastapi --database postgresql --lang py
```

-   frontend
    Scaffold a frontend project with optional flags for customization.

```bash
sersi create frontend [flags]
```

Available Flags for `frontend`:

| Flag        | Type   | Description                                    | Optional |
| ----------- | ------ | ---------------------------------------------- | -------- |
| --name      | string | Name of the project                            | Yes      |
| --framework | string | Framework to use (react, vue, vanilla)         | Yes      |
| --css       | string | CSS framework (tailwind, sass, bootstrap, css) | Yes      |
| --lang      | string | Language to use (ts, js)                       | Yes      |

Example:

```bash
sersi create frontend --name my-site --framework vue --css bootstrap --lang js
```

2. build
   Build a project from a YAML configuration file.

```bash
sersi build -f path/to/sersi.yaml
```

| Flag | Required | Description                  |
| ---- | -------- | ---------------------------- |
| -f   | ✅       | Path to the YAML config file |

🧪 Example sersi.yaml

```yaml
name: sersi-app
structure: polyrepos
scaffold:
    frontend:
        framework: react
        css: tailwind
        language: typescript
    backend:
        language: go
        framework: gin
        database: postgresql
```

## 📄 YAML Reference

These are the currently supported values for each flag & field in `sersi.yaml`:

### `frontend`

#### `framework`

-   `react`
-   `svelte`
-   `vue`
-   `vanilla`

#### `css`

-   `tailwind`
-   `bootstrap`
-   `css`

#### `language`

-   `typescript` or `ts`
-   `javascript` or `js`

### `backend`

#### `language`

-   `typescript` or `ts`
-   `javascript` or `js`
-   `go`
-   `python`

#### `framework`

-   `express`
-   `gin` (Go)
-   `chi` (Go)
-   `fastapi` (Python)

#### `database`

- `postgres` (coming soon)
- `mongodb` (coming soon)
-   `none`

3. version
   Show the version of the CLI.

```bash
sersi version
```

4. help
   Display help information for any command.

```bash
sersi help
sersi help create
```

## 🗒 Notes

-   Frameworks and languages must be among the supported options.
-   You can always override parts of a YAML config with flags when using build.
