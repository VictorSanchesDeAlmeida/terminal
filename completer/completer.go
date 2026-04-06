package completer

import (
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func Complete(d prompt.Document) []prompt.Suggest {

	word := d.GetWordBeforeCursor()

	// 🔥 1. comandos dinâmicos (prioridade alta)
	dynamic := handleDynamic(word)
	if len(dynamic) > 0 {
		return dynamic
	}

	commands := []prompt.Suggest{
		{Text: "git", Description: "Git commands"},
		{Text: "npm", Description: "Node package manager"},
		{Text: "node", Description: "Run Node.js"},
		{Text: "npx", Description: "Run packages"},
		{Text: "cd", Description: "Change directory"},
		{Text: "ls", Description: "List files"},
		{Text: "dir", Description: "List files (Windows)"},
		{Text: "pwd", Description: "Show current directory"},
		{Text: "clear", Description: "Clear terminal"},
	}

	gitCommands := []prompt.Suggest{
		{Text: "git clone", Description: "Clone repository"},
		{Text: "git status", Description: "Show status"},
		{Text: "git add", Description: "Stage changes"},
		{Text: "git commit", Description: "Commit changes"},
		{Text: "git push", Description: "Push changes"},
		{Text: "git pull", Description: "Pull changes"},
		{Text: "git checkout", Description: "Switch branch"},
		{Text: "git branch", Description: "List branches"},
		{Text: "git merge", Description: "Merge branches"},
		{Text: "git log", Description: "Show commit history"},
	}

	npmCommands := []prompt.Suggest{
		{Text: "npm install", Description: "Install dependencies"},
		{Text: "npm run dev", Description: "Run script"},
		{Text: "npm run start:dev", Description: "Run development server"},
		{Text: "npm run build", Description: "Build project"},
		{Text: "npm run start", Description: "Format code"},
		{Text: "npm run lint", Description: "Run linter"},
		{Text: "npm run test", Description: "Run tests"},
	}

	nodeCommands := []prompt.Suggest{
		{Text: "node", Description: "Run Node.js"},
		{Text: "node index.js", Description: "Run main file"},
		{Text: "node --watch", Description: "Run in watch mode"},
	}

	fsCommands := []prompt.Suggest{
		{Text: "cd", Description: "Change directory"},
		{Text: "cd ..", Description: "Go back"},
		{Text: "mkdir", Description: "Create directory"},
		{Text: "rm", Description: "Remove file"},
		{Text: "rmdir", Description: "Remove directory"},
	}

	goCommands := []prompt.Suggest{
		{Text: "go run", Description: "Run Go file"},
		{Text: "go build", Description: "Build Go project"},
		{Text: "go test", Description: "Run Go tests"},
	}

	expoCommands := []prompt.Suggest{
		{Text: "expo start", Description: "Start Expo dev server"},
		{Text: "expo start --clear", Description: "Start and clear cache"},
		{Text: "expo install", Description: "Install compatible dependencies"},
		{Text: "expo doctor", Description: "Check project health"},
		{Text: "expo prebuild", Description: "Generate native code"},
		{Text: "expo run:android", Description: "Run on Android"},
		{Text: "expo run:ios", Description: "Run on iOS"},
		{Text: "expo start --web", Description: "Run on web"},
		{Text: "expo login", Description: "Login Expo"},
		{Text: "expo logout", Description: "Logout Expo"},
		{Text: "expo whoami", Description: "Current user"},
	}

	easCommands := []prompt.Suggest{
		{Text: "eas build", Description: "Build app"},
		{Text: "eas build -p android", Description: "Build Android"},
		{Text: "eas build -p ios", Description: "Build iOS"},
		{Text: "eas submit", Description: "Submit to store"},
		{Text: "eas update", Description: "Publish OTA update"},
		{Text: "eas login", Description: "Login EAS"},
		{Text: "eas whoami", Description: "Current user"},
	}

	sshKeyGitCommands := []prompt.Suggest{
		{Text: "ssh-keygen -t ed25519 -C \"seu_email@example.com\"", Description: "Generate SSH key for GitHub"},
		{Text: "eval $(ssh-agent -s)", Description: "Init SSH agent"},
		{Text: "ssh-add ~/.ssh/id_ed25519", Description: "Add SSH key to agent"},
		{Text: "cat ~/.ssh/id_ed25519.pub", Description: "Show public SSH key"},
		{Text: "pbcopy < ~/.ssh/id_ed25519.pub", Description: "Copy SSH key to clipboard (macOS)"},
		{Text: "ls ~/.ssh", Description: "List SSH keys"},
		{Text: "ssh -T git@github.com", Description: "Test SSH connection to GitHub"},
	}

	suggestions := []prompt.Suggest{}

	suggestions = append(suggestions, commands...)
	suggestions = append(suggestions, gitCommands...)
	suggestions = append(suggestions, npmCommands...)
	suggestions = append(suggestions, nodeCommands...)
	suggestions = append(suggestions, fsCommands...)
	suggestions = append(suggestions, goCommands...)
	suggestions = append(suggestions, expoCommands...)
	suggestions = append(suggestions, easCommands...)
	suggestions = append(suggestions, sshKeyGitCommands...)

	return prompt.FilterHasPrefix(suggestions, word, true)
}

func handleDynamic(word string) []prompt.Suggest {
	var suggestions []prompt.Suggest

	// 👉 local → diretório atual
	if strings.HasPrefix("local", word) {
		dir, _ := os.Getwd()

		suggestions = append(suggestions, prompt.Suggest{
			Text:        dir,
			Description: "Diretório atual",
		})
	}

	// 👉 home → pasta do usuário
	if strings.HasPrefix("home", word) {
		home, _ := os.UserHomeDir()

		suggestions = append(suggestions, prompt.Suggest{
			Text:        home,
			Description: "Home do usuário",
		})
	}

	// 👉 files → arquivos e pastas do diretório atual
	if strings.HasPrefix("files", word) {
		dir, _ := os.Getwd()
		files, _ := os.ReadDir(dir)

		for _, file := range files {
			suggestions = append(suggestions, prompt.Suggest{
				Text:        file.Name(),
				Description: "Arquivo ou pasta",
			})
		}
	}

	if strings.HasPrefix("exit", word) {
		suggestions = append(suggestions, prompt.Suggest{
			Text:        "exit",
			Description: "Sair do terminal",
		})
	}

	return suggestions
}
