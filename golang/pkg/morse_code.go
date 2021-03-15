package pkg

var MORSE_DECODER = map[string]string{
	// Letters
	".-": "A",
	"-...": "B",
	"-.-.": "C",
	"-..": "D",
	".": "E",
	"..-.":"F",
	"--.":"G",
	"....":"H",
	"..":"I",
	".---":"J",
	"-.-":"K",
	".-..":"L",
	"--":"M",
	"-.":"N",
	"---":"O",
	".--.":"P",
	"--.-":"Q",
	".-.":"R",
	"...":"S",
	"-":"T",
	"..-":"U",
	"...-":"V",
	".--":"W",
	"-..-":"X",
	"-.--":"Y",
	"--..":"Z",
	// Numbers
	"-----":"0",
	".----":"1",
	"..---":"2",
	"...--":"3",
	"....-":"4",
	".....":"5",
	"-....":"6",
	"--...":"7",
	"---..":"8",
	"----.":"9",
	// Symbols
	"---...":":",
	"--..--":",",
	"-.--.":"(",
	"-.--.-":")",
	"-.-.--":"!",
	"-.-.-.":";",
	"-..-.":"/",
	"-...-":"=",
	"-....-":"-",
	".----.":"'",
	".--.-.":"@",
	".-.-.":"+",
	".-.-.-":".",
	".-..-.":"\"",
	".-...":"&",
	"..--.-":"_",
	"..--..":"?",
	"...---...":"SOS",
	"...-..-":"$",
}

var MORSE_ENCODER = map[string]string{
	// Letters
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
	// Numbers
	"0": "-----",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	// Symbols
	":": "---...",
	",": "--..--",
	"(": "-.--.",
	")": "-.--.-",
	"!": "-.-.--",
	";": "-.-.-.",
	"/": "-..-.",
	"=": "-...-",
	"-": "-....-",
	"'": ".----.",
	"@": ".--.-.",
	"+": ".-.-.",
	".": ".-.-.-",
	".-..-.":"\"",
	"&": ".-...",
	"_": "..--.-",
	"?": "..--..",
	"SOS": "...---...",
	"$": "...-..-",
}