# RTMS CLI Auto-Completion

RTMS CLI supports command auto-completion for Bash, Zsh, fish, and PowerShell. This feature can significantly improve your productivity by reducing typing and helping you discover available commands and options.

## Generating Completion Scripts

RTMS CLI uses cobra, which provides built-in functionality to generate completion scripts. You can generate these scripts using the `rtmscli completion` command.

### Bash

To generate the auto-completion script for Bash:

```bash
rtmscli completion bash > rtmscli_completion.bash
```

To use it in the current session:

```bash
source rtmscli_completion.bash
```

To make it permanent, add this line to your `~/.bashrc` file:

```bash
source /path/to/rtmscli_completion.bash
```

### Zsh

For Zsh, you might need to enable command completion if you haven't already:

```zsh
autoload -Uz compinit && compinit
```

Then generate and use the completion script:

```zsh
rtmscli completion zsh > "${fpath[1]}/_rtmscli"
```

You may need to start a new shell for this setup to take effect.

### Fish

For fish shell:

```fish
rtmscli completion fish > ~/.config/fish/completions/rtmscli.fish
```

### PowerShell

For PowerShell:

```powershell
rtmscli completion powershell > rtmscli_completion.ps1
```

To use it, you need to include it in your PowerShell profile:

```powershell
echo '. /path/to/rtmscli_completion.ps1' >> $PROFILE
```

## Using Auto-Completion

Once you've set up auto-completion, you can use it by typing `rtmscli` followed by a partial command or option, then pressing the Tab key. For example:

```
rtmscli app<Tab>
```

This will complete to `rtmscli appliances`.

You can also use Tab to see available options:

```
rtmscli appliances <Tab>
```

This will show you all the subcommands available under `appliances`.

## Updating Auto-Completion

If new commands are added to RTMS CLI, you'll need to regenerate the completion script to include these new commands. Simply run the generation command again and source the new script.

## Troubleshooting

If auto-completion isn't working as expected:

1. Ensure you've sourced the completion script correctly.
2. Check that your shell supports command completion.
3. Make sure you're using the latest version of RTMS CLI.
4. Try regenerating the completion script.

If problems persist, please report an issue on the RTMS CLI GitHub repository.

## Further Reading

For more detailed information about command completion in general:

- Bash: `man bash`, search for "Programmable Completion"
- Zsh: `man zshcompsys`
- Fish: `man fish-completion`
- PowerShell: `Get-Help about_Tab_Expansion`

Remember, effective use of command completion can significantly speed up your workflow with RTMS CLI!
