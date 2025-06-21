## YAML Configuration File

SERSI uses a `sersi.yaml` file for project configuration. This file should be placed in your project root directory.

### Basic Structure

```yaml
name: your-project-name
structure: monorepo  # or polyrepos
hooks:
  pre: true
  post: true
scaffold:
  frontend:
    framework: react
    css: tailwind
    lang: js
  backend:
    lang: js
    framework: express
    database: postgresql
  devops:
    ci: github
    docker: true
    monitoring: prometheus
```

## Project Structure Options

| Option | Description |
|--------|-------------|
| `monorepo` | Single repository containing all components |
| `polyrepos` | Multiple repositories for different components |

## Frontend Configuration

### Framework Options

| Framework | Description | Language Support |
|-----------|-------------|------------------|
| `react` | React.js framework | JS, TS |
| `vue` | Vue.js framework | JS, TS |
| `svelte` | Svelte framework | JS, TS |
| `vanilla` | Vanilla JavaScript | JS, TS |

### CSS Framework Options

| CSS Framework | Description |
|---------------|-------------|
| `tailwind` | Tailwind CSS utility framework |
| `bootstrap` | Bootstrap CSS framework |
| `css` | Plain CSS |

### Language Options

| Language | Description |
|----------|-------------|
| `js` or `javascript` | JavaScript |
| `ts` or `typescript` | TypeScript |

### Frontend Configuration Example

```yaml
scaffold:
  frontend:
    framework: react
    css: tailwind
    lang: typescript
```

## Backend Configuration

### Language Options

| Language | Description | Framework Support |
|----------|-------------|-------------------|
| `js` or `javascript` | Node.js | Express, Fastify |
| `ts` or `typescript` | TypeScript | Express, Fastify |
| `go` | Go | Gin, Chi |
| `py` or `python` | Python | FastAPI |

### Framework Options

#### Node.js Frameworks
| Framework | Description |
|-----------|-------------|
| `express` | Express.js web framework |
| `fastify` | Fastify web framework |

#### Go Frameworks
| Framework | Description |
|-----------|-------------|
| `gin` | Gin web framework |
| `chi` | Chi router |

#### Python Frameworks
| Framework | Description |
|-----------|-------------|
| `fastapi` | FastAPI web framework |

### Database Options

| Database | Description | Status |
|----------|-------------|--------|
| `postgresql` | PostgreSQL database | Available |
| `mongodb` | MongoDB database | Available |
| `none` | No database | Available |

### Backend Configuration Example

```yaml
scaffold:
  backend:
    lang: go
    framework: gin
    database: postgresql
```
## Pro Features

SERSI Pro extends the open-source CLI with premium features:

### Scaffold Store
- Save and reuse project configurations
- Share templates with team members
- Version control for scaffold configurations

### Custom Hooks
- Advanced pre and post-hook capabilities
- Custom script execution
- Integration with external tools

### Team Features
- Shared private templates
- Team collaboration tools
- Advanced project management

## Examples

### Full-Stack React Application

```yaml
name: my-react-app
structure: monorepo
hooks:
  pre: true
  post: true
scaffold:
  frontend:
    framework: react
    css: tailwind
    lang: typescript
  backend:
    lang: go
    framework: gin
    database: postgresql
  devops:
    ci: github
    docker: true
    monitoring: prometheus
```

### Vue.js Frontend Only

```yaml
name: my-vue-app
structure: monorepo
scaffold:
  frontend:
    framework: vue
    css: bootstrap
    lang: javascript
```

### Python Backend with FastAPI

```yaml
name: my-python-api
structure: polyrepos
scaffold:
  backend:
    lang: python
    framework: fastapi
    database: mongodb
  devops:
    ci: gitlab
    docker: true
```

### Pro Example with Custom Configuration

```yaml
name: enterprise-app
structure: polyrepos
hooks:
  pre: true
  post: true
scaffold:
  frontend:
    framework: react
    css: tailwind
    lang: typescript
  backend:
    lang: go
    framework: chi
    database: postgresql
  devops:
    ci: github
    docker: true
    monitoring: prometheus
```