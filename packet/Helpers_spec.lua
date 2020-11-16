
local Helpers = require("packet/Helpers")

describe("Helpers test", function()
	it("should match high/low expectations", function()
		assert.are.same(65500, Helpers.bytes_to_int(0xff, 0xdc))
	end)

	it("should parse the generated int/byte", function()

		local input = 40000
		local bytes = Helpers.int_to_bytes(input)

		local high = string.byte(bytes, 1)
		local low = string.byte(bytes, 2)

		print("High: " .. high .. " Low: " .. low)

		local output = Helpers.bytes_to_int(high, low)

		assert.are.same(input, output)
	end)
end)
