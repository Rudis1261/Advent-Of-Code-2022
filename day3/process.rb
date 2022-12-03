#!/usr/bin/env ruby

require "json"

input = File.read("./input.txt")
lines = input.split("\n")
total = 0
badgeTotal = 0
priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ".split("")

lines.each_slice(3) do |group|
    group.each do |item|
        item = item.split("")
        compartment1 = item[0...item.size/2]
        compartment2 = item[item.size/2...]
        commonItem = compartment1.intersection(compartment2).first
        total += priority.index(commonItem) + 1
    end

    group.map!{|item| item.split("")}
    char = group[0].intersection(group[1]).intersection(group[2]).first
    badgeTotal += priority.index(char) + 1
end

puts JSON.pretty_generate({
    "Part 1 total": total,
    "Part 2 total": badgeTotal,
})