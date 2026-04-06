# DevShell 🚀

Um terminal interativo personalizado para desenvolvedores, construído em Go, com autocomplete inteligente para comandos de desenvolvimento comuns.

## 📋 Descrição

DevShell é um terminal que oferece uma experiência aprimorada para desenvolvedores, fornecendo autocomplete contextual para uma variedade de ferramentas de desenvolvimento. O projeto usa a biblioteca `go-prompt` para criar uma interface de linha de comando rica e interativa.

## ✨ Funcionalidades

- **Terminal interativo** com prompt customizável que mostra o diretório atual
- **Autocomplete inteligente** para comandos populares de desenvolvimento:
  - 🔧 **Git**: clone, status, add, commit, push, pull, checkout, merge, etc.
  - 📦 **NPM**: install, run scripts, build, test, lint
  - 🟩 **Node.js**: execução de scripts, watch mode
  - ⚛️ **Expo**: start, build, install, doctor, run android/ios
  - 📱 **EAS**: build, submit, update, login
  - 🐹 **Go**: run, build, test
  - 📁 **Sistema de arquivos**: cd, ls, mkdir, rm, etc.

- **Sugestões dinâmicas**:
  - `local`: mostra o diretório atual
  - `home`: mostra o diretório home do usuário
  - `files`: lista arquivos e pastas do diretório atual

- **Execução de comandos** através do PowerShell no Windows

## 🛠️ Tecnologias

- **Go 1.25.4**
- **go-prompt** - Para interface interativa
- **PowerShell** - Para execução de comandos no Windows

## 📁 Estrutura do Projeto

```
terminal/
├── main.go              # Ponto de entrada da aplicação
├── go.mod               # Dependências do módulo Go
├── go.sum               # Checksums das dependências
├── completer/
│   └── completer.go     # Lógica de autocomplete
└── executor/
    └── executor.go      # Execução de comandos
```

## 🚀 Instalação e Uso

### Pré-requisitos

- Go 1.25+ instalado
- Windows com PowerShell

### Instalação

1. Clone o repositório:
```bash
git clone <url-do-repositorio>
cd terminal
```

2. Instale as dependências:
```bash
go mod download
```

3. Execute o projeto:
```bash
go run main.go
```

### Como usar

1. Após executar, você verá um prompt personalizado: `C:\caminho\atual >> `
2. Digite comandos normalmente ou use as teclas de seta para navegar no histórico
3. Use **Tab** para autocomplete - digite algumas letras e pressione Tab para ver sugestões
4. Digite **Ctrl+C** para sair

### Exemplos de autocomplete

- Digite `git` + Tab → verá opções como `git clone`, `git status`, `git commit`
- Digite `npm` + Tab → verá `npm install`, `npm run dev`, `npm test`
- Digite `expo` + Tab → verá `expo start`, `expo build`, `expo run:android`
- Digite `files` + Tab → verá todos os arquivos e pastas do diretório atual

## 🔧 Personalização

Você pode facilmente adicionar novos comandos editando o arquivo `completer/completer.go`:

```go
newCommands := []prompt.Suggest{
    {Text: "seu-comando", Description: "Descrição do comando"},
    // adicione mais comandos aqui
}
```

## 🎯 Casos de Uso

- **Desenvolvimento Web**: Comandos rápidos para npm, node, git
- **Desenvolvimento Mobile**: Suporte completo para Expo e EAS
- **Desenvolvimento Go**: Comandos go integrados
- **DevOps**: Navegação rápida em diretórios e arquivos
- **Produtividade**: Autocomplete reduz digitação e erros

## 🤝 Contribuição

Contribuições são bem-vindas! Para contribuir:

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

## 🔗 Dependências

- [go-prompt](https://github.com/c-bata/go-prompt) - Biblioteca para terminal interativo

---

**DevShell** - Transformando a experiência de linha de comando para desenvolvedores! 💻✨