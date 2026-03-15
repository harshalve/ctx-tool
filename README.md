# ctx-tool (Context Tool) 🚀

A lightweight, developer-focused CLI tool built in **Go** to eliminate the mental tax of context switching. `ctx-tool` captures your active workspace state—browser links, terminal tasks, and Git branches—and restores them with a single command.

---

## 📋 Table of Contents
- [The Problem](#-the-problem)
- [Key Features](#-key-features)
- [Installation](#-installation)
- [Shell Integration](#-shell-integration)
- [Usage](#-usage)
- [Technical Architecture](#-technical-architecture)

---

## 🧠 The Problem
Context switching is one of the biggest productivity killers for software engineers. Jumping between a Jira ticket, a PR review, and a feature branch usually involves:
1. Re-opening several browser tabs (Docs, GitHub, Jira).
2. Remembering which Git branch was active.
3. Recalling the specific terminal commands or notes for that task.

`ctx-tool` handles the "rehydration" of your workspace so you can focus on writing code.

---

## ✨ Key Features
- **Deep Context Saving:** Interactively save URLs and terminal commands with contextual notes.
- **Git State Persistence:** Automatically detects and switches to the correct Git branch on restore.
- **Clickable Directory Links:** Provides `file://` protocol links to jump straight to your project in Finder or VS Code.

---

## 🛠️ Installation

### 1. Prerequisites
- **Go** (1.20 or higher)
- **Git**

### 2. Build and Install
The project uses a `Makefile` to maintain high code quality. The installation process automatically runs formatting, linting, and tests before installing the binary.

```bash
git clone [https://github.com/harshalve/ctx-tool.git](https://github.com/harshalve/ctx-tool.git)
cd ctx-tool
make install

## Shell Integration

To allow `ctx` to change your terminal's directory, add this helper function to your `~/.zshrc` or `~/.bashrc`:

```bash
# Helper function for ctx-tool
ctx-go() {
    # 1. Restore the context (Tabs/Branch/Notes)
    ctx restore "$1"
    
    # 2. Get the project path and change directory
    local target_dir=$(ctx path "$1")
    if [ -d "$target_dir" ]; then
       cd "$target_dir"
    else
        echo "⚠️  Note: Project directory not found, staying in current folder."
    fi
}


## Usage

### Save a Context
Navigate to your project directory and run:
```bash
ctx save <project-name>
```
The tool will prompt you for links and terminal tasks

### Restore a Context
```bash
ctx restore <project-name>
```
This will open all the links in your default browser and cd you into the project directory and switch to the correct branch

### Use the Shell Integration
```bash
ctx-go <project-name>
```
This will restore the context and change the directory

### List all Contexts
```bash
ls ~/.ctx
```


## Technical Architecture

- Language: Go
- Data Storage: JSON files in `~/.ctx`
- System Interaction: 
    - Uses `os/exec` to interact with Git and the OS default browser.
    - Uses `text/tabwriter` for formatting the output.


## Contributing
I follow an industry-standard development lifecycle.
1. Create a feature branch (git checkout -b feat/my-feature)
2. Make changes
3. Run `make precommit` to format, lint, and test the code
4. Commit the changes (git commit -m "feat: my feature")
5. Push the changes (git push origin feat/my-feature)
6. Sign your commits with GPG
7. Create a Pull Request


Built with ❤️ by Harshal
