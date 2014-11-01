import wave
import struct

wav = wave.open('output.wav')
data = wav.readframes(wav.getnframes())
adata = struct.unpack("<%dI" % (len(data)/4), data)
wav.close()

f = open('output.png', 'wb')
for i in range(0, len(adata), 2):
	f.write(chr(adata[i+1] - adata[i]))
f.close()
