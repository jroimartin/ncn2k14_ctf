#!/usr/bin/env ruby

require 'ruby-audio'

def file_to_snd_a input, amp, samples
	output = []
	period = samples / (input.length * 8)
	nsamples = period/2
	abort "Freq is too high!" if nsamples < 1
	input.each_byte do |v|
		binstr = v.to_s(2)
		binstr = Array.new(8-binstr.length, "0").join + binstr
		binstr.each_char do |b|
			if b == "1"
				output = output.concat([0.0] * nsamples)
				output = output.concat([amp.to_f] * nsamples)
			else
				output = output.concat([amp.to_f] * nsamples)
				output = output.concat([0.0] * nsamples)
			end
		end
	end
	output.concat(Array.new(samples-output.length, 0.0))
end

abort "usage: #{$0} <input_base> <input_hidden> <outut>" if ARGV.length != 3

input_file = ARGV[0]
hidden_file = ARGV[1]
output_file = ARGV[2]

RubyAudio::Sound.open(input_file) do |input_snd|
	samplerate = input_snd.info.samplerate
	samples = (input_snd.info.length * samplerate).to_i
	input_a = input_snd.read(:float, samples).to_a

	# Read data to be hidden
	hidden_data = File.read hidden_file
	hidden_a = file_to_snd_a hidden_data, 0.01, samples

	# Write output
	puts "Writing output..."
	output_buf = RubyAudio::Buffer.float(input_a.length, 2)
	(0...input_a.length).each do |idx|
		output_buf[idx] = [input_a[idx], hidden_a[idx]]
	end
	info = RubyAudio::SoundInfo.new :channels => 2, :samplerate => samplerate,
		:format => RubyAudio::FORMAT_WAV|RubyAudio::FORMAT_PCM_16
	RubyAudio::Sound.open(output_file, 'w', info) do |output_snd|
		output_snd.write(output_buf)
	end
end
