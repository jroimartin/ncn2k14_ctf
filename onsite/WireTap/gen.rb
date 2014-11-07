#!/usr/bin/env ruby

require 'ruby-audio'

abort "usage: #{$0} <input_base> <input_hidden> <outut>" if ARGV.length != 3

input_file = ARGV[0]
hidden_file = ARGV[1]
output_file = ARGV[2]
	
RubyAudio::Sound.open(input_file) do |input_snd|
	samplerate = input_snd.info.samplerate
	samples = (input_snd.info.length * samplerate).to_i
	input_a = input_snd.read(:int, samples).to_a

	# Read data to be hidden
	hidden_data = File.read hidden_file
	hidden_data = hidden_data.force_encoding('BINARY')

	abort "input_hidden is too big!" if input_a.length < hidden_data.length

	hidden_a = input_a.each_with_index.map do |input, idx|
		hidx = idx % hidden_data.length
		input + hidden_data[hidx].ord
	end

	# Write output
	puts "Writing output..."
	output_buf = RubyAudio::Buffer.int(input_a.length, 2)
	(0...input_a.length).each do |idx|
		output_buf[idx] = [input_a[idx], hidden_a[idx]]
	end
	info = RubyAudio::SoundInfo.new :channels => 2, :samplerate => samplerate,
		:format => RubyAudio::FORMAT_WAV|RubyAudio::FORMAT_PCM_32
	RubyAudio::Sound.open(output_file, 'w', info) do |output_snd|
		output_snd.write(output_buf)
	end
end
