package ascii

var (
  Blue = "\033[34m"
  Reset = "\033[0m" 
)

func Ascii_saludo() string {
	saludo1 := (
  Blue +
		`  
  __  ____                                
 (  )(  _ \                               
  )(  ) __/                               
 (__)(__)
  `+ Reset)

	saludo2 := (
  Blue +
		`
 ____   ___   __   __ _  __ _  ____  ____ 
/ ___) / __) / _\ (  ( \(  ( \(  __)(  _ \
\___ \( (__ /    \/    //    / ) _)  )   /
(____/ \___)\_/\_/\_)__)\_)__)(____)(__\_)
  ` + Reset)

	saludo3 := saludo1 + saludo2

	return saludo3

}
