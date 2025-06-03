## ⚡ Go Live Reload with Air

### 📦 Install Air

```bash
go install github.com/air-verse/air@latest
```

> ✅ Make sure your `$GOPATH/bin` is in your `PATH`.

To check:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

You can now run:

```bash
air --version
```

---

### ⚙️ Setup `air` Config (Optional)

Generate `.air.toml` config file (optional, but helpful for customization):

```bash
air init
```

Edit the `.air.toml` file to suit your project layout if needed.

---

### 🚀 Run your project with Air

From your project root, just run:

```bash
air
```

> 🛠️ Air will automatically restart your app when you change `.go` files or templates.

---

### 📁 Example Project Structure

```
my-go-app/
├── main.go
├── go.mod
├── .air.toml  (optional)
```
