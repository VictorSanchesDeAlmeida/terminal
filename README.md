# DevShell 🚀

Um terminal interativo personalizado para desenvolvedores, construído em Go, com autocomplete inteligente para comandos de desenvolvimento comuns.

## 📋 Descrição

DevShell é um terminal que oferece uma experiência aprimorada para desenvolvedores, fornecendo autocomplete contextual para uma variedade de ferramentas de desenvolvimento. O projeto usa a biblioteca `go-prompt` para criar uma interface de linha de comando rica e interativa.

## ✨ Funcionalidades

- **Terminal interativo** com prompt dinâmico que acompanha o diretório atual
- **Atalhos de teclado estilo terminal** (Linux/Windows/macOS):
  - `Ctrl + C`: interrompe comando em execução (sem fechar o DevShell)
  - `Ctrl + L`: limpa a tela
  - `Ctrl + A` / `Ctrl + E`: início/fim da linha
  - `Ctrl + U` / `Ctrl + K`: apaga antes/depois do cursor
  - `Ctrl + W`: apaga palavra anterior
  - `Ctrl + P` / `Ctrl + N`: histórico anterior/próximo
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

- **Built-ins reais do shell** (não abrem subprocesso):
  - `cd` (incluindo `cd ..` e `cd ~`)
  - `pwd`
  - `clear`
  - `exit` e `quit`
  - autocomplete contextual de diretórios para `cd`

- **Execução de comandos** multiplataforma:
  - Windows: `cmd /C`
  - macOS/Linux: `sh -c`
- **Entrada interativa em comandos externos**: comandos que pedem input agora aceitam digitação normalmente

## 🛠️ Tecnologias

- **Go 1.25.4**
- **go-prompt** - Para interface interativa
- **cmd/sh** - Para execução de comandos externos no Windows/macOS/Linux

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
- Windows, macOS ou Linux

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

1. Após executar, você verá um prompt personalizado: `caminho/atual >> `
2. Digite comandos normalmente ou use as teclas de seta para navegar no histórico
3. Use **Tab** para autocomplete - digite algumas letras e pressione Tab para ver sugestões
4. Digite `exit`, `quit` ou **Ctrl+C** para sair

### Exemplos de autocomplete

- Digite `git` + Tab → verá opções como `git clone`, `git status`, `git commit`
- Digite `npm` + Tab → verá `npm install`, `npm run dev`, `npm test`
- Digite `expo` + Tab → verá `expo start`, `expo build`, `expo run:android`
- Digite `files` + Tab → verá todos os arquivos e pastas do diretório atual
- Digite `cd src` + Tab → verá diretórios compatíveis com o prefixo

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