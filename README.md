<h1 align="center">
  SecNex DNS Server
</h1>

<div align="center">
    <img src="https://img.shields.io/github/downloads/secnex/DNS-Server/total?style=for-the-badge" />
    <img src="https://img.shields.io/github/last-commit/secnex/DNS-Server?color=%231BCBF2&style=for-the-badge" />
    <img src="https://img.shields.io/github/issues/secnex/DNS-Server?style=for-the-badge" />
</div>

<br />

# 📝 Table of Contents <a name="table-of-contents"></a>

- [📝 Table of Contents](#table-of-contents)
- [📂 Project Structure](#project-structure)
- [🧐 About](#about)
- [🏁 Getting Started](#getting_started)
  - [💻 Prerequisites](#prerequisites)
  - [🚀 Installing](#installing)
- [📝 License](#license)
- [📝 Acknowledgements](#acknowledgements)

# 📂 Project Structure <a name="project-structure"></a>

```
📂 DNS-Server
├── 📂 src
│   ├── 📂 resolver
│   │   ├── 📄 domains.go
│   │   ├── 📄 records.go
│   │   ├── 📄 resolver.go
│   │   └── 📄 utils.go
│   ├── 📂 sql
│   ├── 📄 go.mod
│   ├── 📄 go.sum
│   └── 📄 main.go
├── 📂 server
├── 📄 .gitignore
├── 📄 build.sh
├── 📄 LICENSE
├── 📄 README.md
├── 📄 secnex.dns.service
└── 📄 SECURITY.md
```

# 🧐 About <a name="about"></a>

This project is a nameserver written in Golang. It is designed to be a simple, lightweight, and easy to use your own nameserver. It is designed to be used with [ByteDNS](https://dns.bytesentinel.io) but can be used with any DNS provider.

# 📝 License <a name="license"></a>

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.