#!/usr/bin/env ruby

require 'ruby-audio'

abort "usage: #{$0} <input> <output>" if ARGV.length != 2

input_file = ARGV[0]
output_file = ARGV[1]

RubyAudio::Sound.open(input_file) do |input_snd|
	samplerate = input_snd.info.samplerate
	samples = (input_snd.info.length * samplerate).to_i
	input_a = input_snd.read(:int, samples).to_a

	hidden_data = Array.new
	input_a.each do |v|
		hidden_data << (v[1] - v[0])
	end

	# Write output file
	File.open(output_file, 'wb' ) do |out|
		  out.write hidden_data.pack("C*")
	end
end
