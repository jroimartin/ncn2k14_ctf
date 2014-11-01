#!/usr/bin/env ruby

require 'socket'
require 'sand'

IP = "127.0.0.1"
PORT = 6767

server = TCPServer.new(IP, PORT)
loop do
	Thread.start(server.accept) do |client|
		line = client.gets.untaint
		input = line.delete("\n\r").split(';')
		output = Sandrbox.perform(input).output
		client.puts output
		client.close
	end
end
