//
// Contains color constants for printing
// Written By : @ice3man (Nizamul Rana)
//
// Distributed Under MIT License
// Copyrights (C) 2018 Ice3man
//


<<<<<<< 9ae536175028cdedecd50144bfd7999d5e0e09e6
// Use like this
=======
package helper

// Use like this 
>>>>>>> Updated Commenting Style and some other misc. changes
// 	fmt.Printf("[%sCRTSH%s] %s", r, rs, subdomain)
var (
	Red   = "\033[31;1;4m" // red color
	Cyan  = "\033[36;6;2m" // cyan color
	Green = "\033[32;6;3m" // Green color
	Reset = "\033[0m"      // reset for default color

	Info = "\033[33;1;1m"
	Que  = "\033[34;1;1m"
	Bad  = "\033[31;1;1m"
	Good = "\033[32;1;1m"
	Run  = "\033[97;1;1m"
)
