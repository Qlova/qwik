concept f()
	print("hi")
}

concept call(a)
	a()
}

concept add(a,b)
	return a+b
}

main
	f()
	b = f
	b()
	call(b)
	print(add(2, 3), 8, 2, 1)
}
