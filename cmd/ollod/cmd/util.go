package cmd

import (
	"regexp"
	"strings"

	color "github.com/fatih/color"
	// "github.com/ignite/cli/ignite/pkg/cliui/view/accountview"
	// "github.com/ignite/cli/ignite/pkg/cliui/icons"
	// "github.com/ignite/cli/ignite/pkg/cliui/clispinner"
	// "github.com/ignite/cli/ignite/pkg/cliui/entrywriter"
	// "github.com/ignite/cli/ignite/pkg/markdownviewer"
	// "github.com/ignite/cli/ignite/pkg/cliui/colors"
	// cosmosanl "github.com/ignite/cli/ignite/pkg/cosmosanalysis/module"
	"github.com/spf13/cobra"
)

var (
	StyleHeading = color.New(color.FgHiGreen, color.Bold).SprintFunc()
	StyleParam   = color.New(color.FgHiYellow).SprintFunc()
	StyleTx      = color.New(color.FgHiRed, color.Bold).SprintFunc()
	StyleQuery   = color.New(color.FgHiYellow, color.Bold).SprintFunc()
	StyleSubcmd  = color.New(color.FgHiBlue, color.Bold).SprintFunc()
	StyleFlags   = color.New(color.FgBlue, color.Bold).SprintFunc()
	StyleError   = color.New(color.FgRed, color.Bold).SprintFunc()
	StyleWarning = color.New(color.FgHiYellow, color.Bold).SprintFunc()
	StyleInfo    = color.New(color.FgHiBlue, color.Bold).SprintFunc()
	StyleMoniker = color.New(color.FgHiGreen).SprintFunc()
	StyleAmt     = color.New(color.FgHiMagenta).SprintFunc()
	StyleCoin    = color.New(color.FgHiCyan).SprintFunc()
	StyleHelp    = color.New(color.FgHiWhite, color.Bold).SprintFunc()
	StyleFlag    = color.New(color.FgBlue).SprintFunc()
	StyleType    = color.New(color.FgGreen, color.Italic).SprintFunc()
	styleRepl    = func(st string) string {
		return strings.NewReplacer(
			`Usage:`, `{{StyleHeading "Usage:"}}`,
			`Aliases:`, `{{StyleHeading "Aliases:"}}`,
			`Available Commands:`, `{{StyleHeading "Available Commands:"}}`,
			`Global Flags:`, `{{StyleHeading "Global Flags:"}}`,
			`Flags:`, `{{StyleHeading "Flags:"}}`,
			`[command]`, `{{StyleParam "[command]"}}`,
			`[flags]`, `{{StyleFlags "[flags]"}}`,
			// `[amount]`, `{{StyleAmt "[amount]"}}`,
			// `[coin]`, `{{StyleCoin "[coin]"}}`,
			// `string`, `{{StyleType "string"}}`,
			// `query`, `{{StyleQuery "query (q)"}}`,
			// `tx`, `{{StyleTx "tx (t)"}}`,
			// `help`, `{{StyleTx "help"}}`,
			// `Error`, `{{StyleError "Error"}}`,
			// `Warning`, `{{StyleWarning "Warning"}}`,
		).Replace(st)
	}
	styleAll = func() {
		cobra.AddTemplateFunc("StyleHeading", StyleHeading)
		cobra.AddTemplateFunc("StyleParam", StyleParam)
		cobra.AddTemplateFunc("StyleTx", StyleTx)
		cobra.AddTemplateFunc("StyleError", StyleError)
		cobra.AddTemplateFunc("StyleWarning", StyleWarning)
		cobra.AddTemplateFunc("StyleInfo", StyleInfo)
		cobra.AddTemplateFunc("StyleSubcmd", StyleSubcmd)
		cobra.AddTemplateFunc("StyleFlags", StyleFlags)
		cobra.AddTemplateFunc("StyleQuery", StyleQuery)
		cobra.AddTemplateFunc("StyleMoniker", StyleMoniker)
	}
	styleCmd = func(cmd *cobra.Command) *cobra.Command {
		styleAll()
		usageTmpl := cmd.UsageTemplate()
		newUsage := styleRepl(usageTmpl)
		re := regexp.MustCompile(`(?m)^Flags:\s*$`)
		usageTmpl = re.ReplaceAllLiteralString(usageTmpl, `{{StyleHeading "Flags:"}}`)
		cmd.SetUsageTemplate(newUsage)
		return cmd
	}
)
var (
	fgMagenta = color.New(color.FgHiMagenta, color.Bold).SprintFunc()
	fgBlue    = color.New(color.FgHiBlue, color.Italic, color.Bold).SprintFunc()
	fgDesc    = color.New(color.Italic, color.Faint).SprintFunc()
	fgBold    = color.New(color.Bold).SprintFunc()
)
var (
	Earth = "🌍"

	// OK is an OK mark.
	OK = color.New(color.FgGreen).SprintFunc()("✔")
	// NotOK is a red cross mark
	NotOK = color.New(color.FgRed).SprintFunc()("✘")
	// Bullet is a bullet mark
	Bullet = color.New(color.FgYellow).SprintFunc()("⋆")
	// Info is an info mark
	Info = color.New(color.FgYellow).SprintFunc()("𝓲")
)
