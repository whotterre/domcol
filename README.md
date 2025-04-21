# 🎨 DomCol

DomCol is a command-line utility that extracts **dominant colors** from an image file.  
It iss written in **Go** and powered by [spf13/cobra](https://github.com/spf13/cobra) for its CLI interface.

![Go Badge](https://img.shields.io/badge/made%20with-Go-00ADD8?logo=go&logoColor=white)
![CLI Badge](https://img.shields.io/badge/cli-cobra-blue?logo=terminal&logoColor=white)

---

## 📦 Installation

```bash
go install github.com/whotterre/domcol@latest
```

> Ensure that `$GOPATH/bin` is in your system's `PATH` so you can run `domcol` from anywhere.

---

## 🚀 Usage

```bash
domcol --imgpath <path-to-image> [--rgba] [--hex] [--top N]
```

### 🔧 Flags

| Flag        | Description                                   |
|-------------|-----------------------------------------------|
| `--imgpath`, `-i` | **(Required)** Path to the image file         |
| `--rgba`     | Output dominant colors in `rgb(r, g, b)` format |
| `--hex`      | Output dominant colors in hex (`#RRGGBB`) format |
| `--top`      | Number of dominant colors to display (default: 5) |

---

## 📷 Example

```bash
domcol -i ./sunset.png --hex --top 3
```
The output consists of the [color swatch](https://www.coreldraw.com/en/tips/swatch-definition/#:~:text=As%20mentioned%20before%2C%20a%20color,a%20specific%20color%20on%20it.), hex or rgb value and the number of pixels with the associated color
```
██████ #FF5733 (1023)
██████ #33A1FF (789)
██████ #1C1C1C (456)
```

You can also display RGB:

```bash
domcol -i ./sunset.png --rgba --top 3
```

```
██████ rgb(255, 87, 51) (1023)
██████ rgb(51, 161, 255) (789)
██████ rgb(28, 28, 28) (456)
```

---

## ✅ Supported Formats

- PNG
- JPG / JPEG

More formats may be supported in future releases.

---

## 📂 Project Structure

```
domcol/
├── cmd/           # Cobra CLI commands
├── pkg/domcol/    # Core image processing logic
├── go.mod
├── go.sum
└── main.go
```

---

## 🧠 How It Works

- The image is loaded and decoded using the standard library.
- All pixels are analyzed, and the RGB values are bucketed into a frequency map.
- The most frequent colors are then sorted and displayed using terminal color escape codes.

---

## ❤️ Author

Made with :heart: by [Iwegbu Jeddy](https://github.com/whotterre)