
concept add(a, b)
	return a+b
}

main
	a = 50
	b = 50
	
	e = 2
	
	c = [a*e, 32, add(a, b), 32, 98]

	for value in c: write(value, " ")
	write("\n")
}
