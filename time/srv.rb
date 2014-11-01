#!/usr/bin/env ruby

require 'socket'

FLAG = "tIMeMaTTerS"
IP = "127.0.0.1"
PORT = 6969
DELAY = 0.25

server = TCPServer.new(IP, PORT)
loop do
	Thread.start(server.accept) do |s|
		while line = s.gets.delete("\n\r")
			gotit = false
			line.each_char.with_index do |ch, idx|
				break if ch != FLAG[idx]
				sleep(DELAY)
				if idx == FLAG.length-1 and line.length == FLAG.length
					gotit = true
					break
				end
			end
			if gotit
				s.puts "Yeah! You got it! :)"
			else
				s.puts "Invalid key"
			end
		end
		s.close
	end
end
