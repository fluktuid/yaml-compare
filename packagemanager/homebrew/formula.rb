class Yaml-Compare < Formula
  desc "A command line tool to compare two yaml files"
  homepage "https://github.com/fluktuid/yaml-compare"
  url "https://github.com/fluktuid/yaml-compare/archive/0.0.0.tar.gz"
  sha256 "e9ace872c860197d85f7650502ecb30f74c180faddae6ea77b49aeb2a7259e4c"
  version "0.0.0"

  depends_on "curl"

  bottle :unneeded

  def install
    bin.install "weather"
  end
end
