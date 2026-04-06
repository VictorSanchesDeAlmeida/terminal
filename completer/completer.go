package completer

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/c-bata/go-prompt"
)

var staticSuggestions = []prompt.Suggest{
	{Text: "cd", Description: "Change directory"},
	{Text: "ls", Description: "List files"},
	{Text: "dir", Description: "List files (Windows)"},
	{Text: "pwd", Description: "Show current directory"},
	{Text: "clear", Description: "Clear terminal"},
	{Text: "exit", Description: "Exit DevShell"},
	{Text: "quit", Description: "Exit DevShell"},

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

	{Text: "npm install", Description: "Install dependencies"},
	{Text: "npm run dev", Description: "Run dev script"},
	{Text: "npm run build", Description: "Build project"},
	{Text: "npm run start", Description: "Run start script"},
	{Text: "npm run lint", Description: "Run linter"},
	{Text: "npm run test", Description: "Run tests"},

	{Text: "node", Description: "Run Node.js"},
	{Text: "node index.js", Description: "Run main file"},
	{Text: "node --watch", Description: "Run in watch mode"},

	{Text: "go run .", Description: "Run current Go module"},
	{Text: "go build", Description: "Build Go project"},
	{Text: "go test ./...", Description: "Run all tests"},

	{Text: "expo start", Description: "Start Expo dev server"},
	{Text: "expo start --clear", Description: "Start and clear cache"},
	{Text: "expo install", Description: "Install compatible dependencies"},
	{Text: "expo doctor", Description: "Check project health"},
	{Text: "expo prebuild", Description: "Generate native code"},
	{Text: "expo run:android", Description: "Run on Android"},
	{Text: "expo run:ios", Description: "Run on iOS"},
	{Text: "expo start --web", Description: "Run on web"},

	{Text: "eas build", Description: "Build app"},
	{Text: "eas build -p android", Description: "Build Android"},
	{Text: "eas build -p ios", Description: "Build iOS"},
	{Text: "eas submit", Description: "Submit to store"},
	{Text: "eas update", Description: "Publish OTA update"},

	{Text: "ssh-keygen -t ed25519 -C \"seu_email@example.com\"", Description: "Generate SSH key for GitHub"},
	{Text: "eval $(ssh-agent -s)", Description: "Init SSH agent"},
	{Text: "ssh-add ~/.ssh/id_ed25519", Description: "Add SSH key to agent"},
	{Text: "cat ~/.ssh/id_ed25519.pub", Description: "Show public SSH key"},
	{Text: "pbcopy < ~/.ssh/id_ed25519.pub", Description: "Copy SSH key to clipboard (macOS)"},
	{Text: "ssh -T git@github.com", Description: "Test SSH connection to GitHub"},
}

func Complete(d prompt.Document) []prompt.Suggest {
	rawText := d.TextBeforeCursor()
	text := strings.TrimSpace(rawText)
	word := d.GetWordBeforeCursor()

	if strings.HasPrefix(text, "cd ") || text == "cd" {
		return completePathsForCd(text)
	}

	if contextual := completeCommandTail(rawText); len(contextual) > 0 {
		return contextual
	}

	dynamic := handleDynamic(word)
	if len(dynamic) > 0 {
		return dynamic
	}

	filtered := prompt.FilterHasPrefix(staticSuggestions, text, true)
	if len(filtered) > 0 {
		return filtered
	}

	return prompt.FilterHasPrefix(staticSuggestions, word, true)
}

func completeCommandTail(rawText string) []prompt.Suggest {
	if strings.TrimSpace(rawText) == "" {
		return nil
	}

	parts := strings.Fields(rawText)
	if len(parts) < 1 {
		return nil
	}
	endsWithSpace := strings.HasSuffix(rawText, " ")

	command := parts[0]
	fullPrefix := command + " "
	typedTail := parts[1:]

	var suggestions []prompt.Suggest
	seen := make(map[string]struct{})
	for _, item := range staticSuggestions {
		if !strings.HasPrefix(strings.ToLower(item.Text), strings.ToLower(fullPrefix)) {
			continue
		}

		tail := strings.TrimPrefix(item.Text, fullPrefix)
		tailParts := strings.Fields(tail)
		if len(tailParts) == 0 {
			continue
		}

		suggestText := ""
		if endsWithSpace {
			if len(typedTail) > len(tailParts) {
				continue
			}

			matched := true
			for i := range typedTail {
				if !strings.EqualFold(typedTail[i], tailParts[i]) {
					matched = false
					break
				}
			}
			if !matched {
				continue
			}

			remaining := tailParts[len(typedTail):]
			if len(remaining) == 0 {
				continue
			}
			suggestText = strings.Join(remaining, " ")
		} else {
			if len(typedTail) == 0 {
				continue
			}

			completeCount := len(typedTail) - 1
			if completeCount > len(tailParts) {
				continue
			}

			matched := true
			for i := 0; i < completeCount; i++ {
				if !strings.EqualFold(typedTail[i], tailParts[i]) {
					matched = false
					break
				}
			}
			if !matched {
				continue
			}

			candidate := strings.Join(tailParts[completeCount:], " ")
			partial := typedTail[len(typedTail)-1]
			if partial != "" && !strings.HasPrefix(strings.ToLower(candidate), strings.ToLower(partial)) {
				continue
			}

			suggestText = candidate
		}

		if suggestText == "" {
			continue
		}

		if _, ok := seen[suggestText]; ok {
			continue
		}
		seen[suggestText] = struct{}{}

		suggestions = append(suggestions, prompt.Suggest{
			Text:        suggestText,
			Description: item.Description,
		})
	}

	return suggestions
}

func handleDynamic(word string) []prompt.Suggest {
	var suggestions []prompt.Suggest

	if strings.HasPrefix("local", word) {
		dir, err := os.Getwd()
		if err != nil {
			return suggestions
		}

		suggestions = append(suggestions, prompt.Suggest{
			Text:        dir,
			Description: "Diretório atual",
		})
	}

	if strings.HasPrefix("home", word) {
		home, err := os.UserHomeDir()
		if err != nil {
			return suggestions
		}

		suggestions = append(suggestions, prompt.Suggest{
			Text:        home,
			Description: "Home do usuário",
		})
	}

	if strings.HasPrefix("files", word) {
		dir, err := os.Getwd()
		if err != nil {
			return suggestions
		}

		files, err := os.ReadDir(dir)
		if err != nil {
			return suggestions
		}

		for _, file := range files {
			description := "Arquivo"
			if file.IsDir() {
				description = "Pasta"
			}

			suggestions = append(suggestions, prompt.Suggest{
				Text:        file.Name(),
				Description: description,
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

func completePathsForCd(text string) []prompt.Suggest {
	parts := strings.Fields(text)
	target := ""
	if len(parts) > 1 {
		target = parts[1]
	}

	searchBase := "."
	prefix := target

	if target != "" {
		dir := filepath.Dir(target)
		if dir != "." {
			searchBase = dir
		}
		prefix = filepath.Base(target)
	}

	entries, err := os.ReadDir(searchBase)
	if err != nil {
		return []prompt.Suggest{{Text: "..", Description: "Go to parent directory"}}
	}

	suggestions := []prompt.Suggest{
		{Text: "..", Description: "Go to parent directory"},
		{Text: "~", Description: "Go to home directory"},
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		if prefix != "" && !strings.HasPrefix(strings.ToLower(name), strings.ToLower(prefix)) {
			continue
		}

		suggestPath := name
		if searchBase != "." {
			suggestPath = filepath.Join(searchBase, name)
		}

		suggestions = append(suggestions, prompt.Suggest{
			Text:        suggestPath,
			Description: "Diretório",
		})
	}

	return suggestions
}
