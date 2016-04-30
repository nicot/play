defmodule Matasano do
  def to64(hex) do
    {:ok, bytes} = Base.decode16(hex, case: :mixed)
    Base.encode64(bytes)
  end

  def toHex(base64) do
    {:ok, bytes} = Base.decode64(base64, case: :mixed)
    Base.encode16(bytes)
  end

  def xor(hex1, hex2) do
    use Bitwise
    {:ok, buf1} = Base.decode16(hex1, case: :mixed)
    {:ok, buf2} = Base.decode16(hex2, case: :mixed)
    buf1 ^^^ buf2
  end
end
