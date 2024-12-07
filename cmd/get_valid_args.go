package cmd

import "github.com/spf13/cobra"

func GetValidArgs(
	cmd *cobra.Command,
	args []string,
	toComplete string,
) ([]string, cobra.ShellCompDirective) {
	notes, err := storageInstance.GetNotes()
	cobra.CheckErr(err)

	var result []string
	for _, note := range notes {
		result = append(result, note.Name)
	}

	return result, cobra.ShellCompDirectiveNoFileComp
}
