# Contributing to Sersi CLI

Thank you for your interest in contributing to **Sersi CLI**! 🚀  
We welcome community involvement in improving the CLI and building the ecosystem around it.

---

## 🧱 Project Structure Overview

core/
├── src/
│ ├── cmd/
│ │ ├── build/
│ │ ├── create/
│ │ └── version/
│ ├── common/
│ ├── core/
│ ├── docs/
│ ├── tea/
│ ├── templates/
│ ├── test/
│ ├── utils/
│ └── validator/
├── go.mod
├── go.sum
├── main.go
├── README.md
├── LICENSE
├── .gitignore
└── .prettierrc

---

## 🛠️ Getting Started

### 1. Fork the Repository

````bash
git clone https://github.com/YOUR_USERNAME/sersi-cli
cd sersi-cli

### 2. Install Dependencies
Ensure you have Go 1.20+

Install dependencies:

```bash
go mod tidy
````

### 3. Build the CLI

```bash
go build
go install
```

### 4. Run the CLI

```bash
sersi <command> <flags>
```

## 🔁 Contribution Flow

- Create a feature branch: `feat/your-feature-name`
- Commit using clear messages (feat:, fix:, docs:)
- Open a PR against main with a description of what you're solving or adding.
- Add screenshots or code samples if relevant.
- One of the maintainers will review your PR.

## 📦 Adding Templates or Languages

- New frameworks go in `templates/frontend/`
- CSS options go in `templates/css/`
- Base shared logic in `templates/common/`
- Make sure they support both TypeScript and JavaScript unless otherwise scoped.

## 📝 Documentation

- Update `docs/` with new features
- Add examples to `examples/`
- Keep README updated with latest info

## 💬 Communication

For questions, ideas, or larger contributions, feel free to:

- Open an issue
- Start a discussion
- Join our Discord (link coming soon)

## 🧾 License

By contributing, you agree that your contributions will be licensed under the MIT license for the community edition and may be dual-licensed or commercialized in the pro version.

Thanks again! 💜
