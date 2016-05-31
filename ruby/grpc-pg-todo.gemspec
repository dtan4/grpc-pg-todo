Gem::Specification.new do |s|
  s.name          = "grpc-pg-todo"
  s.version       = "0.1.0"
  s.authors       = ["Daisuke Fujita"]
  s.email         = "dtanshi45@gmail.com"
  s.homepage      = "https://github.com/dtan4/grpc-pg-todo"
  s.summary       = "gRPC + PostgreSQL sample"
  s.description   = "todo application using gRPC + PostgreSQL"

  s.files         = `git ls-files -- ruby/*`.split("\n")
  s.executables   = `git ls-files -- ruby/greeter*.rb ruby/route_guide/*.rb`.split("\n").map do |f|
    File.basename(f)
  end

  s.require_paths = ["lib"]
  s.platform      = Gem::Platform::RUBY

  s.add_dependency "google-protobuf", "3.0.0.alpha.5.0.5.1"
  s.add_dependency "grpc", "~> 0.14.1"
  s.add_dependency "pg", "~> 0.18.2"

  s.add_development_dependency "bundler", "~> 1.7"
end
