#!/usr/bin/ruby
require "csv"

envstats = {}
CSV.open(ARGV[0], :headers => true) do |csv|
  csv.each do |row|
    generation = row["Generation"].to_i
    environment = row["Environment"]
    bd_ratio = row["B/D Ratio"].to_f
    fitness = row["Fitness"].to_f

    stats = envstats[environment]
    if stats == nil
      stats = {
        :env => environment.to_i,
        :initial_bd => bd_ratio,
        :final_bd => bd_ratio,
        :min_bd => bd_ratio,
        :max_bd => bd_ratio,
        :avg_bd => 0,
        :num_gens => 0,
        :avg_fitness => 0
      }
    end
    stats[:final_bd] = bd_ratio
    if bd_ratio < stats[:min_bd]
      stats[:min_bd] = bd_ratio
    end
    if bd_ratio > stats[:max_bd]
      stats[:max_bd] = bd_ratio
    end
    stats[:avg_bd] += bd_ratio
    stats[:num_gens] += 1
    stats[:avg_fitness] += fitness

    envstats[environment] = stats
  end
end

envstats.each do |k, stats|
  stats.each do |statk, val|
    if statk.to_s[0..3] == "avg_"
      stats[statk] = val / stats[:num_gens]
    end
  end
  stats[:diff_bd] = stats[:final_bd] - stats[:initial_bd]
end

envary = envstats.values.sort_by{|x| x[:env]}

envary.each_with_index do |x, idx| 
  if idx > 0
    x[:bd_jump] = x[:initial_bd] - envary[idx - 1][:final_bd]
  else
    x[:bd_jump] = "N/A"
  end
end

keylist = envary[0].keys.sort
CSV.open(ARGV[1], "w") do |csv|
  csv << keylist
  envary.each do |x|
    csv << x.values_at(*keylist)
  end
end
