#!/usr/bin/env ruby

require 'ruby-audio'

def snd_to_file_a input, t_2, samples
	output = []
	(t_2...samples).step(t_2*2) do |n|
		if input[n] > 0.0
			output << "1"
		else
			output << "0"
		end
	end
	output
end

abort "usage: #{$0} <input> <output> <T/2>" if ARGV.length != 3

input_file = ARGV[0]
output_file = ARGV[1]
t_2 = ARGV[2].to_i
	
RubyAudio::Sound.open(input_file) do |input_snd|
	samplerate = input_snd.info.samplerate
	samples = (input_snd.info.length * samplerate).to_i
	input_a = input_snd.read(:float, samples).to_a

	hidden_data = []
	input_a.each do |v|
		hidden_data << v[1]
	end

	# De-modulation
	hidden_a = snd_to_file_a hidden_data, t_2, samples

	# Write output file
	hidden_a_bin = hidden_a.each_slice(8).map do |a|
		a.join.to_i(2)
	end
	File.open(output_file, 'wb' ) do |out|
		  out.write hidden_a_bin.pack("C*")
	end
end
