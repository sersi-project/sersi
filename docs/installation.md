# Installation

Welcome to **Sersi** — a powerful tool to scaffold full-stack applications quickly with your preferred tech stack.

---

## 🔁 Installation Options

### Option 1: Install via Homebrew (macOS/Linux)

```bash
brew tap sersi-project/sersi
brew install sersi
```

### Option 2: Install via Scoop (Windows)

```bash
scoop bucket add sersi-project
scoop install sersi
```

### Option 3: Install via `npm`

```bash
npm install -g sersi
#or
npx sersi@latest
```

### Option 4: Download Prebuilt Binary (Windows, macOS, Linux)

1. Visit the [Releases](https://github.com/sersi-project/sersi/releases) page
2. Download the latest binary for your platform

-   Windows: `sersi-windows-amd64/sersi.exe`
-   macOS (Intel): `sersi-darwin-amd64/sersi`
-   macOS (Apple Silicon): `sersi-darwin-arm64/sersi`
-   Linux: `sersi-linux-amd64/sersi`

3. Make the binary executable

```bash
chmod +x sersi
```

4. Move the binary to a directory in your PATH

```bash
sudo mv sersi /usr/local/bin/sersi
```

5. Verify the installation

```bash
sersi --version
```

## Need help?

If you need help, please refer to the [Usage](./USAGE.md) or [contact us](https://sersi.dev/help).
