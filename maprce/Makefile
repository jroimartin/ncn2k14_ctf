all:
	gcc -Wall -o maprce maprce.c
	strip maprce

solve:
	dd if=maprce of=sol.bin bs=1 count=128 &> /dev/null
	cat sol.bin | ./maprce
	shasum sol.bin | awk '{print "NcN_"$$1}'

clean:
	rm -f maprce sol.bin
