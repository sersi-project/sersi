# Usage

## 🚀 Commands

1. create
   Generate a new project with optional flags for customization.

```bash
sersi create [flags]
```

Alternatively, you can use flags to do everything in one go.

Available Flags:

| Flag        | Type   | Description                                    | Optional |
| ----------- | ------ | ---------------------------------------------- | -------- |
| --name      | string | Name of the project                            | Yes      |
| --framework | string | Framework to use (react, vue, vanilla)         | Yes      |
| --css       | string | CSS framework (tailwind, sass, bootstrap, css) | Yes      |
| --lang      | string | Language to use (ts, js)                       | Yes      |

Example:

```bash
sersi create --name my-site --framework vue --css bootstrap --lang js
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
scaffold:
  frontend:
    framework: react
  css: tailwind
  language: typescript
  opts:
    - material-ui
    - prettier
```

## 📄 YAML Reference

These are the currently supported values for each field in `sersi.yaml`:

### `frontend`

#### `framework`

- `react`
- `svelte`
- `vue`
- `vanilla`

#### `css`

- `tailwind`
- `bootstrap`
- `css`

#### `language`

- `typescript` or `ts`
- `javascript` or `js`

#### `opts`

A list of additional tooling or UI frameworks. Supported:

- `material-ui`
- `prettier`

## 🔐 Extended YAML Options [Sersi Pro](https://sersi.dev/pro)

In addition to standard fields, Sersi Pro supports:

### `backend`

#### `language`

- `typescript` or `ts`
- `javascript` or `js`
- `go`

#### `framework`

- `deno`
- `express`
- `gin` (Go)

#### `database`

- `postgres`
- `sqlite`
- `mongodb`

### `devops`

#### `docker`

- `true` or `false`

#### `cicd`

- `github-actions`
- `gitlab`

#### `monitoring`

- `prometheus`

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

- Frameworks and languages must be among the supported options.
- You can always override parts of a YAML config with flags when using build.
