package ascii

func Ascii_saludo() string {
	saludo1 :=
		`
  __  ____                                
 (  )(  _ \                               
  )(  ) __/                               
 (__)(__)
  `

	saludo2 :=
		`
  ____   ___   __   __ _  __ _  ____  ____ 
/ ___) / __) / _\ (  ( \(  ( \(  __)(  _ \
\___ \( (__ /    \/    //    / ) _)  )   /
(____/ \___)\_/\_/\_)__)\_)__)(____)(__\_)
  `

	saludo3 := saludo1 + saludo2

	return saludo3

}
