package utils

type Color struct {
	Name string
	ColorCode string
}

type ColorService struct {
	colors map[int]Color
}

func (cs ColorService) New() *ColorService {
  
  colorMap := make(map[int]Color)

  //Normal
  colorMap[0] = Color{
		Name: "Normal",
		ColorCode: "\x1B[0;0m",
	}
	colorMap[1] = Color{
		Name: "Red",
		ColorCode: "\x1B[0;31m",
  }
  colorMap[2] = Color{
		Name: "Green",
		ColorCode: "\x1B[0;32m",
	}
	colorMap[3] = Color{
		Name: "Yellow",
		ColorCode: "\x1B[0;33m",
  }
  colorMap[4] = Color{
		Name: "Blue",
		ColorCode: "\x1B[0;34m",
	}
	colorMap[5] = Color{
		Name: "Magenta",
		ColorCode: "\x1B[0;35m",
  }
  colorMap[6] = Color{
		Name: "Cyan",
		ColorCode: "\x1B[0;36m",
	}
	colorMap[7] = Color{
		Name: "White",
		ColorCode: "\x1B[0;37m",
  }

  //Bold
	colorMap[8] = Color{
		Name: "Bold Red",
		ColorCode: "\x1B[1;31m",
  }
  colorMap[9] = Color{
		Name: "Bold Green",
		ColorCode: "\x1B[1;32m",
	}
	colorMap[10] = Color{
		Name: "Bold Yellow",
		ColorCode: "\x1B[1;33m",
  }
  colorMap[11] = Color{
		Name: "Bold Blue",
		ColorCode: "\x1B[1;34m",
	}
	colorMap[12] = Color{
		Name: "Bold Magenta",
		ColorCode: "\x1B[1;35m",
  }
  colorMap[13] = Color{
		Name: "Bold Cyan",
		ColorCode: "\x1B[1;36m",
	}
	colorMap[14] = Color{
		Name: "Bold White",
		ColorCode: "\x1B[1;37m",
  }

  //Background
  colorMap[15] = Color{
		Name: "Bold Red",
		ColorCode: "\x1B[41m",
  }
  colorMap[16] = Color{
		Name: "Bold Green",
		ColorCode: "\x1B[42m",
	}
	colorMap[17] = Color{
		Name: "Bold Yellow",
		ColorCode: "\x1B[43m",
  }
  colorMap[18] = Color{
		Name: "Bold Blue",
		ColorCode: "\x1B[44m",
	}
	colorMap[19] = Color{
		Name: "Bold Magenta",
		ColorCode: "\x1B[45m",
  }
  colorMap[20] = Color{
		Name: "Bold Cyan",
		ColorCode: "\x1B[46m",
	}
	colorMap[21] = Color{
		Name: "Bold White",
		ColorCode: "\x1B[47m",
  }

  colorService := &ColorService{
    colors: colorMap,
  }

  return colorService
}

func (cs ColorService) ProcessString(s string) string {
	outString := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '&' {
			colorID := getColorCode(s[i+1])
			if colorID >= 0 {
        outString += cs.colors[colorID].ColorCode
        i++
			}
    } else {
      outString += string(s[i])
    }
  }
  
  return outString
}

func getColorCode(code byte) int {
  switch (code) {
  /* Normal colours */
  case  'k': return 25; break;	/* Black */
  case  'r': return 1;	break;	/* Red */
  case  'g': return 2;	break;	/* Green */
  case  'y': return 3;	break;	/* Yellow */
  case  'b': return 4;	break;	/* Blue */
  case  'm': return 5;	break;	/* Magenta */
  case  'c': return 6;	break;	/* Cyan */
  case  'w': return 7;	break;	/* White */

  /* Bold colours */
  case  'K': return 29; break;  /* Bold black (Just for completeness) */
  case  'R': return 8;	break;	/* Bold red */
  case  'G': return 9;	break;	/* Bold green */
  case  'Y': return 10;	break;	/* Bold yellow */
  case  'B': return 11;	break;	/* Bold blue */
  case  'M': return 12;	break;	/* Bold magenta */
  case  'C': return 13;	break;	/* Bold cyan */
  case  'W': return 14;	break;	/* Bold white */
  
  /* Background colours */
  case  '0': return 24;	break; 	/* Black background */
  case  '1': return 15;	break;	/* Red background */
  case  '2': return 16;	break;	/* Green background */
  case  '3': return 17;	break;	/* Yellow background */
  case  '4': return 18;	break;	/* Blue background */
  case  '5': return 19;	break;	/* Magenta background */
  case  '6': return 20;	break;	/* Cyan background */
  case  '7': return 21;	break;	/* White background */

  /* Misc characters */
  case  '&': return 22;	break;	/* The & character */
  case '\\': return 23;	break;	/* The \ character */
  
  /* Special codes */
  case  'n': return 0;	break;	/* Normal */
  case  'f': return 26;	break;	/* Flash */
  case  'v': return 27; break;	/* Reverse video */
  case  'u': return 28; break;	/* Underline (Only for mono screens) */

  default:   return -1;	break;
  }
  return -1;
}
// #define CNRM  "\x1B[0;0m"
// #define CBLK  "\x1B[0;30m"
// #define CRED  "\x1B[0;31m"
// #define CGRN  "\x1B[0;32m"
// #define CYEL  "\x1B[0;33m"
// #define CBLU  "\x1B[0;34m"
// #define CMAG  "\x1B[0;35m"
// #define CCYN  "\x1B[0;36m"
// #define CWHT  "\x1B[0;37m"
// #define CNUL  ""

// #define BBLK  "\x1B[1;30m"
// #define BRED  "\x1B[1;31m"
// #define BGRN  "\x1B[1;32m"
// #define BYEL  "\x1B[1;33m"
// #define BBLU  "\x1B[1;34m"
// #define BMAG  "\x1B[1;35m"
// #define BCYN  "\x1B[1;36m"
// #define BWHT  "\x1B[1;37m"

// #define BKBLK  "\x1B[40m"
// #define BKRED  "\x1B[41m"
// #define BKGRN  "\x1B[42m"
// #define BKYEL  "\x1B[43m"
// #define BKBLU  "\x1B[44m"
// #define BKMAG  "\x1B[45m"
// #define BKCYN  "\x1B[46m"
// #define BKWHT  "\x1B[47m"

// #define CAMP  "&"
// #define CSLH  "\\"

// #define CUDL  "\x1B[4m"	/* Underline ANSI code */
// #define CFSH  "\x1B[5m"	/* Flashing ANSI code.  Change to #define CFSH "" if
//                          * you want to disable flashing colour codes
//                          */
// #define CRVS  "\x1B[7m" /* Reverse video ANSI code */

// const char *COLOURLIST[] = {CNRM, CRED, CGRN, CYEL, CBLU, CMAG, CCYN, CWHT,
// 			    BRED, BGRN, BYEL, BBLU, BMAG, BCYN, BWHT,
// 			    BKRED, BKGRN, BKYEL, BKBLU, BKMAG, BKCYN, BKWHT,
// 			    CAMP, CSLH, BKBLK, CBLK, CFSH, CRVS, CUDL, BBLK };

// #define MAX_COLORS 30

// int isnum(char s)
// {
//   return( (s>='0') && (s<='9') );
// }

// int is_colour(char code)
// {
//   switch (code) {
//   /* Normal colours */
//   case  'k': return 25; break;	/* Black */
//   case  'r': return 1;	break;	/* Red */
//   case  'g': return 2;	break;	/* Green */
//   case  'y': return 3;	break;	/* Yellow */
//   case  'b': return 4;	break;	/* Blue */
//   case  'm': return 5;	break;	/* Magenta */
//   case  'c': return 6;	break;	/* Cyan */
//   case  'w': return 7;	break;	/* White */

//   /* Bold colours */
//   case  'K': return 29; break;  /* Bold black (Just for completeness) */
//   case  'R': return 8;	break;	/* Bold red */
//   case  'G': return 9;	break;	/* Bold green */
//   case  'Y': return 10;	break;	/* Bold yellow */
//   case  'B': return 11;	break;	/* Bold blue */
//   case  'M': return 12;	break;	/* Bold magenta */
//   case  'C': return 13;	break;	/* Bold cyan */
//   case  'W': return 14;	break;	/* Bold white */
  
//   /* Background colours */
//   case  '0': return 24;	break; 	/* Black background */
//   case  '1': return 15;	break;	/* Red background */
//   case  '2': return 16;	break;	/* Green background */
//   case  '3': return 17;	break;	/* Yellow background */
//   case  '4': return 18;	break;	/* Blue background */
//   case  '5': return 19;	break;	/* Magenta background */
//   case  '6': return 20;	break;	/* Cyan background */
//   case  '7': return 21;	break;	/* White background */

//   /* Misc characters */
//   case  '&': return 22;	break;	/* The & character */
//   case '\\': return 23;	break;	/* The \ character */
  
//   /* Special codes */
//   case  'n': return 0;	break;	/* Normal */
//   case  'f': return 26;	break;	/* Flash */
//   case  'v': return 27; break;	/* Reverse video */
//   case  'u': return 28; break;	/* Underline (Only for mono screens) */

//   default:   return -1;	break;
//   }
//   return -1;
// }

// void proc_color(char *inbuf, int colour)
// {
//   register int j = 0, p = 0;
//   int k, max, c = 0;
//   char out_buf[32768];

//   if (inbuf[0] == '\0')
//     return;

//   while (inbuf[j] != '\0') {
//     if ((inbuf[j]=='\\') && (inbuf[j+1]=='c')
//         && isnum(inbuf[j + 2]) && isnum(inbuf[j + 3])) {
//       c = (inbuf[j + 2] - '0')*10 + inbuf[j + 3]-'0';
//       j += 4;
//     } else if ((inbuf[j] == '&') && !(is_colour(inbuf[j + 1]) == -1)) {
//       c = is_colour(inbuf[j + 1]);
//       j += 2;
//     } else {
//       out_buf[p] = inbuf[j];
//       j++;
//       p++;
//       continue;
//     }
//     if (c > MAX_COLORS)
//       c = 0;
//     max = strlen(COLOURLIST[c]);
//     if (colour || max == 1)
//       for (k = 0; k < max; k++) {
//         out_buf[p] = COLOURLIST[c][k];
// 	p++;
//       }
//   }

//   out_buf[p] = '\0';

//   strcpy(inbuf, out_buf);
// }
