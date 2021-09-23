all:
	gcc -c gmp/mini-gmp.c -o mini-gmp.o
	gcc -c sha256.c -o sha.o
	gcc -Igmp -c srp.cpp -o srp.o
	#TODO link properly
	gcc -I/usr/include/lua5.1 -llua5.3 -lgmp -shared -fpic main.cpp -o ../srp.so
