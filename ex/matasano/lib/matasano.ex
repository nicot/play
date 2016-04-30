defmodule Matasano do
  def to64(hex) do
    {:ok, bytes} = Base.decode16(hex, case: :mixed)
    Base.encode64(bytes)
  end

  def toHex(base64) do
    {:ok, bytes} = Base.decode64(base64, case: :mixed)
    Base.encode16(bytes)
  end

  def xor(a, b) do
    use Bitwise
    buf1 = :binary.bin_to_list(a)
    buf2 = :binary.bin_to_list(b)
    bytes = Enum.zip(buf1, buf2)
    r = Enum.map(bytes, fn({a, b}) -> a ^^^ b end)
    Base.encode16(to_string(r))
  end

  def xorCipher(hex, char) do
    use Bitwise
    {:ok, buf} = Base.decode16(hex, case: :mixed)
    bytes = :binary.bin_to_list(buf)
    Enum.map(bytes, fn(b) -> b  ^^^ char end)
  end

  # Now that the party is jumping.
  def detectXor() do
    :inets.start
    url = 'http://cryptopals.com/static/challenge-data/4.txt'
    {:ok, {{_, 200, _}, _, body}} = :httpc.request url
    lines = String.split to_string(body), "\n"
    for n <- 1..255, do:
    Enum.map(lines, fn(line) ->
      decoded = to_string xorCipher line, n
      if String.printable?(decoded) and String.length(decoded) == byte_size(decoded) do
        IO.puts decoded
      end
    end)
  end

  def repeatingXor(string, key) do
    for x <-
    chunks = Enum.chunk(string, byte_size key)
    Enum.map(chunks, fn cs -> xor cs, key end)
  end
end
