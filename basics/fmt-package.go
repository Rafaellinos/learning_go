package basics

import "fmt"

/*
fmt verbs

%v = default
%t = true or false
%c = character
%X = hex
%U = unicode format
%e = scientific notation
*/

/*
fmt Escape sequences
\\ backslash
\' single quote
\" double quote
\n new line
\u or \U unicode (2b and 4b)
\x raw bytes
*/

func mainFmt() string {
	return fmt.Sprintf("Rafael %c %v", 'v', "lino")
}
