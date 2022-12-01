#!/usr/bin/env ruby 

require 'json'

lines = File.read("input.txt").split("\n")
elves = []
elfIndex = 0

lines.each do |line|
	if line == ""
		elfIndex += 1 
		next 
	end

	calories = Integer(line)
	elves[elfIndex] ||= 0 
	elves[elfIndex] = elves[elfIndex] + calories
end.sort!

puts JSON.pretty_generate({
	max: elves.last,
	sum_of_max_three: elves[-3..].sum,
})