#!/usr/bin/env ruby

sol = "rrrrrffffrrrrrrffffffffffllllbblllllllrrbbbbbbrrrrbbbrrrfffrrrrf"
rsol = ""
sol.each_char.with_index do |ch, i|
  rch = ""
  if ch.ord + i%26 > "z".ord
    rch = (ch.ord + i%26 - 26).chr
  else
    rch = (ch.ord + i%26).chr
  end
  rsol << rch
end
puts rsol
